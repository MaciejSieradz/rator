package generator

import (
	"fmt"
	"os"

	"github.com/MaciejSieradz/rator/internal/lexer"
	"github.com/MaciejSieradz/rator/internal/token"
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

	lexer := lexer.New(string(dat))
	var tokens []token.Token
	for tok := lexer.NextToken(); tok.Type != token.EOF; {
		tokens = append(tokens, tok)
		fmt.Println(tok)
		tok = lexer.NextToken()
	}

	return nil
}

func New(inputPath, language, pkg string) (Generator, error) {
	return &BasicGenerator{
		InputPath: inputPath,
		Language:  language,
		Pkg:       pkg,
	}, nil
}
