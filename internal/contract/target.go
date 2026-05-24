package contract

type GenerationTargetKind string

const (
	GenerationTargetKindFrame GenerationTargetKind = "frame"
	GenerationTargetKindIndex GenerationTargetKind = "index"
	GenerationTargetKindRule  GenerationTargetKind = "rule"
)

type GenerationMaterializer string

const (
	GenerationMaterializerGomplate  GenerationMaterializer = "gomplate"
	GenerationMaterializerCueExport GenerationMaterializer = "cue-export"
)

type GenerationFormat string

const (
	GenerationFormatMarkdown GenerationFormat = "markdown"
	GenerationFormatJSON     GenerationFormat = "json"
	GenerationFormatText     GenerationFormat = "text"
)

type GenerationTarget struct {
	Name         string                 `json:"name"`
	Kind         GenerationTargetKind   `json:"kind"`
	Input        string                 `json:"input"`
	Schema       string                 `json:"schema,omitempty"`
	Template     string                 `json:"template,omitempty"`
	Output       string                 `json:"output"`
	Format       GenerationFormat       `json:"format"`
	Materializer GenerationMaterializer `json:"materializer"`
	EditPolicy   string                 `json:"edit_policy"`
}
