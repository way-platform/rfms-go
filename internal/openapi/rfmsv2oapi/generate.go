package rfmsv2oapi

//go:generate echo [rfmsv2oapi] copying original...
//go:generate cp volvotrucks.yaml 01-original.yaml

//go:generate echo [rfmsv2oapi] preprocessing...
//go:generate sh -c "sed -f preprocess.sed 01-original.yaml > 02-preprocessed.yaml"

//go:generate echo [rfmsv2oapi] converting from Swagger to OpenAPI...
//go:generate npx swagger2openapi@v7.0.8 02-preprocessed.yaml --outfile 03-converted.yaml --yaml

//go:generate echo [rfmsv2oapi] applying overlay...
//go:generate sh -c "openapi-overlay apply overlay.yaml 03-converted.yaml > 04-overlayed.yaml"

//go:generate echo [rfmsv2oapi] generating code...
//go:generate oapi-codegen -config config.yaml 04-overlayed.yaml

//go:generate echo [rfmsv2oapi] OK
