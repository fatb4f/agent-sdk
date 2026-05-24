package bashcli

import "github.com/fatb4f/agent-sdk/cue/base"

workflow: base.#Workflow & {
	name: "bashly-source-edit"
	phases: [...base.#Phase] & [
		{
			id: "inspect"
			tool: "bashly"
			mode: "check"
			mutates_source: false
			after: ""
			blocks_on: ""
		},
		{
			id: "edit_source"
			tool: "bashly"
			mode: "write"
			mutates_source: true
			after: "inspect"
			blocks_on: "generated_bash_edited"
		},
		{
			id: "format_shellharden"
			tool: "shell-validation"
			mode: "write"
			mutates_source: true
			after: "edit_source"
			blocks_on: "shellharden_failed"
		},
		{
			id: "format_shfmt"
			tool: "shell-validation"
			mode: "write"
			mutates_source: true
			after: "format_shellharden"
			blocks_on: "shfmt_failed"
		},
		{
			id: "lint_source_shellcheck"
			tool: "shell-validation"
			mode: "check"
			after: "format_shfmt"
			mutates_source: false
			blocks_on: "shellcheck_source_failed"
		},
		{
			id: "generate_bashly"
			tool: "bashly"
			mode: "generate"
			command: ["bashly", "generate"]
			env: {
				BASHLY_FORMATTER: "none"
			}
			after: "lint_source_shellcheck"
			mutates_source: false
			blocks_on: "bashly_generate_failed"
			source_mutation_guard: true
		},
		{
			id: "report"
			tool: "report"
			mode: "check"
			after: "generate_bashly"
			mutates_source: false
			blocks_on: ""
		},
	]
	deferred: ["lint_generated", "bats-core", "shellspec"]

	adapters: {
		pre_commit: {
			kind: "local_hook"
			authority: "advisory"
			mode: "write"
			format_source: true
			generate_mutates_source: false
			report: {
				kind: "local_report"
			}
		}

		ci: {
			kind: "ci_job"
			authority: "blocking"
			mode: "check"
			format_source: false
			generate_mutates_source: false
			report: {
				kind: "ci_report"
			}
		}
	}

	gate: {
		kind: "ci_job"
		adapter: "ci"
		blocks_on: [
			"unsafe_write_target",
			"generated_bash_edited",
			"shellharden_failed",
			"shfmt_failed",
			"shellcheck_source_failed",
			"bashly_generate_failed",
		]
		deferred: ["lint_generated", "bats-core", "shellspec"]
	}
}
