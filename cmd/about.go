package cmd

import (
	"fmt"

	"github.com/stanzinofree/kcsi/pkg/version"
	"github.com/spf13/cobra"
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

func runAbout(cmd *cobra.Command, args []string) {
	fmt.Println(version.GetAbout())
}
