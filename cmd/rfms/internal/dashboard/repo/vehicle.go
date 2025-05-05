package repo

import (
	"context"

	"github.com/way-platform/rfms-go/v4"
	"github.com/way-platform/rfms-go/v4/api/rfmsv4"
)

type Vehicle struct {
	*rfmsv4.Vehicle
	LatestPosition *rfmsv4.VehiclePosition
	LatestStatus   *rfmsv4.VehicleStatus
}

func (r *Repo) LoadVehicles(ctx context.Context) ([]*Vehicle, error) {
	vehicles, err := r.loadAllVehicles(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*Vehicle, 0, len(vehicles))
	for _, vehicle := range vehicles {
		result = append(result, &Vehicle{
			Vehicle: vehicle,
		})
	}
	latestVehiclePositions, err := r.loadLatestVehiclePositions(ctx)
	if err != nil {
		return nil, err
	}
	for _, vehiclePosition := range latestVehiclePositions {
		for _, vehicle := range result {
			if vehicle.VIN == vehiclePosition.VIN {
				vehicle.LatestPosition = vehiclePosition
				break
			}
		}
	}
	latestVehicleStatuses, err := r.loadLatestVehicleStatuses(ctx)
	if err != nil {
		return nil, err
	}
	for _, vehicleStatus := range latestVehicleStatuses {
		for _, vehicle := range result {
			if vehicle.VIN == vehicleStatus.VIN {
				vehicle.LatestStatus = vehicleStatus
				break
			}
		}
	}
	return result, nil
}

func (r *Repo) loadAllVehicles(ctx context.Context) ([]*rfmsv4.Vehicle, error) {
	var result []*rfmsv4.Vehicle
	lastVIN, moreDataAvailable := "", true
	for moreDataAvailable {
		response, err := r.client.Vehicles(ctx, &rfms.VehiclesRequest{
			LastVIN: lastVIN,
		})
		if err != nil {
			return nil, err
		}
		moreDataAvailable = response.MoreDataAvailable
		for _, vehicle := range response.Vehicles {
			result = append(result, &vehicle)
		}
		lastVIN = response.Vehicles[len(response.Vehicles)-1].VIN
	}
	return result, nil
}

func (r *Repo) loadLatestVehiclePositions(ctx context.Context) ([]*rfmsv4.VehiclePosition, error) {
	var result []*rfmsv4.VehiclePosition
	lastVIN, moreDataAvailable := "", true
	for moreDataAvailable {
		response, err := r.client.VehiclePositions(ctx, &rfms.VehiclePositionsRequest{
			LatestOnly: true,
			LastVIN:    lastVIN,
		})
		if err != nil {
			return nil, err
		}
		moreDataAvailable = response.MoreDataAvailable
		for _, vehiclePosition := range response.VehiclePositions {
			result = append(result, &vehiclePosition)
		}
		lastVIN = response.VehiclePositions[len(response.VehiclePositions)-1].VIN
	}
	return result, nil
}

func (r *Repo) loadLatestVehicleStatuses(ctx context.Context) ([]*rfmsv4.VehicleStatus, error) {
	var result []*rfmsv4.VehicleStatus
	lastVIN, moreDataAvailable := "", true
	for moreDataAvailable {
		response, err := r.client.VehicleStatuses(ctx, &rfms.VehicleStatusesRequest{
			LatestOnly: true,
			LastVIN:    lastVIN,
		})
		if err != nil {
			return nil, err
		}
		moreDataAvailable = response.MoreDataAvailable
		for _, vehicleStatus := range response.VehicleStatuses {
			result = append(result, &vehicleStatus)
		}
		lastVIN = response.VehicleStatuses[len(response.VehicleStatuses)-1].VIN
	}
	return result, nil
}
