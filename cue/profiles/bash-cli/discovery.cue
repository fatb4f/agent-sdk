package bashcli

import "github.com/fatb4f/agent-sdk/cue/base"

discovery: {
	boot: base.#DiscoveryRule & {
		kind: "boot"
		include: ["repoFrame"]
		search_allowed: false
	}

	inventory: base.#DiscoveryRule & {
		kind: "inventory"
		include: ["skillIndex", "skillFrame"]
		search_allowed: false
	}

	workflow: base.#DiscoveryRule & {
		kind: "workflow"
		include: ["workflowFrame", "commandRules"]
		search_allowed: true
	}

	fallback: base.#DiscoveryRule & {
		kind: "fallback"
		include: ["repoFrame", "skillIndex", "workflowFrame"]
		search_allowed: true
	}
}
