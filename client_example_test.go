package rfms_test

import (
	"context"
	"fmt"
	"os"

	"github.com/way-platform/rfms-go/v4"
)

func ExampleClient_scania() {
	client := rfms.NewClient(
		rfms.WithScania(os.Getenv("SCANIA_CLIENT_ID"), os.Getenv("SCANIA_CLIENT_SECRET")),
	)
	lastVIN, moreDataAvailable := "", true
	for moreDataAvailable {
		response, err := client.Vehicles(context.Background(), &rfms.VehiclesRequest{
			LastVIN: lastVIN,
		})
		if err != nil {
			panic(err)
		}
		for _, vehicle := range response.Vehicles {
			fmt.Println(vehicle.VIN)
		}
		if response.MoreDataAvailable {
			lastVIN = response.Vehicles[len(response.Vehicles)-1].VIN
		}
	}
}

func ExampleClient_volvoTrucks() {
	client := rfms.NewClient(
		rfms.WithVolvoTrucks(os.Getenv("VOLVO_TRUCKS_USERNAME"), os.Getenv("VOLVO_TRUCKS_PASSWORD")),
	)
	lastVIN, moreDataAvailable := "", true
	for moreDataAvailable {
		response, err := client.Vehicles(context.Background(), &rfms.VehiclesRequest{
			LastVIN: lastVIN,
		})
		if err != nil {
			panic(err)
		}
		for _, vehicle := range response.Vehicles {
			fmt.Println(vehicle.VIN)
		}
		if response.MoreDataAvailable {
			lastVIN = response.Vehicles[len(response.Vehicles)-1].VIN
		}
	}
}
