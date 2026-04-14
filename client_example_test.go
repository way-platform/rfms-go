package rfms_test

import (
	"context"
	"fmt"
	"os"

	rfms "github.com/way-platform/rfms-go"
	rfmsv5 "github.com/way-platform/rfms-go/proto/gen/go/wayplatform/connect/rfms/v5"
)

func ExampleClient_scania() {
	client, err := rfms.NewClient(
		rfms.WithScania(os.Getenv("SCANIA_CLIENT_ID"), os.Getenv("SCANIA_CLIENT_SECRET")),
	)
	if err != nil {
		panic(err)
	}
	request := rfmsv5.VehiclesRequest_builder{}.Build()
	for {
		response, err := client.Vehicles(context.Background(), request)
		if err != nil {
			panic(err)
		}
		for _, vehicle := range response.GetVehicles() {
			fmt.Println(vehicle.GetVin())
		}
		if !response.GetMoreDataAvailable() || len(response.GetVehicles()) == 0 {
			break
		}
		request.SetLastVin(response.GetVehicles()[len(response.GetVehicles())-1].GetVin())
	}
}

func ExampleClient_volvoTrucks() {
	client, err := rfms.NewClient(
		rfms.WithVolvoTrucks(
			os.Getenv("VOLVO_TRUCKS_USERNAME"),
			os.Getenv("VOLVO_TRUCKS_PASSWORD"),
		),
	)
	if err != nil {
		panic(err)
	}
	request := rfmsv5.VehiclesRequest_builder{}.Build()
	for {
		response, err := client.Vehicles(context.Background(), request)
		if err != nil {
			panic(err)
		}
		for _, vehicle := range response.GetVehicles() {
			fmt.Println(vehicle.GetVin())
		}
		if !response.GetMoreDataAvailable() || len(response.GetVehicles()) == 0 {
			break
		}
		request.SetLastVin(response.GetVehicles()[len(response.GetVehicles())-1].GetVin())
	}
}
