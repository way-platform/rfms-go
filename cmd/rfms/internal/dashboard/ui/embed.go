package ui

import "embed"

//go:embed templates
var TemplateFS embed.FS

//go:embed static
var StaticFS embed.FS
