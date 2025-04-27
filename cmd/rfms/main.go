package main

import (
	"bytes"
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
	return cmd
}

func newVehiclesCommand() *cobra.Command {
	cmd := newCommand()
	cmd.Use = "vehicles"
	cmd.Short = "List vehicles"
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		moreDataAvailable := true
		lastVIN := ""
		for moreDataAvailable {
			response, err := client.ListVehicles(cmd.Context(), &rfms.ListVehiclesRequest{
				LastVIN: lastVIN,
			})
			if err != nil {
				return err
			}
			for _, vehicle := range response.Vehicles {
				printRawJSON(cmd, vehicle.Raw)
			}
			moreDataAvailable = response.MoreDataAvailable
			lastVIN = response.Vehicles[len(response.Vehicles)-1].VIN
		}
		return nil
	}
	return cmd
}

func printRawJSON(cmd *cobra.Command, raw json.RawMessage) error {
	var buf bytes.Buffer
	json.Indent(&buf, raw, "", "  ")
	cmd.Println(buf.String())
	return nil
}
