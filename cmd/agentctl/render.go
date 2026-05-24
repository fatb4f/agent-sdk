package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	sdkcue "github.com/fatb4f/agent-sdk/cue"
	"github.com/fatb4f/agent-sdk/internal/cuegraph"
	sdktpl "github.com/fatb4f/agent-sdk/templates"
)

type codexGenerationData struct {
	Repo struct {
		Name              string `json:"name"`
		Module            string `json:"module"`
		AuthorityContract string `json:"authority_contract"`
		BootContract      string `json:"boot_contract"`
	} `json:"repo"`
	RepoFrame  string `json:"repoFrame"`
	SkillFrame struct {
		SourceOfTruth       string          `json:"source_of_truth"`
		InventoryProjection string          `json:"inventory_projection"`
		DiscoveryRule       string          `json:"discovery_rule"`
		Rows                []codexSkillRow `json:"rows"`
	} `json:"skillFrame"`
	WorkflowFrame struct {
		SourceOfTruth   string               `json:"source_of_truth"`
		ValidationOrder string               `json:"validation_order"`
		Deferred        []string             `json:"deferred"`
		Phases          []codexWorkflowPhase `json:"phases"`
		GenerateBashly  codexWorkflowPhase   `json:"generate_bashly"`
	} `json:"workflowFrame"`
	CommandRules []codexCommandRule `json:"commandRules"`
}

type codexWorkflowPhase struct {
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

type codexCommandRule struct {
	Kind          string   `json:"kind"`
	Pattern       []string `json:"pattern"`
	Decision      string   `json:"decision"`
	Justification string   `json:"justification"`
}

type codexSkillRow struct {
	ID         string `json:"id"`
	Status     string `json:"status"`
	Path       string `json:"path"`
	Purpose    string `json:"purpose"`
	LoadPolicy string `json:"load_policy,omitempty"`
}

type codexSkillIndexEntry struct {
	ID         string   `json:"id"`
	Path       string   `json:"path"`
	Entrypoint string   `json:"entrypoint"`
	Purpose    string   `json:"purpose"`
	Status     string   `json:"status"`
	LoadPolicy string   `json:"load_policy,omitempty"`
	Triggers   []string `json:"triggers,omitempty"`
}

type codexSurfaceIndexEntry struct {
	Path       string `json:"path"`
	Source     string `json:"source"`
	EditPolicy string `json:"edit_policy"`
	Kind       string `json:"kind"`
}

func generateAgentSurfaces(projectRoot, outputRoot string) error {
	sdkRoot := filepath.Join(projectRoot, "agent-sdk")
	graph, err := cuegraph.New(sdkRoot, sdkcue.FS)
	if err != nil {
		return err
	}

	var generation codexGenerationData
	if err := graph.Decode("cue/adapters/codex", "generationData", &generation); err != nil {
		return err
	}

	var skillIndex []codexSkillIndexEntry
	if err := graph.Decode("cue/adapters/codex", "skillIndex", &skillIndex); err != nil {
		return err
	}

	var surfaceIndex []codexSurfaceIndexEntry
	if err := graph.Decode("cue/adapters/codex", "surfaceIndex", &surfaceIndex); err != nil {
		return err
	}

	return writeCodexOutputs(outputRoot, generation, skillIndex, surfaceIndex)
}

func writeCodexOutputs(outputRoot string, generation codexGenerationData, skillIndex []codexSkillIndexEntry, surfaceIndex []codexSurfaceIndexEntry) error {
	framesDir := filepath.Join(outputRoot, "meta", "agent", "frames")
	generatedDir := filepath.Join(outputRoot, "meta", "agent", "generated")
	rulesDir := filepath.Join(outputRoot, "meta", "agent", "rules")

	for _, dir := range []string{framesDir, generatedDir, rulesDir} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	repoFrame, err := renderTemplate("repo-frame.md.tmpl", map[string]any{
		"repo": map[string]any{
			"name":               generation.Repo.Name,
			"module":             generation.Repo.Module,
			"authority_contract": generation.Repo.AuthorityContract,
			"boot_contract":      generation.Repo.BootContract,
		},
	})
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(framesDir, "repo-frame.md"), []byte(repoFrame), 0o644); err != nil {
		return err
	}

	skillFrame, err := renderTemplate("skills.md.tmpl", map[string]any{
		"skillFrame": map[string]any{
			"source_of_truth":      generation.SkillFrame.SourceOfTruth,
			"inventory_projection": generation.SkillFrame.InventoryProjection,
			"discovery_rule":       generation.SkillFrame.DiscoveryRule,
			"rows":                 skillRowsToMaps(generation.SkillFrame.Rows),
		},
	})
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(framesDir, "skills.md"), []byte(skillFrame), 0o644); err != nil {
		return err
	}

	workflowFrame, err := renderTemplate("workflow.md.tmpl", map[string]any{
		"workflowFrame": map[string]any{
			"source_of_truth":  generation.WorkflowFrame.SourceOfTruth,
			"validation_order": generation.WorkflowFrame.ValidationOrder,
			"deferred":         generation.WorkflowFrame.Deferred,
			"phases":           workflowPhasesToMaps(generation.WorkflowFrame.Phases),
			"generate_bashly":  workflowPhaseToMap(generation.WorkflowFrame.GenerateBashly),
		},
	})
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(framesDir, "workflow.md"), []byte(workflowFrame), 0o644); err != nil {
		return err
	}

	if err := writeJSON(filepath.Join(generatedDir, "skill-index.json"), skillIndex); err != nil {
		return err
	}
	if err := writeJSON(filepath.Join(generatedDir, "surface-index.json"), surfaceIndex); err != nil {
		return err
	}

	rules, err := renderTemplate("default.rules.tmpl", map[string]any{
		"commandRules": commandRulesToMaps(generation.CommandRules),
	})
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(rulesDir, "default.rules"), []byte(rules), 0o644); err != nil {
		return err
	}

	return nil
}

func renderTemplate(name string, data any) (string, error) {
	tpl, err := template.ParseFS(sdktpl.FS, name)
	if err != nil {
		return "", fmt.Errorf("parse %s: %w", name, err)
	}

	var out bytes.Buffer
	if err := tpl.Execute(&out, data); err != nil {
		return "", fmt.Errorf("render %s: %w", name, err)
	}

	return out.String(), nil
}

func writeJSON(path string, value any) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(value)
}

func skillRowsToMaps(rows []codexSkillRow) []map[string]any {
	out := make([]map[string]any, 0, len(rows))
	for _, row := range rows {
		out = append(out, map[string]any{
			"id":          row.ID,
			"status":      row.Status,
			"path":        row.Path,
			"purpose":     row.Purpose,
			"load_policy": row.LoadPolicy,
		})
	}
	return out
}

func workflowPhasesToMaps(phases []codexWorkflowPhase) []map[string]any {
	out := make([]map[string]any, 0, len(phases))
	for _, phase := range phases {
		out = append(out, workflowPhaseToMap(phase))
	}
	return out
}

func workflowPhaseToMap(phase codexWorkflowPhase) map[string]any {
	return map[string]any{
		"id":                    phase.ID,
		"tool":                  phase.Tool,
		"mode":                  phase.Mode,
		"command":               phase.Command,
		"env":                   phase.Env,
		"mutates_source":        phase.MutatesSource,
		"after":                 phase.After,
		"blocks_on":             phase.BlocksOn,
		"source_mutation_guard": phase.SourceMutationGuard,
	}
}

func commandRulesToMaps(rules []codexCommandRule) []map[string]any {
	out := make([]map[string]any, 0, len(rules))
	for _, rule := range rules {
		out = append(out, map[string]any{
			"kind":          rule.Kind,
			"pattern":       rule.Pattern,
			"decision":      rule.Decision,
			"justification": rule.Justification,
		})
	}
	return out
}
