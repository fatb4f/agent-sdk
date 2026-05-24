package bashcli

import "github.com/fatb4f/agent-sdk/cue/base"

#SkillID: base.#SkillID

skills: {
	bashly: base.#Skill & {
		id: "bashly"
		path: ".agents/skills/bashly"
		entrypoint: ".agents/skills/bashly/SKILL.md"
		purpose: "Bashly source/config workflow and generated artifact boundary."
		required_tools: ["bashly"]
		optional_tools: ["ruby", "argc", "bash-ast", "tree-sitter", "sem"]
		triggers: ["bashly", "generated cli", "source script", "bashly.yml"]
		delegates: ["shell-validation", "bats-core", "shellspec"]
		status: "active"
	}

	argc: base.#Skill & {
		id: "argc"
		path: ".agents/skills/argc"
		entrypoint: ".agents/skills/argc/SKILL.md"
		purpose: "Argc annotation context and argv-guidance for shell source work."
		required_tools: ["argc"]
		triggers: ["argc", "argv context", "annotation context"]
		status: "active"
	}

	"bash-ast": base.#Skill & {
		id: "bash-ast"
		path: ".agents/skills/bash-ast"
		entrypoint: ".agents/skills/bash-ast/SKILL.md"
		purpose: "Bash AST evidence and structural inspection for shell source."
		optional_tools: ["bash-ast", "ast-bash"]
		triggers: ["bash-ast", "ast-bash", "parse tree", "structural evidence"]
		status: "active"
	}

	"shell-validation": base.#Skill & {
		id: "shell-validation"
		path: ".agents/skills/shell-validation"
		entrypoint: ".agents/skills/shell-validation/SKILL.md"
		purpose: "Shell normalization, formatting, linting, and local CI gate interpretation."
		required_tools: ["shellharden", "shfmt", "shellcheck"]
		triggers: ["shellharden", "shfmt", "shellcheck", "format", "lint"]
		status: "active"
	}

	"bats-core": base.#Skill & {
		id: "bats-core"
		path: ".agents/skills/bats-core"
		entrypoint: ".agents/skills/bats-core/SKILL.md"
		purpose: "Black-box Bash CLI behavior tests."
		optional_tools: ["bats"]
		status: "deferred"
		triggers: ["bats", "behavior tests"]
	}

	"shellspec": base.#Skill & {
		id: "shellspec"
		path: ".agents/skills/shellspec"
		entrypoint: ".agents/skills/shellspec/SKILL.md"
		purpose: "ShellSpec source-level Bash behavior tests."
		optional_tools: ["shellspec"]
		status: "deferred"
		triggers: ["shellspec", "behavior tests", "source-level tests"]
	}

	"tree-sitter": base.#Skill & {
		id: "tree-sitter"
		path: ".agents/skills/tree-sitter"
		entrypoint: ".agents/skills/tree-sitter/SKILL.md"
		purpose: "Structural parsing and grammar evidence for shell source."
		status: "experimental"
		triggers: ["tree-sitter", "parse tree", "structural evidence"]
	}
}

skillIndex: [...base.#SkillIndexEntry] & [
	{
		id: skills.bashly.id
		path: skills.bashly.path
		entrypoint: skills.bashly.entrypoint
		purpose: skills.bashly.purpose
		status: skills.bashly.status
		triggers: skills.bashly.triggers
	},
	{
		id: skills.argc.id
		path: skills.argc.path
		entrypoint: skills.argc.entrypoint
		purpose: skills.argc.purpose
		status: skills.argc.status
		triggers: skills.argc.triggers
	},
	{
		id: skills."bash-ast".id
		path: skills."bash-ast".path
		entrypoint: skills."bash-ast".entrypoint
		purpose: skills."bash-ast".purpose
		status: skills."bash-ast".status
		triggers: skills."bash-ast".triggers
	},
	{
		id: skills."shell-validation".id
		path: skills."shell-validation".path
		entrypoint: skills."shell-validation".entrypoint
		purpose: skills."shell-validation".purpose
		status: skills."shell-validation".status
		triggers: skills."shell-validation".triggers
	},
	{
		id: skills."bats-core".id
		path: skills."bats-core".path
		entrypoint: skills."bats-core".entrypoint
		purpose: skills."bats-core".purpose
		status: skills."bats-core".status
		triggers: skills."bats-core".triggers
	},
	{
		id: skills.shellspec.id
		path: skills.shellspec.path
		entrypoint: skills.shellspec.entrypoint
		purpose: skills.shellspec.purpose
		status: skills.shellspec.status
		triggers: skills.shellspec.triggers
	},
	{
		id: skills."tree-sitter".id
		path: skills."tree-sitter".path
		entrypoint: skills."tree-sitter".entrypoint
		purpose: skills."tree-sitter".purpose
		status: skills."tree-sitter".status
		triggers: skills."tree-sitter".triggers
	},
]
