package base

#SDKRef: {
	name!: string
	module!: string
	version?: string
}

#RepoConfig: {
	name!: string
	module!: string
	root!: string
}

#ComponentConfig: {
	name?: string
	module?: string
	root?: string
}

#OutputConfig: {
	root!: string
}

#ProfileConfig: {
	id!: string
	source!: string
}

#SkillConfig: {
	id!: string
	required?: bool
	path?: string
	entrypoint?: string
	purpose?: string
	status?: string
	load_policy?: string
	triggers?: [...string]
	delegates?: [...string]
	attrs?: [string]: string
}

#AdapterConfig: {
	enabled!: bool
}

#WorkflowPhase: {
	id!: string
	tool!: string
	mode!: "write" | "check" | "generate"
	command?: [...string]
	env?: [string]: string
	mutates_source?: bool
	after?: string
	blocks_on?: string
	source_mutation_guard?: bool
}

#Workflow: {
	name?: string
	phases?: [...#WorkflowPhase]
	deferred?: [...string]
}

#Surface: {
	name?: string
	kind!: "frame" | "index" | "rule"
	input?: string
	schema?: string
	template?: string
	output?: string
	format?: string
	materializer?: string
	edit_policy?: string
	path?: string
	source?: string
}

#Rule: {
	kind?: string
	pattern?: [...string]
	decision?: string
	justification?: string
}

#Project: {
	schema_version!: "agent-sdk/v0"
	sdk?: #SDKRef
	repo!: #RepoConfig
	component?: #ComponentConfig
	output!: #OutputConfig
	profile!: #ProfileConfig
	skills!: [...#SkillConfig]
	workflow?: #Workflow
	surfaces?: [...#Surface]
	rules?: [...#Rule]
	adapters!: [string]: #AdapterConfig
}
