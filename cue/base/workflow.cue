package base

#PhaseID: string
#PhaseMode: "write" | "check" | "generate"
#Blocker: string
#DeferredID: string

#AdapterID: string
#AdapterKind: "local_hook" | "ci_job"
#Authority: "advisory" | "blocking"
#MutationMode: "write" | "check"

#Workflow: {
	name!: string
	phases!: [...#Phase]
	deferred?: [...#DeferredID]
	adapters?: [string]: #Adapter
	gate?: #Gate
}

#Phase: {
	id!: #PhaseID
	tool!: string
	mode!: #PhaseMode
	command?: [...string]
	env?: [string]: string
	mutates_source: bool | *false
	after?: #PhaseID | ""
	blocks_on?: #Blocker | ""
	source_mutation_guard?: bool
}

#Adapter: {
	kind!: #AdapterKind
	authority!: #Authority
	mode!: #MutationMode
	format_source: bool | *false
	generate_mutates_source: bool | *false
	report?: {
		kind!: string
	}
}

#Gate: {
	kind!: "ci_job"
	adapter!: #AdapterID
	blocks_on?: [...#Blocker]
	deferred?: [...#DeferredID]
}
