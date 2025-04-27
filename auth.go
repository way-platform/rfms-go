package rfms

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// TokenCredentials for an authenticated rFMS API client.
type TokenCredentials struct {
	// Token is the bearer token for the authenticated client.
	Token string `json:"token"`
	// TokenExpireTime is the time when the token expires.
	TokenExpireTime time.Time `json:"tokenExpireTime"`
	// RefreshToken is the refresh token for the authenticated client.
	RefreshToken string `json:"refreshToken"`
	// RefreshTokenExpireTime is the time when the refresh token expires.
	RefreshTokenExpireTime time.Time `json:"refreshTokenExpireTime"`
}

// TokenAuthenticator is a pluggable interface for authenticating requests to an rFMS API.
type TokenAuthenticator interface {
	// Authenticate the client and return a set of [TokenCredentials].
	Authenticate(ctx context.Context) (TokenCredentials, error)
	// Refresh the token credentials.
	Refresh(ctx context.Context, refreshToken string) (TokenCredentials, error)
}

type tokenAuthenticatorTransport struct {
	tokenAuthenticator TokenAuthenticator
	transport          http.RoundTripper
	mu                 sync.Mutex
	credentials        TokenCredentials
}

func (t *tokenAuthenticatorTransport) RoundTrip(req *http.Request) (_ *http.Response, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("token authenticator transport: %w", err)
		}
	}()
	token, err := t.getToken(req.Context())
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	return t.transport.RoundTrip(req)
}

func (t *tokenAuthenticatorTransport) getToken(ctx context.Context) (string, error) {
	t.mu.Lock()
	credentials := t.credentials
	t.mu.Unlock()
	if credentials.TokenExpireTime.After(time.Now()) {
		return credentials.Token, nil
	}
	if credentials.RefreshToken != "" && credentials.RefreshTokenExpireTime.After(time.Now()) {
		if newCredentials, err := t.tokenAuthenticator.Refresh(ctx, credentials.RefreshToken); err == nil {
			t.mu.Lock()
			t.credentials = newCredentials
			t.mu.Unlock()
			return newCredentials.Token, nil
		}
	}
	newCredentials, err := t.tokenAuthenticator.Authenticate(ctx)
	if err != nil {
		return "", fmt.Errorf("authenticate: %w", err)
	}
	t.mu.Lock()
	t.credentials = newCredentials
	t.mu.Unlock()
	return newCredentials.Token, nil
}

type reuseTokenCredentialsTransport struct {
	transport   http.RoundTripper
	credentials TokenCredentials
}

func (t *reuseTokenCredentialsTransport) RoundTrip(req *http.Request) (_ *http.Response, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("reuse token credentials transport: %w", err)
		}
	}()
	req.Header.Set("Authorization", "Bearer "+t.credentials.Token)
	return t.transport.RoundTrip(req)
}
