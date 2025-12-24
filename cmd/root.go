package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kcsi",
	Short: "A kubectl wrapper with smart autocompletion",
	Long: `kcsi is a wrapper around kubectl that provides intelligent 
autocompletion for namespaces, pods, and other Kubernetes resources.`,
	Version: "0.1.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can go here
}
