package apiv2

//go:generate npx swagger2openapi@v7.0.8 volvotrucks.swagger2.yaml --outfile volvotrucks.oas3.yaml --yaml
//go:generate go tool oapi-codegen -config cfg.yaml volvotrucks.oas3.yaml
