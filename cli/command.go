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
		creds := &ScaniaCredentials{}
		if cfg.scaniaCredentialStore != nil {
			if loaded, err := cfg.scaniaCredentialStore.Load(); err != nil && !errors.Is(err, fs.ErrNotExist) {
				return fmt.Errorf("read credentials: %w", err)
			} else if err == nil {
				creds = loaded
			}
		}
		// Override with flags.
		if *clientID != "" {
			creds.ClientID = *clientID
		}
		if *clientSecret != "" {
			creds.ClientSecret = *clientSecret
		}
		// Prompt for missing fields.
		if creds.ClientID == "" {
			val, err := promptSecret(cmd, "Enter OAuth2 client ID: ")
			if err != nil {
				return fmt.Errorf("read client ID: %w", err)
			}
			creds.ClientID = val
		}
		if creds.ClientSecret == "" {
			val, err := promptSecret(cmd, "Enter OAuth2 client secret: ")
			if err != nil {
				return fmt.Errorf("read client secret: %w", err)
			}
			creds.ClientSecret = val
		}
		// Persist credentials.
		if cfg.scaniaCredentialStore != nil {
			if err := cfg.scaniaCredentialStore.Save(creds); err != nil {
				return fmt.Errorf("write credentials: %w", err)
			}
		}
		// Clear Volvo credentials when switching provider.
		if cfg.volvoCredentialStore != nil {
			_ = cfg.volvoCredentialStore.Clear()
		}
		// Run OAuth2 flow.
		tokenSource := rfms.ScaniaAuthConfig{
			ClientID:     creds.ClientID,
			ClientSecret: creds.ClientSecret,
		}.TokenSource(context.Background())
		token, err := tokenSource.Token()
		if err != nil {
			return fmt.Errorf("get token: %w", err)
		}
		// Cache token.
		if cfg.tokenStore != nil {
			if err := cfg.tokenStore.Save(token); err != nil {
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
		creds := &VolvoCredentials{}
		if cfg.volvoCredentialStore != nil {
			if loaded, err := cfg.volvoCredentialStore.Load(); err != nil && !errors.Is(err, fs.ErrNotExist) {
				return fmt.Errorf("read credentials: %w", err)
			} else if err == nil {
				creds = loaded
			}
		}
		// Override with flags.
		if *username != "" {
			creds.Username = *username
		}
		if *password != "" {
			creds.Password = *password
		}
		// Prompt for missing fields.
		if creds.Username == "" {
			val, err := promptSecret(cmd, "Enter username: ")
			if err != nil {
				return fmt.Errorf("read username: %w", err)
			}
			creds.Username = val
		}
		if creds.Password == "" {
			val, err := promptSecret(cmd, "Enter password: ")
			if err != nil {
				return fmt.Errorf("read password: %w", err)
			}
			creds.Password = val
		}
		// Persist credentials.
		if cfg.volvoCredentialStore != nil {
			if err := cfg.volvoCredentialStore.Save(creds); err != nil {
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
			if cfg.volvoCredentialStore != nil {
				if err := cfg.volvoCredentialStore.Clear(); err != nil {
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
		if _, err := cfg.scaniaCredentialStore.Load(); err == nil {
			if cfg.tokenStore != nil {
				token, err := cfg.tokenStore.Load()
				if err != nil {
					if errors.Is(err, fs.ErrNotExist) {
						return nil, fmt.Errorf(
							"session expired, please login again using `rfms auth login scania`",
						)
					}
					return nil, fmt.Errorf("read token: %w", err)
				}
				if !token.Valid() {
					return nil, fmt.Errorf(
						"session expired, please login again using `rfms auth login scania`",
					)
				}
				opts = append(opts,
					rfms.WithBaseURL(rfms.ScaniaBaseURL),
					rfms.WithVersion(rfms.V4),
					rfms.WithTokenSource(oauth2.StaticTokenSource(token)),
				)
			}
			return rfms.NewClient(opts...)
		}
	}
	// Try Volvo Trucks credentials.
	if cfg.volvoCredentialStore != nil {
		if creds, err := cfg.volvoCredentialStore.Load(); err == nil {
			opts = append(opts,
				rfms.WithVolvoTrucks(creds.Username, creds.Password),
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
