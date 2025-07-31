package rfmsv4

//go:generate cp rFMS_4_0.original.yaml api.yaml
//go:generate sed -i -f preprocess.sed api.yaml
//go:generate go tool -modfile ../../tools/go.mod oapi-codegen -config cfg.yaml api.yaml
//go:generate sed -i -f postprocess.sed api.gen.go
