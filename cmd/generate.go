/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
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
	Run: func(cmd *cobra.Command, args []string) {
		inputFlag, err := cmd.Flags().GetString("input")
		if err != nil {
			fmt.Printf("Error while invoking generate command: %v\n", err)
			return
		}

		if inputFlag == "" {
			fmt.Printf("Input flag is required\n")
			return
		}

		fmt.Printf("input: %v", inputFlag)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")
	generateCmd.PersistentFlags().StringP("input", "i", "", "Relative path to JSON file")
	generateCmd.PersistentFlags().StringP("package", "p", "com.example.model.rator", "Package declaration for generated files")
	generateCmd.PersistentFlags().StringP("lang", "l", "java", "Language in which the files should be generated")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
