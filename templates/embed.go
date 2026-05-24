package templates

import "embed"

// FS embeds the templates for frames, adapters, and rules.
//
//go:embed *.tmpl
var FS embed.FS
