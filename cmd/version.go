package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of kcsi",
	Long:  `Display the current version of kcsi`,
	Run:   runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(_ *cobra.Command, _ []string) {
	fmt.Printf("kcsi version %s\n", version.GetVersion())
}
