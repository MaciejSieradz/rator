/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rator",
	Short: "A generator of type-safe JVM models from JSON payloads",
	Long: `rator is a developer-friendly command-line utility designed to accelerate JVM backend development.
By analyzing raw JSON structures, rator automatically infers types and generates robust, 
type-safe Java Records, classic POJOs, or Kotlin Data Classes.

Equipped with support for modern industry standards like JSpecify nullability annotations 
and popular serialization frameworks (Jackson, kotlinx.serialization), rator handles 
complex, nested API structures locally, safely, and instantaneously.

Example Usage:
  # Generate Kotlin data classes with kotlinx.serialization annotations
  rator generate -i response.json -l kotlin -f kotlinx

  # Generate Java Records with JSpecify null-safety annotations
  rator generate -i user.json -l java --nullable`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
