package codex

import repopkg "github.com/fatb4f/agent-sdk/cue/profiles/bash-cli:bashcli"

generationData: {
	repo: {
		name: repopkg.name
		module: repopkg.module
		authority_contract: repopkg.authority_contract
		boot_contract: repopkg.boot_contract
	}

	repoFrame: repopkg.repo_boot_projection.value

	skillFrame: {
		source_of_truth: "agent-sdk/cue/profiles/bash-cli/skills.cue"
		inventory_projection: "meta/agent/generated/skill-index.json"
		discovery_rule: "answer inventory questions from this frame or `skill-index.json` before opening `SKILL.md` files."
		rows: [
			for id in skillOrder {
				let s = repopkg.skills[id]
				{
					id: s.id
					status: s.status
					path: s.path
					purpose: s.purpose
					load_policy: s.load_policy
				}
			},
		]
	}

	workflowFrame: {
		source_of_truth: "agent-sdk/cue/profiles/bash-cli/workflow.cue"
		validation_order: "shellharden -> shfmt -> shellcheck source -> bashly generate with Bashly formatting disabled -> report."
		deferred: repopkg.workflow.deferred
		phases: repopkg.workflow.phases
		generate_bashly: repopkg.workflow.phases[5]
	}

	commandRules: [
		for r in repopkg.command_rules {
			{
				kind: r.kind
				pattern: r.pattern
				decision: r.decision
				justification: r.justification
			}
		},
	]
}
