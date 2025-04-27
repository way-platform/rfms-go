package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/way-platform/rfms-go/v4"
)

type authFile struct {
	Provider         string                 `json:"provider"`
	TokenCredentials *rfms.TokenCredentials `json:"tokenCredentials,omitzero"`
	Username         string                 `json:"username,omitzero"`
	Password         string                 `json:"password,omitzero"`
}

const authConfigFile = "rfms-go/auth.json"

func readAuthFile() (*authFile, error) {
	authFilepath, err := xdg.ConfigFile(authConfigFile)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(authFilepath); err != nil {
		return nil, err
	}
	data, err := os.ReadFile(authFilepath)
	if err != nil {
		return nil, err
	}
	var authFile authFile
	if err := json.Unmarshal(data, &authFile); err != nil {
		return nil, err
	}
	return &authFile, nil
}

func writeAuthFile(authFile *authFile) error {
	authFilepath, err := xdg.ConfigFile(authConfigFile)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(authFile, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(authFilepath, data, 0o600)
}

func removeAuthFile() error {
	authFilepath, err := xdg.ConfigFile(authConfigFile)
	if err != nil {
		return err
	}
	return os.RemoveAll(authFilepath)
}

func newAuthCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "auth"
	cmd.Short = "Authenticate with an rFMS API"
	cmd.AddCommand(newLoginCommand())
	cmd.AddCommand(newLogoutCommand())
	return cmd
}

func newLoginCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "login"
	cmd.Short = "Login to an rFMS API"
	cmd.AddCommand(newLoginScaniaCommand())
	cmd.AddCommand(newLoginVolvoTrucksCommand())
	// TODO: Implement login with unknown provider.
	return cmd
}

func newLoginScaniaCommand() *cobra.Command {
	const provider = rfms.Scania
	cmd := newCommand()
	cmd.Use = "scania"
	cmd.Short = "Login to Scania's rFMS API"
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
		auth := &authFile{
			Provider:         provider,
			TokenCredentials: &credentials,
		}
		if err := writeAuthFile(auth); err != nil {
			return err
		}
		cmd.Printf("Logged in to %s.", provider)
		return nil
	}
	return cmd
}

func newLoginVolvoTrucksCommand() *cobra.Command {
	const provider = rfms.VolvoTrucks
	cmd := newCommand()
	cmd.Use = "volvo-trucks"
	cmd.Short = "Login to Volvo Trucks' rFMS API"
	username := cmd.Flags().String("username", "", "username to use for authentication")
	cmd.MarkFlagRequired("username")
	password := cmd.Flags().String("password", "", "password to use for authentication")
	cmd.MarkFlagRequired("password")
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		auth := &authFile{
			Provider: provider,
			Username: *username,
			Password: *password,
		}
		if err := writeAuthFile(auth); err != nil {
			return err
		}
		cmd.Printf("Logged in to %s.", provider)
		return nil
	}
	return cmd
}

func newLogoutCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "logout"
	cmd.Short = "Logout from the current authenticated rFMS API"
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		if err := removeAuthFile(); err != nil {
			return err
		}
		cmd.Println("Logged out.")
		return nil
	}
	return cmd
}

func newClient() (*rfms.Client, error) {
	auth, err := readAuthFile()
	if err != nil {
		return nil, err
	}
	switch auth.Provider {
	case rfms.Scania:
		if auth.TokenCredentials.TokenExpireTime.Before(time.Now()) {
			if auth.TokenCredentials.RefreshTokenExpireTime.Before(time.Now()) {
				// TODO: Refresh token credentials if they are expired.
				_ = "TODO"
			}
			return nil, fmt.Errorf("session expired - please login again")
		}
		return rfms.NewClient(
			rfms.ScaniaBaseURL,
			rfms.WithReuseTokenAuth(*auth.TokenCredentials),
		), nil
	case rfms.VolvoTrucks:
		return rfms.NewClient(
			rfms.VolvoTrucksBaseURL,
			rfms.WithBasicAuth(auth.Username, auth.Password),
		), nil
	default:
		return nil, fmt.Errorf("unknown provider: %s", auth.Provider)
	}
}
