package apiv2

//go:generate cp volvotrucks.original.yaml tmp.yaml
//go:generate sed -i s/com\.fms_standard\.rfms\.v2_1\.//g tmp.yaml
//go:generate sed -i s/AccumulatedType/Accumulated/g tmp.yaml
//go:generate sed -i s/AlternatorInfoType/AlternatorInfo/g tmp.yaml
//go:generate sed -i s/DoorStatusType/DoorStatus/g tmp.yaml
//go:generate sed -i s/DriverIdType/DriverId/g tmp.yaml
//go:generate sed -i s/FromToClassType/FromToClass/g tmp.yaml
//go:generate sed -i s/FromToClassesType/FromToClasses/g tmp.yaml
//go:generate sed -i s/GNSSPositionType/GNSSPosition/g tmp.yaml
//go:generate sed -i s/LabelClassType/LabelClass/g tmp.yaml
//go:generate sed -i s/LabelClassesType/LabelClasses/g tmp.yaml
//go:generate sed -i s/OemDriverIdentificationType/OemDriverIdentification/g tmp.yaml
//go:generate sed -i s/ProductionDateType/ProductionDate/g tmp.yaml
//go:generate sed -i s/SnapshotType/Snapshot/g tmp.yaml
//go:generate sed -i s/TachoDriverIdentificationType/TachoDriverIdentification/g tmp.yaml
//go:generate sed -i s/TellTaleInfoType/TellTaleInfo/g tmp.yaml
//go:generate sed -i s/TriggerType/Trigger/g tmp.yaml
//go:generate sed -i s/UptimeType/Uptime/g tmp.yaml
//go:generate sed -i s/VehiclePositionType/VehiclePosition/g tmp.yaml
//go:generate sed -i s/VehicleStatusType/VehicleStatus/g tmp.yaml
//go:generate sed -i s/VehicleType/Vehicle/g tmp.yaml
//go:generate npx swagger2openapi@v7.0.8 tmp.yaml --outfile api.yaml --yaml
//go:generate rm tmp.yaml
//go:generate go tool oapi-codegen -config cfg.yaml api.yaml
