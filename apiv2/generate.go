package apiv2

//go:generate npx swagger2openapi@v7.0.8 volvotrucks.yaml --outfile spec.yaml --yaml
//go:generate sed -i s/com\.fms_standard\.rfms\.v2_1\.//g spec.yaml
//go:generate go tool oapi-codegen -config cfg.yaml spec.yaml
