package rfms

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"
)

// ScaniaAuthBaseURL is the Scania auth server's base URL.
const ScaniaAuthBaseURL = "https://dataaccess.scania.com/auth"

// ScaniaAuthConfig describes Scania's custom OAuth2-like flow using HMAC-SHA256 challenge-response.
type ScaniaAuthConfig struct {
	// ClientID is the application's ID.
	ClientID string

	// ClientSecret is the application's secret key (base64 URL encoded).
	ClientSecret string

	// BaseURL is the Scania auth server's token endpoint URL.
	// Defaults to [ScaniaAuthBaseURL] if empty.
	BaseURL string

	// MaxRetries is the maximum number of times to retry a request.
	MaxRetries int

	// Timeout is the timeout for a request.
	Timeout time.Duration

	// Debug is whether to enable debug logging.
	Debug bool
}

// TokenSource returns an oauth2.TokenSource that retrieves tokens using
// Scania's HMAC-SHA256 challenge-response mechanism.
// The returned TokenSource automatically handles token caching and refresh.
func (c ScaniaAuthConfig) TokenSource(ctx context.Context) oauth2.TokenSource {
	source := &scaniaTokenSource{
		ctx:    ctx,
		config: c,
	}
	// Use oauth2.ReuseTokenSource for automatic token caching and refresh
	return oauth2.ReuseTokenSource(nil, source)
}

// scaniaTokenSource implements oauth2.TokenSource for Scania's custom auth flow.
type scaniaTokenSource struct {
	ctx    context.Context
	config ScaniaAuthConfig
}

var _ oauth2.TokenSource = (*scaniaTokenSource)(nil)

// Token implements oauth2.TokenSource.
// It performs the complete Scania authentication flow: challenge -> response -> token.
func (s *scaniaTokenSource) Token() (*oauth2.Token, error) {
	// Step 1: Get challenge
	challenge, err := s.getChallenge()
	if err != nil {
		return nil, fmt.Errorf("scania auth: get challenge: %w", err)
	}
	// Step 2: Compute challenge response
	challengeResponse, err := computeScaniaChallengeResponse(challenge, s.config.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("scania auth: compute challenge response: %w", err)
	}
	// Step 3: Exchange response for token
	token, err := s.exchangeToken(challengeResponse)
	if err != nil {
		return nil, fmt.Errorf("scania auth: exchange token: %w", err)
	}
	return token, nil
}

// getChallenge requests an authentication challenge from Scania.
func (s *scaniaTokenSource) getChallenge() (string, error) {
	req, err := s.newRequest(http.MethodPost, "/clientid2challenge", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	body := url.Values{}
	body.Set("clientId", s.config.ClientID)
	req.Body = io.NopCloser(bytes.NewBufferString(body.Encode()))
	resp, err := s.httpClient().Do(req)
	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var response struct {
		Challenge string `json:"challenge"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}
	return response.Challenge, nil
}

// exchangeToken exchanges a challenge response for an OAuth2 token.
func (s *scaniaTokenSource) exchangeToken(challengeResponse string) (*oauth2.Token, error) {
	req, err := s.newRequest(http.MethodPost, "/response2token", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	body := url.Values{}
	body.Set("clientId", s.config.ClientID)
	body.Set("Response", challengeResponse)
	req.Body = io.NopCloser(bytes.NewBufferString(body.Encode()))
	resp, err := s.httpClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var response struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	now := time.Now()
	token := &oauth2.Token{
		AccessToken:  response.Token,
		TokenType:    "Bearer", // Standard OAuth2 token type
		RefreshToken: response.RefreshToken,
		Expiry:       now.Add(time.Hour), // Scania tokens expire in 1 hour
	}
	// Store additional Scania-specific data using WithExtra
	return token.WithExtra(map[string]any{
		"refresh_token_expiry": now.Add(24 * time.Hour).Unix(),
	}), nil
}

// newRequest creates a new HTTP request with proper context and user agent.
func (s *scaniaTokenSource) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	tokenURL := s.config.BaseURL
	if tokenURL == "" {
		tokenURL = ScaniaAuthBaseURL
	}
	req, err := http.NewRequestWithContext(s.ctx, method, tokenURL+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", getUserAgent())
	return req, nil
}

// httpClient returns the HTTP client to use, with retry transport if none specified.
func (s *scaniaTokenSource) httpClient() *http.Client {
	transport := http.DefaultTransport
	if s.config.Debug {
		transport = &debugTransport{
			next: transport,
		}
	}
	if s.config.MaxRetries > 0 {
		transport = &retryTransport{
			maxRetries: s.config.MaxRetries,
			next:       transport,
		}
	}
	return &http.Client{
		Transport: transport,
		Timeout:   s.config.Timeout,
	}
}

// computeScaniaChallengeResponse calculates the HMAC-SHA256 of a challenge using a secret key.
// The challenge, secret key, and result are base64 URL encoded without padding.
func computeScaniaChallengeResponse(challenge, secretKey string) (string, error) {
	decodedSecretKey, err := base64.RawURLEncoding.DecodeString(secretKey)
	if err != nil {
		return "", fmt.Errorf("decode secret key: %w", err)
	}
	decodedChallenge, err := base64.RawURLEncoding.DecodeString(challenge)
	if err != nil {
		return "", fmt.Errorf("decode challenge: %w", err)
	}
	mac := hmac.New(sha256.New, decodedSecretKey)
	mac.Write(decodedChallenge)
	responseBytes := mac.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(responseBytes), nil
}
