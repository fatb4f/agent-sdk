package base

#SkillID: string
#SkillStatus: "active" | "deferred" | "experimental" | "deprecated"
#SkillLoadPolicy: "on_select" | "always" | "never"

#Skill: {
	id!: #SkillID
	path!: string
	entrypoint!: string
	purpose!: string

	required_tools?: [...string]
	optional_tools?: [...string]
	triggers?: [...string]
	delegates?: [...#SkillID]

	status: *"active" | #SkillStatus
	load_policy: *"on_select" | #SkillLoadPolicy
}

#SkillIndexEntry: {
	id!: #SkillID
	path!: string
	entrypoint!: string
	purpose!: string
	status!: #SkillStatus
	load_policy?: #SkillLoadPolicy
	triggers?: [...string]
	required_tools?: [...string]
	optional_tools?: [...string]
}
