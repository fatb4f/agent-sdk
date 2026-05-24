package bashcli

import "github.com/fatb4f/agent-sdk/cue/base"

commandRules: [
	{
		kind: "prefix"
		pattern: ["cue", "vet"]
		decision: "allow"
		justification: "CUE validation is repo-local and non-mutating."
	},
	{
		kind: "prefix"
		pattern: ["cue", "export"]
		decision: "allow"
		justification: "Exporting projections is read-only and part of the static authority slice."
	},
	{
		kind: "prefix"
		pattern: ["rm", "-rf"]
		decision: "forbidden"
		justification: "Destructive delete must not be agent-default."
	},
]

command_rules: [...base.#CommandRule] & commandRules
