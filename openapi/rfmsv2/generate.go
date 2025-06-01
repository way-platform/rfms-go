package rfmsv2

//go:generate cp volvotrucks.original.yaml tmp.yaml
//go:generate sed -i -f preprocess.sed tmp.yaml
//go:generate npx swagger2openapi@v7.0.8 tmp.yaml --outfile api.yaml --yaml
//go:generate rm tmp.yaml
//go:generate go tool oapi-codegen -config cfg.yaml api.yaml
//go:generate sed -i -f postprocess.sed api.gen.go
