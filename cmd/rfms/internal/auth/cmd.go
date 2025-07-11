package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/way-platform/rfms-go"
)

// Credentials for the rFMS CLI.
type Credentials struct {
	Provider         string                 `json:"provider"`
	TokenCredentials *rfms.TokenCredentials `json:"tokenCredentials,omitzero"`
	Username         string                 `json:"username,omitzero"`
	Password         string                 `json:"password,omitzero"`
}

func resolveCredentialsFilepath() (string, error) {
	return xdg.ConfigFile("rfms-go/auth.json")
}

// ReadCredentials reads the rFMS CLI credentials.
func ReadCredentials() (*Credentials, error) {
	credentialsFilepath, err := resolveCredentialsFilepath()
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(credentialsFilepath); err != nil {
		return nil, err
	}
	data, err := os.ReadFile(credentialsFilepath)
	if err != nil {
		return nil, err
	}
	var credentials Credentials
	if err := json.Unmarshal(data, &credentials); err != nil {
		return nil, err
	}
	return &credentials, nil
}

// NewClient creates a new rFMS client using the CLI credentials.
func NewClient() (*rfms.Client, error) {
	auth, err := ReadCredentials()
	if err != nil {
		return nil, err
	}
	switch auth.Provider {
	case rfms.BrandScania:
		if auth.TokenCredentials.TokenExpireTime.Before(time.Now()) {
			if auth.TokenCredentials.RefreshTokenExpireTime.Before(time.Now()) {
				// TODO: Refresh token credentials if they are expired.
				_ = "TODO"
			}
			return nil, fmt.Errorf("session expired - please login again")
		}
		return rfms.NewClient(
			rfms.WithBaseURL(rfms.ScaniaBaseURL),
			rfms.WithVersion(rfms.V4),
			rfms.WithReuseTokenAuth(*auth.TokenCredentials),
		), nil
	case rfms.BrandVolvoTrucks:
		return rfms.NewClient(
			rfms.WithVolvoTrucks(auth.Username, auth.Password),
		), nil
	default:
		return nil, fmt.Errorf("unknown provider: %s", auth.Provider)
	}
}

func writeCredentials(credentials *Credentials) error {
	credentialsFilepath, err := resolveCredentialsFilepath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(credentials, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(credentialsFilepath, data, 0o600)
}

func removeCredentials() error {
	credentialsFilepath, err := resolveCredentialsFilepath()
	if err != nil {
		return err
	}
	return os.RemoveAll(credentialsFilepath)
}

// NewCommand returns a new [cobra.Command] for rFMS CLI authentication.
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate with an rFMS API",
	}
	cmd.AddCommand(newLoginCommand())
	cmd.AddCommand(newLogoutCommand())
	return cmd
}

func newLoginCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to an rFMS API",
	}
	cmd.AddCommand(newLoginScaniaCommand())
	cmd.AddCommand(newLoginVolvoTrucksCommand())
	// TODO: Implement login with unknown provider.
	return cmd
}

func newLoginScaniaCommand() *cobra.Command {
	const provider = rfms.BrandScania
	cmd := &cobra.Command{
		Use:   "scania",
		Short: "Login to the Scania rFMS API",
	}
	clientID := cmd.Flags().String("client-id", "", "client ID to use for authentication")
	cmd.MarkFlagRequired("client-id")
	clientSecret := cmd.Flags().String("client-secret", "", "client secret to use for authentication")
	cmd.MarkFlagRequired("client-secret")
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		authenticator := rfms.NewScaniaTokenAuthenticator(*clientID, *clientSecret)
		credentials, err := authenticator.Authenticate(cmd.Context())
		if err != nil {
			return err
		}
		auth := &Credentials{
			Provider:         provider,
			TokenCredentials: &credentials,
		}
		if err := writeCredentials(auth); err != nil {
			return err
		}
		cmd.Printf("Logged in to %s.", provider)
		return nil
	}
	return cmd
}

func newLoginVolvoTrucksCommand() *cobra.Command {
	const provider = rfms.BrandVolvoTrucks
	cmd := &cobra.Command{
		Use:   "volvo-trucks",
		Short: "Login to the Volvo Trucks rFMS API",
	}
	username := cmd.Flags().String("username", "", "username to use for authentication")
	cmd.MarkFlagRequired("username")
	password := cmd.Flags().String("password", "", "password to use for authentication")
	cmd.MarkFlagRequired("password")
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		auth := &Credentials{
			Provider: provider,
			Username: *username,
			Password: *password,
		}
		if err := writeCredentials(auth); err != nil {
			return err
		}
		cmd.Printf("Logged in to %s.", provider)
		return nil
	}
	return cmd
}

func newLogoutCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Logout from the current authenticated rFMS API",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := removeCredentials(); err != nil {
				return err
			}
			cmd.Println("Logged out.")
			return nil
		},
	}
}
