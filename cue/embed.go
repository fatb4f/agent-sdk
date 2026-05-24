package cue

import "embed"

// FS embeds the CUE schemas, profiles, and adapters.
//
//go:embed base/*.cue profiles/**/*.cue adapters/**/*.cue
var FS embed.FS
