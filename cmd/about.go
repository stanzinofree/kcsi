package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/version"
)

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "About kcsi - project information and philosophy",
	Long:  `Display detailed information about kcsi, including its spirit, philosophy, and author information`,
	Run:   runAbout,
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}

func runAbout(_ *cobra.Command, _ []string) {
	fmt.Println(version.GetAbout())
}
