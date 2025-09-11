package main

import (
	"context"
	"fmt"
	"image/color"
	"log/slog"
	"os"
	"time"

	"github.com/charmbracelet/fang"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/spf13/cobra"
	"github.com/way-platform/rfms-go"
	"github.com/way-platform/rfms-go/cmd/rfms/internal/auth"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	if err := fang.Execute(
		context.Background(),
		newRootCommand(),
		fang.WithColorSchemeFunc(func(c lipgloss.LightDarkFunc) fang.ColorScheme {
			base := c(lipgloss.Black, lipgloss.White)
			baseInverted := c(lipgloss.White, lipgloss.Black)
			return fang.ColorScheme{
				Base:         base,
				Title:        base,
				Description:  base,
				Comment:      base,
				Flag:         base,
				FlagDefault:  base,
				Command:      base,
				QuotedString: base,
				Argument:     base,
				Help:         base,
				Dash:         base,
				ErrorHeader:  [2]color.Color{baseInverted, base},
				ErrorDetails: base,
			}
		}),
	); err != nil {
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rfms",
		Short: "rFMS CLI",
	}
	cmd.AddGroup(&cobra.Group{
		ID:    "rfms",
		Title: "rFMS Commands",
	})
	cmd.AddCommand(newVehiclesCommand())
	cmd.AddCommand(newVehiclePositionsCommand())
	cmd.AddCommand(newVehicleStatusesCommand())
	cmd.AddGroup(&cobra.Group{
		ID:    "auth",
		Title: "Authentication",
	})
	cmd.AddCommand(auth.NewCommand())
	cmd.AddGroup(&cobra.Group{
		ID:    "utils",
		Title: "Utils",
	})
	cmd.SetHelpCommandGroupID("utils")
	cmd.SetCompletionCommandGroupID("utils")
	cmd.PersistentFlags().BoolP("debug", "d", false, "enable debug logging")
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		level := slog.LevelInfo
		if cmd.Flags().Changed("debug") {
			level = slog.LevelDebug
		}
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})))
		return nil
	}
	return cmd
}

func newVehiclesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "vehicles",
		Short:   "List vehicles",
		GroupID: "rfms",
	}
	limit := cmd.Flags().Int("limit", 100, "max vehicles queried")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := auth.NewClient()
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

func newVehiclePositionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "vehicle-positions",
		Short:   "List vehicle positions",
		GroupID: "rfms",
	}
	limit := cmd.Flags().Int("limit", 100, "max vehicle positions queried")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := auth.NewClient()
		if err != nil {
			return err
		}
		moreDataAvailable, lastVIN, count := true, "", 0
		for moreDataAvailable && count < *limit {
			response, err := client.VehiclePositions(cmd.Context(), rfms.VehiclePositionsRequest{
				LastVIN:    lastVIN,
				LatestOnly: true,
			})
			if err != nil {
				return err
			}
			for _, vehiclePosition := range response.VehiclePositions {
				fmt.Println(protojson.Format(vehiclePosition))
			}
			count += len(response.VehiclePositions)
			moreDataAvailable = response.MoreDataAvailable
			lastVIN = response.VehiclePositions[len(response.VehiclePositions)-1].GetVin()
		}
		return nil
	}
	return cmd
}

func newVehicleStatusesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "vehicle-statuses [VIN]",
		Short:   "List vehicle statuses",
		GroupID: "rfms",
		Args:    cobra.MaximumNArgs(1),
	}
	startTime := cmd.Flags().Time("start", time.Time{}, []string{time.DateOnly, time.RFC3339}, "start time")
	stopTime := cmd.Flags().Time("stop", time.Time{}, []string{time.DateOnly, time.RFC3339}, "stop time")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := auth.NewClient()
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
