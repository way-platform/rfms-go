package rfmsv4

//go:generate cp rFMS_4_0.original.yaml api.yaml
//go:generate go tool oapi-codegen -config cfg.yaml api.yaml
