package base

#GenerationTargetKind: "frame" | "index" | "rule"
#GenerationMaterializer: "gomplate" | "cue-export"
#GenerationFormat: "markdown" | "json" | "text"

#GenerationTarget: {
	name!: string
	kind!: #GenerationTargetKind
	input!: string
	schema?: string
	template?: string
	output!: string
	format!: #GenerationFormat
	materializer!: #GenerationMaterializer
	edit_policy!: "never-hand-edit"
}
