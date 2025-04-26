package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/way-platform/rfms-go/v4"
)

func main() {
	if err := newRootCommand().Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "rfmsctl"
	cmd.Short = "rFMS CLI"
	cmd.AddCommand(newAuthCommand())
	return cmd
}

func newAuthCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "auth"
	cmd.Short = "Authenticate with an rFMS API"
	cmd.AddCommand(newLoginCommand())
	return cmd
}

type authFile struct {
	Provider         string                `json:"provider"`
	TokenCredentials rfms.TokenCredentials `json:"tokenCredentials"`
}

func newLoginCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "login"
	cmd.Short = "Login to an rFMS API"
	provider := cmd.Flags().String("provider", "", "provider to login with")
	_ = cmd.MarkFlagRequired("provider")
	clientID := cmd.Flags().String("client-id", "", "client ID to use for authentication")
	_ = cmd.MarkFlagRequired("client-id")
	clientSecret := cmd.Flags().String("client-secret", "", "client secret to use for authentication")
	_ = cmd.MarkFlagRequired("client-secret")
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		cmd.Println(*provider)
		cmd.Println(*clientID)
		cmd.Println(*clientSecret)
		cmd.Println()
		authenticator := rfms.NewScaniaTokenAuthenticator(*clientID, *clientSecret)
		credentials, err := authenticator.Authenticate(cmd.Context())
		if err != nil {
			return err
		}
		cmd.Println(credentials)
		data, err := json.MarshalIndent(&authFile{
			Provider:         *provider,
			TokenCredentials: credentials,
		}, "", "  ")
		if err != nil {
			return err
		}
		cmd.Println(string(data))
		cmd.Println("Logged in.")
		return nil
	}
	return cmd
}

func newCommand() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	return cmd
}
