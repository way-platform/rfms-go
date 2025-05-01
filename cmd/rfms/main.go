package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/way-platform/rfms-go/cmd/rfms/v4/internal/auth"
	"github.com/way-platform/rfms-go/cmd/rfms/v4/internal/dashboard"
	"github.com/way-platform/rfms-go/v4"
)

func main() {
	if err := newRootCommand().Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rfms",
		Short: "rFMS CLI",
	}
	cmd.AddCommand(auth.NewCommand())
	cmd.AddCommand(dashboard.NewCommand())
	cmd.AddCommand(newVehiclesCommand())
	cmd.AddCommand(newVehiclePositionsCommand())
	cmd.AddCommand(newVehicleStatusesCommand())
	return cmd
}

func newVehiclesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vehicles",
		Short: "List vehicles",
	}
	limit := cmd.Flags().Int("limit", 100, "max vehicles queried")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := auth.NewClient()
		if err != nil {
			return err
		}
		moreDataAvailable, lastVIN, count := true, "", 0
		for moreDataAvailable && count < *limit {
			response, err := client.Vehicles(cmd.Context(), &rfms.VehiclesRequest{
				LastVIN: lastVIN,
			})
			if err != nil {
				return err
			}
			for _, vehicle := range response.Vehicles {
				printJSON(cmd, vehicle)
			}
			count += len(response.Vehicles)
			moreDataAvailable = response.MoreDataAvailable
			lastVIN = response.Vehicles[len(response.Vehicles)-1].VIN
		}
		return nil
	}
	return cmd
}

func newVehiclePositionsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vehicle-positions",
		Short: "List vehicle positions",
	}
	limit := cmd.Flags().Int("limit", 100, "max vehicle positions queried")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := auth.NewClient()
		if err != nil {
			return err
		}
		moreDataAvailable, lastVIN, count := true, "", 0
		for moreDataAvailable && count < *limit {
			response, err := client.VehiclePositions(cmd.Context(), &rfms.VehiclePositionsRequest{
				LastVIN:    lastVIN,
				LatestOnly: true,
			})
			if err != nil {
				return err
			}
			for _, vehiclePosition := range response.VehiclePositions {
				printJSON(cmd, vehiclePosition)
			}
			count += len(response.VehiclePositions)
			moreDataAvailable = response.MoreDataAvailable
			lastVIN = response.VehiclePositions[len(response.VehiclePositions)-1].VIN
		}
		return nil
	}
	return cmd
}

func newVehicleStatusesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vehicle-statuses",
		Short: "List vehicle statuses",
	}
	limit := cmd.Flags().Int("limit", 100, "max vehicle statuses queried")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := auth.NewClient()
		if err != nil {
			return err
		}
		moreDataAvailable, lastVIN, count := true, "", 0
		for moreDataAvailable && count < *limit {
			response, err := client.VehicleStatuses(cmd.Context(), &rfms.VehicleStatusesRequest{
				LastVIN:    lastVIN,
				LatestOnly: true,
			})
			if err != nil {
				return err
			}
			for _, vehicleStatus := range response.VehicleStatuses {
				printJSON(cmd, vehicleStatus)
			}
			count += len(response.VehicleStatuses)
			moreDataAvailable = response.MoreDataAvailable
			lastVIN = response.VehicleStatuses[len(response.VehicleStatuses)-1].VIN
		}
		return nil
	}
	return cmd
}

func printJSON(cmd *cobra.Command, msg any) error {
	data, err := json.MarshalIndent(msg, "", "  ")
	if err != nil {
		return err
	}
	cmd.Println(string(data))
	return nil
}
