package bashcli

import "github.com/fatb4f/agent-sdk/cue/base"

surfaces: {
	generation_targets: [...base.#GenerationTarget] & [
		{
			name: "repo-frame"
			kind: "frame"
			input: "generationData.repoFrame"
			template: "agent-sdk/templates/repo-frame.md.tmpl"
			output: "meta/agent/frames/repo-frame.md"
			format: "markdown"
			materializer: "gomplate"
			edit_policy: "never-hand-edit"
		},
		{
			name: "skills-frame"
			kind: "frame"
			input: "generationData.skillFrame"
			template: "agent-sdk/templates/skills.md.tmpl"
			output: "meta/agent/frames/skills.md"
			format: "markdown"
			materializer: "gomplate"
			edit_policy: "never-hand-edit"
		},
		{
			name: "workflow-frame"
			kind: "frame"
			input: "generationData.workflowFrame"
			template: "agent-sdk/templates/workflow.md.tmpl"
			output: "meta/agent/frames/workflow.md"
			format: "markdown"
			materializer: "gomplate"
			edit_policy: "never-hand-edit"
		},
		{
			name: "skill-index"
			kind: "index"
			input: "skillIndex"
			output: "meta/agent/generated/skill-index.json"
			format: "json"
			materializer: "cue-export"
			edit_policy: "never-hand-edit"
		},
		{
			name: "surface-index"
			kind: "index"
			input: "surfaceIndex"
			output: "meta/agent/generated/surface-index.json"
			format: "json"
			materializer: "cue-export"
			edit_policy: "never-hand-edit"
		},
		{
			name: "default-rules"
			kind: "rule"
			input: "generationData.commandRules"
			template: "agent-sdk/templates/default.rules.tmpl"
			output: "meta/agent/rules/default.rules"
			format: "text"
			materializer: "gomplate"
			edit_policy: "never-hand-edit"
		},
	]

	future_frames: [
		{
			path: "meta/agent/frames/repo-frame.md"
			source: "agent-sdk/templates/repo-frame.md.tmpl"
			edit_policy: "never-hand-edit"
			kind: "frame"
		},
		{
			path: "meta/agent/frames/skills.md"
			source: "agent-sdk/templates/skills.md.tmpl"
			edit_policy: "never-hand-edit"
			kind: "frame"
		},
		{
			path: "meta/agent/frames/workflow.md"
			source: "agent-sdk/templates/workflow.md.tmpl"
			edit_policy: "never-hand-edit"
			kind: "frame"
		},
	]

	future_indexes: [
		{
			path: "meta/agent/generated/skill-index.json"
			source: "agent-sdk/cue/adapters/codex/indexes.cue"
			edit_policy: "never-hand-edit"
			kind: "index"
		},
		{
			path: "meta/agent/generated/surface-index.json"
			source: "agent-sdk/cue/adapters/codex/indexes.cue"
			edit_policy: "never-hand-edit"
			kind: "index"
		},
	]

	future_rules: [
		{
			path: "meta/agent/rules/default.rules"
			source: "agent-sdk/templates/default.rules.tmpl"
			edit_policy: "never-hand-edit"
			kind: "rule"
		},
	]
}

generated_surfaces: [...base.#GeneratedSurface] & [
	for s in surfaces.future_frames {
		s
	},
	for s in surfaces.future_indexes {
		s
	},
	for s in surfaces.future_rules {
		s
	},
]
