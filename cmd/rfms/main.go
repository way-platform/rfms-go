package main

import (
	"context"
	"image/color"
	"net/http"
	"os"

	"charm.land/fang/v2"
	"charm.land/lipgloss/v2"
	"github.com/adrg/xdg"
	rfms "github.com/way-platform/rfms-go"
	"github.com/way-platform/rfms-go/cli"
)

func main() {
	credPath, _ := xdg.ConfigFile("rfms-go/credentials.json")
	tokenPath, _ := xdg.ConfigFile("rfms-go/token.json")
	var debug bool
	debugTransport := &rfms.DebugTransport{Enabled: &debug}
	cmd := cli.NewCommand(
		cli.WithCredentialStore(cli.NewFileStore(credPath)),
		cli.WithTokenStore(cli.NewFileStore(tokenPath)),
		cli.WithHTTPClient(&http.Client{Transport: debugTransport}),
	)
	cmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging of HTTP requests")
	if err := fang.Execute(
		context.Background(),
		cmd,
		fang.WithColorSchemeFunc(func(c lipgloss.LightDarkFunc) fang.ColorScheme {
			base := c(lipgloss.Black, lipgloss.White)
			baseInverted := c(lipgloss.White, lipgloss.Black)
			return fang.ColorScheme{
				Base:         base,
				Title:        base,
				Description:  base,
				Comment:      base,
				Flag:         base,
				FlagDefault:  base,
				Command:      base,
				QuotedString: base,
				Argument:     base,
				Help:         base,
				Dash:         base,
				ErrorHeader:  [2]color.Color{baseInverted, base},
				ErrorDetails: base,
			}
		}),
	); err != nil {
		os.Exit(1)
	}
}
