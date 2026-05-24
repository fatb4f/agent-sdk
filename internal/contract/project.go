package contract

type ProjectGraph struct {
	SchemaVersion string                   `json:"schema_version"`
	SDK           *SDKConfig               `json:"sdk,omitempty"`
	Repo          RepoConfig               `json:"repo"`
	Component     *ComponentConfig         `json:"component,omitempty"`
	Output        OutputConfig             `json:"output"`
	Profile       ProfileConfig            `json:"profile"`
	Skills        []SkillConfig            `json:"skills"`
	Workflow      *WorkflowConfig          `json:"workflow,omitempty"`
	Surfaces      []SurfaceConfig          `json:"surfaces,omitempty"`
	Rules         []RuleConfig             `json:"rules,omitempty"`
	Adapters      map[string]AdapterConfig `json:"adapters"`
}

type SDKConfig struct {
	Name    string `json:"name"`
	Module  string `json:"module"`
	Version string `json:"version,omitempty"`
}

type RepoConfig struct {
	Name   string `json:"name"`
	Module string `json:"module"`
	Root   string `json:"root"`
}

type ComponentConfig struct {
	Name   string `json:"name,omitempty"`
	Module string `json:"module,omitempty"`
	Root   string `json:"root,omitempty"`
}

type OutputConfig struct {
	Root string `json:"root"`
}

type ProfileConfig struct {
	ID     string `json:"id"`
	Source string `json:"source"`
}

type SkillConfig struct {
	ID         string            `json:"id"`
	Required   bool              `json:"required,omitempty"`
	Path       string            `json:"path,omitempty"`
	Entrypoint string            `json:"entrypoint,omitempty"`
	Purpose    string            `json:"purpose,omitempty"`
	Status     string            `json:"status,omitempty"`
	LoadPolicy string            `json:"load_policy,omitempty"`
	Triggers   []string          `json:"triggers,omitempty"`
	Delegates  []string          `json:"delegates,omitempty"`
	Attrs      map[string]string `json:"attrs,omitempty"`
}

type WorkflowConfig struct {
	Name     string          `json:"name,omitempty"`
	Phases   []WorkflowPhase `json:"phases,omitempty"`
	Deferred []string        `json:"deferred,omitempty"`
}

type WorkflowPhase struct {
	ID                  string            `json:"id"`
	Tool                string            `json:"tool"`
	Mode                string            `json:"mode"`
	Command             []string          `json:"command,omitempty"`
	Env                 map[string]string `json:"env,omitempty"`
	MutatesSource       bool              `json:"mutates_source,omitempty"`
	After               string            `json:"after,omitempty"`
	BlocksOn            string            `json:"blocks_on,omitempty"`
	SourceMutationGuard bool              `json:"source_mutation_guard,omitempty"`
}

type SurfaceConfig struct {
	Name         string `json:"name,omitempty"`
	Kind         string `json:"kind"`
	Input        string `json:"input,omitempty"`
	Schema       string `json:"schema,omitempty"`
	Template     string `json:"template,omitempty"`
	Output       string `json:"output,omitempty"`
	Format       string `json:"format,omitempty"`
	Materializer string `json:"materializer,omitempty"`
	EditPolicy   string `json:"edit_policy,omitempty"`
	Path         string `json:"path,omitempty"`
	Source       string `json:"source,omitempty"`
}

type RuleConfig struct {
	Kind          string   `json:"kind,omitempty"`
	Pattern       []string `json:"pattern,omitempty"`
	Decision      string   `json:"decision,omitempty"`
	Justification string   `json:"justification,omitempty"`
}

type AdaptersConfig map[string]AdapterConfig

type AdapterConfig struct {
	Enabled bool `json:"enabled"`
}
