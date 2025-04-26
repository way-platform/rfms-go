module github.com/way-platform/rfms-go/cmd/rfmsctl/v4

go 1.24.1

require (
	github.com/spf13/cobra v1.9.1
	github.com/wayplatform/rfms-go/v4 v4.0.0-00010101000000-000000000000
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
)

replace github.com/way-platform/rfms-go/v4 => ../../
