package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Store reads and writes JSON-serializable data.
type Store interface {
	Read(target any) error
	Write(data any) error
	Clear() error
}

// Credentials for rFMS CLI authentication.
// Provider determines which auth method is used:
//   - "SCANIA": OAuth2 client credentials (ClientID + ClientSecret)
//   - "VOLVO TRUCKS": HTTP Basic Auth (Username + Password)
type Credentials struct {
	Provider     string `json:"provider"`
	ClientID     string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
}

// Option configures the CLI command tree.
type Option func(*config)

type config struct {
	credentialStore Store
	tokenStore      Store
}

// WithCredentialStore sets the credential store.
func WithCredentialStore(s Store) Option {
	return func(c *config) { c.credentialStore = s }
}

// WithTokenStore sets the token store (used for Scania OAuth2 tokens).
func WithTokenStore(s Store) Option {
	return func(c *config) { c.tokenStore = s }
}

// FileStore is a JSON file-backed store.
type FileStore struct {
	path string
}

// NewFileStore creates a new file-backed store at the given path.
func NewFileStore(path string) *FileStore {
	return &FileStore{path: path}
}

// Read unmarshals the file contents into target.
func (s *FileStore) Read(target any) error {
	data, err := os.ReadFile(s.path)
	if err != nil {
		return fmt.Errorf("read store: %w", err)
	}
	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("unmarshal store: %w", err)
	}
	return nil
}

// Write marshals data and writes it to the file.
func (s *FileStore) Write(data any) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal store: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(s.path), 0o700); err != nil {
		return fmt.Errorf("create store dir: %w", err)
	}
	return os.WriteFile(s.path, bytes, 0o600)
}

// Clear removes the file.
func (s *FileStore) Clear() error {
	err := os.Remove(s.path)
	if err != nil && os.IsNotExist(err) {
		return nil
	}
	return err
}
