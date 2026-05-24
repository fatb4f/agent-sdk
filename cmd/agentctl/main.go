package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	sdkcue "github.com/fatb4f/agent-sdk/cue"
	"github.com/fatb4f/agent-sdk/internal/contract"
	"github.com/fatb4f/agent-sdk/internal/cuegraph"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	if len(args) == 0 {
		usage()
		return 2
	}

	switch args[0] {
	case "generate":
		return runGenerate(args[1:])
	case "check-generated":
		return runCheckGenerated(args[1:])
	case "vet":
		return runVet(args[1:])
	case "doctor":
		return runDoctor(args[1:])
	case "targets":
		return runTargets(args[1:])
	case "-h", "--help", "help":
		usage()
		return 0
	default:
		fmt.Fprintf(os.Stderr, "agentctl: unknown command: %s\n", args[0])
		usage()
		return 2
	}
}

func runGenerate(args []string) int {
	fs := flag.NewFlagSet("generate", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	root := fs.String("project-root", ".", "project root")
	outputRoot := fs.String("output-root", "", "output root; defaults to project root")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	if *outputRoot == "" {
		*outputRoot = *root
	}

	if err := generateProjectGraph(*root, *outputRoot); err != nil {
		fmt.Fprintf(os.Stderr, "agentctl generate: %v\n", err)
		return 1
	}

	if err := generateAgentSurfaces(*root, *outputRoot); err != nil {
		fmt.Fprintf(os.Stderr, "agentctl generate: %v\n", err)
		return 1
	}

	return 0
}

func generateProjectGraph(projectRoot, outputRoot string) error {
	graph, err := cuegraph.New(projectRoot, sdkcue.FS)
	if err != nil {
		return err
	}

	var project contract.ProjectGraph
	if err := graph.DecodeValidated("agent.cue", "project", "base/project.cue", "#Project", &project); err != nil {
		return err
	}

	generatedDir := filepath.Join(outputRoot, project.Output.Root, "generated")
	if err := os.MkdirAll(generatedDir, 0o755); err != nil {
		return err
	}

	graphPath := filepath.Join(generatedDir, "project-graph.json")
	f, err := os.Create(graphPath)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(project); err != nil {
		return err
	}

	fmt.Printf("agentctl generate: wrote %s\n", graphPath)
	return nil
}

func runCheckGenerated(args []string) int {
	fs := flag.NewFlagSet("check-generated", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	root := fs.String("project-root", ".", "project root")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	tempDir, err := os.MkdirTemp("", "agentctl-generated-*")
	if err != nil {
		fmt.Fprintf(os.Stderr, "agentctl check-generated: %v\n", err)
		return 1
	}
	defer os.RemoveAll(tempDir)

	if err := generateProjectGraph(*root, tempDir); err != nil {
		fmt.Fprintf(os.Stderr, "agentctl check-generated: %v\n", err)
		return 1
	}

	if err := generateAgentSurfaces(*root, tempDir); err != nil {
		fmt.Fprintf(os.Stderr, "agentctl check-generated: %v\n", err)
		return 1
	}

	expected := filepath.Join(*root, "meta", "agent")
	actual := filepath.Join(tempDir, "meta", "agent")
	if err := compareTrees(expected, actual); err != nil {
		fmt.Fprintf(os.Stderr, "agentctl check-generated: %v\n", err)
		return 1
	}

	fmt.Println("agentctl check-generated: ok")
	return 0
}

func runVet(args []string) int {
	fs := flag.NewFlagSet("vet", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	root := fs.String("project-root", ".", "project root")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	graph, err := cuegraph.New(*root, sdkcue.FS)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	if err := graph.Vet("agent.cue"); err != nil {
		fmt.Fprintf(os.Stderr, "agentctl vet: %v\n", err)
		return 1
	}

	fmt.Println("agentctl vet: ok")
	return 0
}

func runDoctor(args []string) int {
	fmt.Println("agentctl doctor: ok")
	return 0
}

func runTargets(args []string) int {
	fs := flag.NewFlagSet("targets", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	root := fs.String("project-root", ".", "project root")
	pkg := fs.String("package", "agent.cue", "CUE package to evaluate")
	expr := fs.String("expr", "surfaces.generation_targets", "CUE expression to decode")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	graph, err := cuegraph.New(*root, sdkcue.FS)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	var targets []contract.GenerationTarget
	if err := graph.Decode(*pkg, *expr, &targets); err != nil {
		fmt.Fprintf(os.Stderr, "agentctl targets: %v\n", err)
		return 1
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(targets); err != nil {
		fmt.Fprintf(os.Stderr, "agentctl targets: %v\n", err)
		return 1
	}

	return 0
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: agentctl <generate|check-generated|vet|doctor|targets> [--project-root PATH]")
}

func compareTrees(expectedRoot, actualRoot string) error {
	expected := make(map[string]string)
	if err := filepath.WalkDir(expectedRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(expectedRoot, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		if d.IsDir() {
			expected[rel] = "dir"
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		expected[rel] = string(data)
		return nil
	}); err != nil {
		return err
	}

	actual := make(map[string]string)
	if err := filepath.WalkDir(actualRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(actualRoot, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		if d.IsDir() {
			actual[rel] = "dir"
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		actual[rel] = string(data)
		return nil
	}); err != nil {
		return err
	}

	var paths []string
	for path := range expected {
		paths = append(paths, path)
	}
	for path := range actual {
		if _, ok := expected[path]; !ok {
			paths = append(paths, path)
		}
	}
	sort.Strings(paths)

	seen := make(map[string]struct{})
	var diffs []string
	for _, path := range paths {
		if _, ok := seen[path]; ok {
			continue
		}
		seen[path] = struct{}{}
		want, wok := expected[path]
		got, gok := actual[path]
		switch {
		case !wok:
			diffs = append(diffs, fmt.Sprintf("unexpected generated path: %s", path))
		case !gok:
			diffs = append(diffs, fmt.Sprintf("missing generated path: %s", path))
		case want != got:
			diffs = append(diffs, fmt.Sprintf("drift in %s", path))
		}
	}

	if len(diffs) > 0 {
		return errors.New(strings.Join(diffs, "; "))
	}

	return nil
}
