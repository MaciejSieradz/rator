/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/MaciejSieradz/rator/internal/generator"
	"github.com/spf13/cobra"
)

type generateOptions struct {
	inputPath string
	pkgName   string
	language  string
}

func newGenerateCmd() *cobra.Command {
	opts := &generateOptions{}

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generates code from the JSON file/files",
		Long: `Generates code from JSON file or files. It can generate records, POJOs, or 
Kotlin data classes based on the flags.

Example Usage:
# Generates Java records
rator generate -i input.json -o com/example/destination
# Generates Kotlin mutable classes
rator generate -i input.json -l kotlin --mutable true
	`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.Validate()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.inputPath, "input", "i", "", "Relative path to JSON file")
	flags.StringVarP(&opts.pkgName, "package", "p", "com.example.model.rator", "Package declaration for generated files")
	flags.StringVarP(&opts.language, "lang", "l", "java", "Language in which the files should be generated")

	_ = cmd.MarkFlagRequired("input")

	return cmd
}

func (o *generateOptions) Validate() error {
	if _, err := os.Stat(o.inputPath); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist: %s", o.inputPath)
	}

	switch o.language {
	case "java", "kotlin":
		return nil
	default:
		return fmt.Errorf("unsupported language %q (must be 'java' or 'kotlin')", o.language)
	}
}

func (o *generateOptions) Run() error {
	gen, err := generator.New(o.inputPath, o.language, o.pkgName)
	if err != nil {
		return fmt.Errorf("failed to initialize generator: %w", err)
	}

	if err := gen.Generate(); err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(newGenerateCmd())
}
