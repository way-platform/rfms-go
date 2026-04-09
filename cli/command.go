package cli

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/spf13/cobra"
	rfms "github.com/way-platform/rfms-go"
	scaniacreds "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/scania/rfms/v1"
	volvocreds "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/volvotrucks/rfms/v1"
	"golang.org/x/oauth2"
	"golang.org/x/term"
	"google.golang.org/protobuf/encoding/protojson"
)

// NewCommand builds the full CLI command tree for the rFMS SDK.
func NewCommand(opts ...Option) *cobra.Command {
	cfg := config{}
	for _, opt := range opts {
		opt(&cfg)
	}
	cmd := &cobra.Command{
		Use:   "rfms",
		Short: "rFMS CLI",
	}
	cmd.AddGroup(&cobra.Group{ID: "rfms", Title: "rFMS Commands"})
	cmd.AddCommand(newVehiclesCommand(&cfg))
	cmd.AddCommand(newVehiclePositionsCommand(&cfg))
	cmd.AddCommand(newVehicleStatusesCommand(&cfg))
	cmd.AddGroup(&cobra.Group{ID: "auth", Title: "Authentication"})
	cmd.AddCommand(newAuthCommand(&cfg))
	cmd.AddGroup(&cobra.Group{ID: "utils", Title: "Utils"})
	cmd.SetHelpCommandGroupID("utils")
	cmd.SetCompletionCommandGroupID("utils")
	return cmd
}

func newAuthCommand(cfg *config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "auth",
		Short:   "Authenticate with an rFMS API",
		GroupID: "auth",
	}
	cmd.AddCommand(newLoginCommand(cfg))
	cmd.AddCommand(newLogoutCommand(cfg))
	return cmd
}

func newLoginCommand(cfg *config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to an rFMS API",
	}
	cmd.AddCommand(newLoginScaniaCommand(cfg))
	cmd.AddCommand(newLoginVolvoTrucksCommand(cfg))
	return cmd
}

func newLoginScaniaCommand(cfg *config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scania",
		Short: "Login to the Scania rFMS API",
	}

	clientID := cmd.Flags().String("client-id", "", "client ID for authentication")
	clientSecret := cmd.Flags().String("client-secret", "", "client secret for authentication")

	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		// Try loading stored credentials first.
		creds := &scaniacreds.Credentials{}
		if cfg.scaniaCredentialStore != nil {
			if err := cfg.scaniaCredentialStore.Read(creds); err != nil && !errors.Is(err, fs.ErrNotExist) {
				return fmt.Errorf("read credentials: %w", err)
			}
		}
		// Override with flags.
		if *clientID != "" {
			creds.SetClientId(*clientID)
		}
		if *clientSecret != "" {
			creds.SetClientSecret(*clientSecret)
		}
		// Prompt for missing fields.
		if creds.GetClientId() == "" {
			val, err := promptSecret(cmd, "Enter OAuth2 client ID: ")
			if err != nil {
				return fmt.Errorf("read client ID: %w", err)
			}
			creds.SetClientId(val)
		}
		if creds.GetClientSecret() == "" {
			val, err := promptSecret(cmd, "Enter OAuth2 client secret: ")
			if err != nil {
				return fmt.Errorf("read client secret: %w", err)
			}
			creds.SetClientSecret(val)
		}
		// Persist credentials.
		if cfg.scaniaCredentialStore != nil {
			if err := cfg.scaniaCredentialStore.Write(creds); err != nil {
				return fmt.Errorf("write credentials: %w", err)
			}
		}
		// Clear Volvo credentials when switching provider.
		if cfg.volvoTrucksCredentialStore != nil {
			_ = cfg.volvoTrucksCredentialStore.Clear()
		}
		// Run OAuth2 flow.
		tokenSource := rfms.ScaniaAuthConfig{
			ClientID:     creds.GetClientId(),
			ClientSecret: creds.GetClientSecret(),
		}.TokenSource(context.Background())
		token, err := tokenSource.Token()
		if err != nil {
			return fmt.Errorf("get token: %w", err)
		}
		// Cache token.
		if cfg.tokenStore != nil {
			if err := cfg.tokenStore.Write(token); err != nil {
				return fmt.Errorf("write token: %w", err)
			}
		}
		cmd.Printf("Logged in to %s.", rfms.BrandScania)
		return nil
	}
	return cmd
}

func newLoginVolvoTrucksCommand(cfg *config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "volvo-trucks",
		Short: "Login to the Volvo Trucks rFMS API",
	}

	username := cmd.Flags().String("username", "", "username for authentication")
	password := cmd.Flags().String("password", "", "password for authentication")

	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		// Try loading stored credentials first.
		creds := &volvocreds.Credentials{}
		if cfg.volvoTrucksCredentialStore != nil {
			if err := cfg.volvoTrucksCredentialStore.Read(creds); err != nil && !errors.Is(err, fs.ErrNotExist) {
				return fmt.Errorf("read credentials: %w", err)
			}
		}
		// Override with flags.
		if *username != "" {
			creds.SetUsername(*username)
		}
		if *password != "" {
			creds.SetPassword(*password)
		}
		// Prompt for missing fields.
		if creds.GetUsername() == "" {
			val, err := promptSecret(cmd, "Enter username: ")
			if err != nil {
				return fmt.Errorf("read username: %w", err)
			}
			creds.SetUsername(val)
		}
		if creds.GetPassword() == "" {
			val, err := promptSecret(cmd, "Enter password: ")
			if err != nil {
				return fmt.Errorf("read password: %w", err)
			}
			creds.SetPassword(val)
		}
		// Persist credentials.
		if cfg.volvoTrucksCredentialStore != nil {
			if err := cfg.volvoTrucksCredentialStore.Write(creds); err != nil {
				return fmt.Errorf("write credentials: %w", err)
			}
		}
		// Clear Scania credentials when switching provider.
		if cfg.scaniaCredentialStore != nil {
			_ = cfg.scaniaCredentialStore.Clear()
		}
		cmd.Printf("Logged in to %s.", rfms.BrandVolvoTrucks)
		return nil
	}
	return cmd
}

func newLogoutCommand(cfg *config) *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Logout from the current authenticated rFMS API",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if cfg.tokenStore != nil {
				if err := cfg.tokenStore.Clear(); err != nil {
					return fmt.Errorf("clear token: %w", err)
				}
			}
			if cfg.scaniaCredentialStore != nil {
				if err := cfg.scaniaCredentialStore.Clear(); err != nil {
					return fmt.Errorf("clear scania credentials: %w", err)
				}
			}
			if cfg.volvoTrucksCredentialStore != nil {
				if err := cfg.volvoTrucksCredentialStore.Clear(); err != nil {
					return fmt.Errorf("clear volvo credentials: %w", err)
				}
			}
			cmd.Println("Logged out.")
			return nil
		},
	}
}

func newVehiclesCommand(cfg *config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "vehicles",
		Short:   "List vehicles",
		GroupID: "rfms",
	}
	limit := cmd.Flags().Int("limit", 100, "max vehicles queried")
	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		client, err := newClient(cmd, cfg)
		if err != nil {
			return err
		}
		moreDataAvailable, lastVIN, count := true, "", 0
		for moreDataAvailable && count < *limit {
			response, err := client.Vehicles(cmd.Context(), rfms.VehiclesRequest{
				LastVIN: lastVIN,
			})
			if err != nil {
				return err
			}
			for _, vehicle := range response.Vehicles {
				fmt.Println(protojson.Format(vehicle))
			}
			count += len(response.Vehicles)
			moreDataAvailable = response.MoreDataAvailable
			lastVIN = response.Vehicles[len(response.Vehicles)-1].GetVin()
		}
		return nil
	}
	return cmd
}

func newVehiclePositionsCommand(cfg *config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "vehicle-positions [VIN]",
		Short:   "List vehicle positions",
		GroupID: "rfms",
		Args:    cobra.MaximumNArgs(1),
	}
	startTime := cmd.Flags().
		Time("start", time.Time{}, []string{time.DateOnly, time.RFC3339}, "start time")
	stopTime := cmd.Flags().
		Time("stop", time.Time{}, []string{time.DateOnly, time.RFC3339}, "stop time")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := newClient(cmd, cfg)
		if err != nil {
			return err
		}
		var vin string
		if len(args) > 0 {
			vin = args[0]
		}
		moreDataAvailable, lastVIN := true, ""
		for moreDataAvailable {
			request := rfms.VehiclePositionsRequest{
				LastVIN:    lastVIN,
				VIN:        vin,
				LatestOnly: startTime.IsZero() && stopTime.IsZero(),
				StartTime:  *startTime,
				StopTime:   *stopTime,
			}
			response, err := client.VehiclePositions(cmd.Context(), request)
			if err != nil {
				return err
			}
			for _, vehiclePosition := range response.VehiclePositions {
				fmt.Println(protojson.Format(vehiclePosition))
			}
			moreDataAvailable = response.MoreDataAvailable
			if !moreDataAvailable {
				break
			}
			lastVIN = response.VehiclePositions[len(response.VehiclePositions)-1].GetVin()
		}
		return nil
	}
	return cmd
}

func newVehicleStatusesCommand(cfg *config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "vehicle-statuses [VIN]",
		Short:   "List vehicle statuses",
		GroupID: "rfms",
		Args:    cobra.MaximumNArgs(1),
	}
	startTime := cmd.Flags().
		Time("start", time.Time{}, []string{time.DateOnly, time.RFC3339}, "start time")
	stopTime := cmd.Flags().
		Time("stop", time.Time{}, []string{time.DateOnly, time.RFC3339}, "stop time")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := newClient(cmd, cfg)
		if err != nil {
			return err
		}
		var vin string
		if len(args) > 0 {
			vin = args[0]
		}
		moreDataAvailable, lastVIN := true, ""
		for moreDataAvailable {
			request := rfms.VehicleStatusesRequest{
				LastVIN:    lastVIN,
				VIN:        vin,
				LatestOnly: startTime.IsZero() && stopTime.IsZero(),
				StartTime:  *startTime,
				StopTime:   *stopTime,
			}
			response, err := client.VehicleStatuses(cmd.Context(), request)
			if err != nil {
				return err
			}
			for _, vehicleStatus := range response.VehicleStatuses {
				fmt.Println(protojson.Format(vehicleStatus))
			}
			moreDataAvailable = response.MoreDataAvailable
			if !moreDataAvailable {
				break
			}
			lastVIN = response.VehicleStatuses[len(response.VehicleStatuses)-1].GetVin()
		}
		return nil
	}
	return cmd
}

func newClient(_ *cobra.Command, cfg *config) (*rfms.Client, error) {
	var opts []rfms.ClientOption
	if cfg.httpClient != nil {
		opts = append(opts, rfms.WithHTTPClient(cfg.httpClient))
	}
	// Try Scania credentials.
	if cfg.scaniaCredentialStore != nil {
		creds := &scaniacreds.Credentials{}
		if err := cfg.scaniaCredentialStore.Read(creds); err == nil {
			var token oauth2.Token
			if cfg.tokenStore != nil {
				if err := cfg.tokenStore.Read(&token); err != nil {
					if errors.Is(err, fs.ErrNotExist) {
						return nil, fmt.Errorf(
							"session expired, please login again using `rfms auth login scania`",
						)
					}
					return nil, fmt.Errorf("read token: %w", err)
				}
			}
			if !token.Valid() {
				return nil, fmt.Errorf(
					"session expired, please login again using `rfms auth login scania`",
				)
			}
			opts = append(opts,
				rfms.WithBaseURL(rfms.ScaniaBaseURL),
				rfms.WithVersion(rfms.V4),
				rfms.WithTokenSource(oauth2.StaticTokenSource(&token)),
			)
			return rfms.NewClient(opts...)
		}
	}
	// Try Volvo Trucks credentials.
	if cfg.volvoTrucksCredentialStore != nil {
		creds := &volvocreds.Credentials{}
		if err := cfg.volvoTrucksCredentialStore.Read(creds); err == nil {
			opts = append(opts,
				rfms.WithVolvoTrucks(creds.GetUsername(), creds.GetPassword()),
			)
			return rfms.NewClient(opts...)
		}
	}
	return nil, fmt.Errorf("no credentials found, please login using `rfms auth login`")
}

func promptSecret(cmd *cobra.Command, prompt string) (string, error) {
	cmd.Print(prompt)
	input, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	cmd.Println()
	return string(input), nil
}
