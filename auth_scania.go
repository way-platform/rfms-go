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
)

// scaniaTokenAuthenticator is an [TokenAuthenticator] for Scania's rFMS API.
type scaniaTokenAuthenticator struct {
	baseURL      string
	clientID     string
	clientSecret string
	httpClient   *http.Client
}

// NewScaniaTokenAuthenticator creates a new [TokenAuthenticator] for Scania's rFMS API.
func NewScaniaTokenAuthenticator(clientID string, clientSecret string) TokenAuthenticator {
	return &scaniaTokenAuthenticator{
		baseURL:      ScaniaAuthBaseURL,
		clientID:     clientID,
		clientSecret: clientSecret,
		httpClient:   http.DefaultClient,
	}
}

var _ TokenAuthenticator = &scaniaTokenAuthenticator{}

// Authenticate implements the [TokenAuthenticator] interface.
func (a *scaniaTokenAuthenticator) Authenticate(ctx context.Context) (_ TokenCredentials, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("authenticate Scania: %w", err)
		}
	}()
	challenge, err := a.getChallenge(ctx)
	if err != nil {
		return TokenCredentials{}, err
	}
	challengeResponse, err := computeScaniaChallengeResponse(challenge, a.clientSecret)
	if err != nil {
		return TokenCredentials{}, err
	}
	credentials, err := a.exchangeToken(ctx, challengeResponse)
	if err != nil {
		return TokenCredentials{}, err
	}
	return credentials, nil
}

// Refresh implements the [TokenAuthenticator] interface.
func (a *scaniaTokenAuthenticator) Refresh(
	ctx context.Context,
	refreshToken string,
) (_ TokenCredentials, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("refresh Scania: %w", err)
		}
	}()
	req, err := a.newRequest(ctx, http.MethodPost, "/refreshtoken", nil)
	if err != nil {
		return TokenCredentials{}, nil
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	body := url.Values{}
	body.Set("clientId", a.clientID)
	body.Set("RefreshToken", refreshToken)
	req.Body = io.NopCloser(bytes.NewBufferString(body.Encode()))
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return TokenCredentials{}, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return TokenCredentials{}, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	var response struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return TokenCredentials{}, fmt.Errorf("decode response: %w", err)
	}
	now := time.Now()
	return TokenCredentials{
		Token:                  response.Token,
		TokenExpireTime:        now.Add(time.Hour),
		RefreshToken:           response.RefreshToken,
		RefreshTokenExpireTime: now.Add(time.Hour * 24),
	}, nil
}

// getChallenge requests an authentication challenge.
func (a *scaniaTokenAuthenticator) getChallenge(ctx context.Context) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("get challenge: %w", err)
		}
	}()
	req, err := a.newRequest(ctx, http.MethodPost, "/clientid2challenge", nil)
	if err != nil {
		return "", nil
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	body := url.Values{}
	body.Set("clientId", a.clientID)
	req.Body = io.NopCloser(bytes.NewBufferString(body.Encode()))
	res, err := a.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("response status code: %d", res.StatusCode)
	}
	var response struct {
		Challenge string `json:"challenge"`
	}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}
	return response.Challenge, nil
}

// exchangeToken exchanges a challenge response for a token.
func (a *scaniaTokenAuthenticator) exchangeToken(
	ctx context.Context,
	challengeResponse string,
) (_ TokenCredentials, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("exchange token: %w", err)
		}
	}()
	req, err := a.newRequest(ctx, http.MethodPost, "/response2token", nil)
	if err != nil {
		return TokenCredentials{}, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	body := url.Values{}
	body.Set("clientId", a.clientID)
	body.Set("Response", challengeResponse)
	req.Body = io.NopCloser(bytes.NewBufferString(body.Encode()))
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return TokenCredentials{}, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return TokenCredentials{}, fmt.Errorf("response status code: %d", resp.StatusCode)
	}
	var response struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return TokenCredentials{}, fmt.Errorf("decode response: %w", err)
	}
	now := time.Now().UTC()
	return TokenCredentials{
		Token:                  response.Token,
		TokenExpireTime:        now.Add(time.Hour),
		RefreshToken:           response.RefreshToken,
		RefreshTokenExpireTime: now.Add(time.Hour * 24),
	}, nil
}

func (a *scaniaTokenAuthenticator) newRequest(
	ctx context.Context,
	method string,
	path string,
	body io.Reader,
) (_ *http.Request, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("new request: %w", err)
		}
	}()
	request, err := http.NewRequestWithContext(ctx, method, a.baseURL+path, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", getUserAgent())
	return request, nil
}

// computeScaniaChallengeResponse calculates the HMAC-SHA256 of a challenge using a secret key.
// The challenge, secret key, and result are base64 URL encoded without padding.
func computeScaniaChallengeResponse(challenge string, secretKey string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("compute challenge response: %w", err)
		}
	}()
	decodedSecretKey, err := base64.RawURLEncoding.DecodeString(secretKey)
	if err != nil {
		return "", fmt.Errorf("decode secret key: %w", err)
	}
	decodedChallenge, err := base64.RawURLEncoding.DecodeString(challenge)
	if err != nil {
		return "", fmt.Errorf("decode challenge: %w", err)
	}
	mac := hmac.New(sha256.New, decodedSecretKey)
	_, _ = mac.Write(decodedChallenge) // never returns an error for hash implementsations
	responseBytes := mac.Sum(nil)      // nil allocates a new slice to hold the result
	return base64.RawURLEncoding.EncodeToString(responseBytes), nil
}
