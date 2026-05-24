package codex

import repo "github.com/fatb4f/agent-sdk/cue/profiles/bash-cli:bashcli"

skillOrder: [
	"argc",
	"bash-ast",
	"bashly",
	"bats-core",
	"shell-validation",
	"shellspec",
	"tree-sitter",
]

skillIndex: [
	for id in skillOrder {
		let s = repo.skills[id]
		{
			id: s.id
			path: s.path
			entrypoint: s.entrypoint
			purpose: s.purpose
			status: s.status
			load_policy: s.load_policy
			triggers: s.triggers
		}
	}
]

surfaceIndex: repo.generated_surfaces
