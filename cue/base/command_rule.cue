package base

#CommandRule: {
	kind!: "prefix" | "exact"
	pattern!: [...string]
	decision!: "allow" | "forbidden" | "review"
	justification!: string
}
