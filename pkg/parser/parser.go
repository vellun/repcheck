package parser

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

type Parser interface {
	Parse(tmpDir string) (*modfile.File, error)
}

type GoModParser struct{}

func (g *GoModParser) Parse(tmpDir string) (*modfile.File, error) {
	goModPath := filepath.Join(tmpDir, "go.mod")
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("No go.mod file in repo")
	}

	data, err := os.ReadFile(goModPath)
	if err != nil {
		return nil, fmt.Errorf("Error while reading go.mod file: %v", err)
	}

	file, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing go.mod: %v", err)
	}
	return file, nil
}
