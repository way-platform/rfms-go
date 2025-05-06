# rFMS Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/way-platform/rfms-go)](https://pkg.go.dev/github.com/way-platform/rfms-go)
[![GoReportCard](https://goreportcard.com/badge/github.com/way-platform/rfms-go)](https://goreportcard.com/report/github.com/way-platform/rfms-go)
[![CI](https://github.com/way-platform/rfms-go/actions/workflows/ci.yaml/badge.svg)](https://github.com/way-platform/rfms-go/actions/workflows/ci.yaml)

A Go SDK and CLI tool for [rFMS](https://www.fms-standard.com/) APIs.

## SDK

### Features

* Full support for the [rFMS 4.0.0 standard](https://www.fms-standard.com/Truck/down_load/Technical_Specification_rFMS_vehicle_data_V4.0.0_17.09.2021.pdf)
* Backwards-compatibility with the [rFMS 2.1.1 standard](https://www.fms-standard.com/Truck/down_load/Technical_Specification_rFMS_V2.1.0_13.10.2017.pdf) (JSON)
* Support for vendor-specific fields from Scania

### Supported OEMs

* [Scania](https://developer.scania.com)
* [Volvo Trucks](https://developer.volvotrucks.com)

### Installing

```bash
$ go get github.com/way-platform/rfms-go@latest
```

### Using

```go
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
    moreDataAvailable = response.MoreDataAvailable
    if moreDataAvailable {
        lastVIN = response.Vehicles[len(response.Vehicles)-1].VIN
    }
}
```

### Developing

#### (re-)Generate OpenAPI models

```bash
$ go generate ./...
```

#### Run tests

```bash
$ go test ./...
```

## CLI tool

### Installing

```bash
go install github.com/way-platform/rfms-go/cmd/rfms@latest
```

### Using

```bash
$ rfms auth login scania --client-id $CLIENT_ID --client-secret $CLIENT_SECRET

Logged in to SCANIA.
```

```bash
$ rfms vehicles
{
  "authorizedPaths": [
    "/vehiclepositions",
    "/vehiclestatuses"
  ],
  "brand": "SCANIA",
  "customerVehicleName": "testingapp3",
  "type": "TRUCK",
  "vin": "YS2RKSTO01TT00068"
}
{
  "authorizedPaths": [
    "/vehiclepositions",
    "/vehiclestatuses"
  ],
  "brand": "SCANIA",
  "customerVehicleName": "Pins. Truck",
  "type": "TRUCK",
  "vin": "YS2R4X2000TT00396"
}
```
