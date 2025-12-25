package cmd

import (
	"fmt"
	"os"

	"github.com/alessandro/kcsi/pkg/version"
	"github.com/spf13/cobra"
)

var (
	showDetailedVersion bool
)

var rootCmd = &cobra.Command{
	Use:   "kcsi",
	Short: "A kubectl wrapper with smart autocompletion",
	Long: `kcsi is a wrapper around kubectl that provides intelligent 
autocompletion for namespaces, pods, and other Kubernetes resources.`,
	Version: version.GetVersion(),
}

func Execute() {
	// Check for --version-detailed flag before cobra processes anything
	for _, arg := range os.Args {
		if arg == "--version-detailed" {
			fmt.Println(version.GetDetailedVersion())
			os.Exit(0)
		}
	}
	
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Custom version template with author info
	rootCmd.SetVersionTemplate(`{{with .Name}}{{printf "%s " .}}{{end}}{{printf "version %s" .Version}}
Author: Alessandro
{{if .Short}}{{.Short}}{{end}}
`)

	// Add detailed version flag - this needs to be handled before command execution
	rootCmd.PersistentFlags().BoolVar(&showDetailedVersion, "version-detailed", false, "Show detailed version information")
	
	// PersistentPreRun executes before any command
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if showDetailedVersion {
			fmt.Println(version.GetDetailedVersion())
			os.Exit(0)
		}
	}
}
