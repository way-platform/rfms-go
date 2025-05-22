package proto

//go:generate openapi-overlay apply overlay.yaml api.yaml > tmp.yaml
//go:generate java -jar openapi-generator-cli-7.13.0.jar generate -g protobuf-schema -i tmp.yaml --additional-properties=addJsonNameAnnotation=true,aggregateModelsName=api.proto
//go:generate buf generate
