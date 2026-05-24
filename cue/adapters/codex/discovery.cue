package codex

import (
	"strings"

	repo "github.com/fatb4f/agent-sdk/cue/profiles/bash-cli:bashcli"
)

discoveryFrame: strings.Join([
	"# Agent entrypoint",
	"",
	"This repository is CUE-authored.",
	"",
	"Start with `cue.mod/`.",
	"",
	"Do not treat `AGENTS.md` as policy authority.",
	"",
	"Authority graph: `agent.cue` / `agent-sdk/cue/profiles/bash-cli`",
	"Projection graph: `meta/agent/generated`",
	"Static check: `just agent-vet`",
	"",
	"Discovery contract:",
	repo.authority_contract,
	repo.boot_contract,
], "\n")
