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

func newCommand() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	return cmd
}

func newRootCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "rfms"
	cmd.Short = "rFMS CLI"
	cmd.AddCommand(newAuthCommand())
	cmd.AddCommand(newVehiclesCommand())
	cmd.AddCommand(newVehiclePositionsCommand())
	return cmd
}

func newVehiclesCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "vehicles"
	cmd.Short = " vehicles"
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		moreDataAvailable := true
		lastVIN := ""
		for moreDataAvailable {
			response, err := client.Vehicles(cmd.Context(), &rfms.VehiclesRequest{
				LastVIN: lastVIN,
			})
			if err != nil {
				return err
			}
			for _, vehicle := range response.Vehicles {
				printJSON(cmd, vehicle)
			}
			moreDataAvailable = response.MoreDataAvailable
			lastVIN = response.Vehicles[len(response.Vehicles)-1].VIN
		}
		return nil
	}
	return cmd
}

func newVehiclePositionsCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "vehicle-positions"
	cmd.Short = " vehicle positions"
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		moreDataAvailable := true
		lastVIN := ""
		for moreDataAvailable {
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
			moreDataAvailable = response.MoreDataAvailable
			lastVIN = response.VehiclePositions[len(response.VehiclePositions)-1].VIN
		}
		return nil
	}
	return cmd
}

func newVehicleStatusesCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "vehicle-statuses"
	cmd.Short = " vehicle statuses"
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		moreDataAvailable := true
		lastVIN := ""
		for moreDataAvailable {
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
