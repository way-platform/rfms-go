package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
)

// ScaniaCredentials holds Scania rFMS OAuth2 client credentials.
type ScaniaCredentials struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// VolvoCredentials holds Volvo Trucks rFMS HTTP Basic credentials.
type VolvoCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ScaniaCredentialStore reads and writes Scania credentials.
type ScaniaCredentialStore interface {
	Load() (*ScaniaCredentials, error)
	Save(*ScaniaCredentials) error
	Clear() error
}

// VolvoCredentialStore reads and writes Volvo Trucks credentials.
type VolvoCredentialStore interface {
	Load() (*VolvoCredentials, error)
	Save(*VolvoCredentials) error
	Clear() error
}

// fileStore is a generic JSON file-backed credential store.
type fileStore[T any] struct{ path string }

func (s *fileStore[T]) Load() (*T, error) {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return nil, fmt.Errorf("read credentials: %w", err)
	}
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, fmt.Errorf("unmarshal credentials: %w", err)
	}
	return &v, nil
}

func (s *fileStore[T]) Save(v *T) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal credentials: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(s.path), 0o700); err != nil {
		return fmt.Errorf("create credentials dir: %w", err)
	}
	return os.WriteFile(s.path, data, 0o600)
}

func (s *fileStore[T]) Clear() error {
	err := os.Remove(s.path)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

// NewScaniaCredentialFileStore creates a file-backed Scania credential store.
func NewScaniaCredentialFileStore(path string) ScaniaCredentialStore {
	return &fileStore[ScaniaCredentials]{path: path}
}

// NewVolvoCredentialFileStore creates a file-backed Volvo credential store.
func NewVolvoCredentialFileStore(path string) VolvoCredentialStore {
	return &fileStore[VolvoCredentials]{path: path}
}

// TokenStore reads and writes OAuth2 tokens.
type TokenStore interface {
	Load() (*oauth2.Token, error)
	Save(*oauth2.Token) error
	Clear() error
}

// NewTokenFileStore creates a file-backed token store.
func NewTokenFileStore(path string) TokenStore {
	return &fileStore[oauth2.Token]{path: path}
}

// Option configures the CLI command tree.
type Option func(*config)

type config struct {
	scaniaCredentialStore ScaniaCredentialStore
	volvoCredentialStore  VolvoCredentialStore
	tokenStore            TokenStore
	httpClient            *http.Client
}

// WithScaniaCredentialStore sets the credential store for Scania rFMS.
func WithScaniaCredentialStore(s ScaniaCredentialStore) Option {
	return func(c *config) { c.scaniaCredentialStore = s }
}

// WithVolvoCredentialStore sets the credential store for Volvo Trucks rFMS.
func WithVolvoCredentialStore(s VolvoCredentialStore) Option {
	return func(c *config) { c.volvoCredentialStore = s }
}

// WithTokenStore sets the token store (used for Scania OAuth2 tokens).
func WithTokenStore(s TokenStore) Option {
	return func(c *config) { c.tokenStore = s }
}

// WithHTTPClient sets the base [http.Client] passed to the SDK client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *config) { c.httpClient = httpClient }
}
