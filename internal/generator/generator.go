package generator

import (
	"fmt"
	"os"
)

type Generator interface {
	Generate() error
}

type BasicGenerator struct {
	InputPath string
	Language  string
	Pkg       string
}

func (g *BasicGenerator) Generate() error {
	dat, err := os.ReadFile(g.InputPath)
	if err != nil {
		return fmt.Errorf("error reading file %q: %w", g.InputPath, err)
	}

	fmt.Println(string(dat))
	return nil
}

func New(inputPath, language, pkg string) (Generator, error) {
	return &BasicGenerator{
		InputPath: inputPath,
		Language:  language,
		Pkg:       pkg,
	}, nil
}
