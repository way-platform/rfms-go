module github.com/way-platform/rfms-go/cmd/rfms

go 1.25.0

toolchain go1.26.0

require (
	charm.land/fang/v2 v2.0.1
	charm.land/lipgloss/v2 v2.0.2
	github.com/adrg/xdg v0.5.3
	github.com/way-platform/rfms-go v0.14.0
	github.com/way-platform/rfms-go/cli v0.0.0
)

require (
	github.com/charmbracelet/colorprofile v0.4.2 // indirect
	github.com/charmbracelet/ultraviolet v0.0.0-20260205113103-524a6607adb8 // indirect
	github.com/charmbracelet/x/ansi v0.11.6 // indirect
	github.com/charmbracelet/x/exp/charmtone v0.0.0-20250911160549-0e720abcae8b // indirect
	github.com/charmbracelet/x/exp/color v0.0.0-20250911160549-0e720abcae8b // indirect
	github.com/charmbracelet/x/term v0.2.2 // indirect
	github.com/charmbracelet/x/termios v0.1.1 // indirect
	github.com/charmbracelet/x/windows v0.2.2 // indirect
	github.com/clipperhouse/displaywidth v0.11.0 // indirect
	github.com/clipperhouse/uax29/v2 v2.7.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.3.0 // indirect
	github.com/mattn/go-runewidth v0.0.20 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/mango v0.2.0 // indirect
	github.com/muesli/mango-cobra v1.3.0 // indirect
	github.com/muesli/mango-pflag v0.2.0 // indirect
	github.com/muesli/roff v0.1.0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/spf13/cobra v1.10.2 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	golang.org/x/oauth2 v0.31.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/term v0.41.0 // indirect
	golang.org/x/text v0.29.0 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace (
	github.com/way-platform/rfms-go => ../../
	github.com/way-platform/rfms-go/cli => ../../cli/
)
