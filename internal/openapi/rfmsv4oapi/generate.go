package rfmsv4oapi

//go:generate echo [rfmsv4oapi] copying original...
//go:generate cp rFMS_4_0.yaml 01-original.yaml

//go:generate echo [rfmsv4oapi] applying overlay...
//go:generate sh -c "go tool -modfile ../../../tools/go.mod openapi-overlay apply overlay.yaml 01-original.yaml > 02-overlayed.yaml"

//go:generate echo [rfmsv4oapi] generating code...
//go:generate go tool -modfile ../../../tools/go.mod oapi-codegen -config config.yaml 02-overlayed.yaml

//go:generate echo [rfmsv4oapi] OK
