module github.com/way-platform/rfms-go/cli

go 1.25.0

toolchain go1.26.0

require (
	github.com/spf13/cobra v1.10.2
	github.com/way-platform/rfms-go v0.14.0
	golang.org/x/oauth2 v0.31.0
	golang.org/x/term v0.41.0
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	golang.org/x/sys v0.42.0 // indirect
)

replace github.com/way-platform/rfms-go => ../
