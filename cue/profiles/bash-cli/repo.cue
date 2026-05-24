package bashcli

import "github.com/fatb4f/agent-sdk/cue/base"

repo: {
	name: "bashly.sh"
	module: "github.com/fatb4f/bashly.sh"

	authority_contract: """
	CUE is authority.
	Skills are procedural manuals.
	Generated surfaces come later.
	Use CUE projections before filesystem crawling once available.
	"""
}

module: repo.module
name: repo.name
authority_contract: repo.authority_contract

boot_contract: """
## Repo-local discovery order

1. Read Codex-facing projections.
2. Read the skill index for inventory questions.
3. Read the workflow projection for validation and mutation policy.
4. Open `SKILL.md` only after selecting a relevant skill.
5. Use repo search only when generated projections lack the requested fact.
"""

repo_boot_projection: base.#Projection & {
	name: "repoFrame"
	format: "text"
	source: "agent-sdk/cue/profiles/bash-cli/repo.cue"
	value: boot_contract
}
