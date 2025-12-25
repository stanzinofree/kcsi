package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/completion"
)

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Get cluster events",
	Long:  `Get Kubernetes events with namespace filtering and autocompletion`,
	RunE:  runEvents,
}

var (
	eventsNamespace string
	eventsWatch     bool
)

func init() {
	rootCmd.AddCommand(eventsCmd)

	eventsCmd.Flags().StringVarP(&eventsNamespace, "namespace", "n", "", "Kubernetes namespace (all namespaces if not specified)")
	eventsCmd.Flags().BoolVarP(&eventsWatch, "watch", "w", false, "Watch for events")
	eventsCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)
}

func runEvents(cmd *cobra.Command, args []string) error {
	kubectlArgs := []string{"get", "events"}

	if eventsNamespace != "" {
		kubectlArgs = append(kubectlArgs, "-n", eventsNamespace)
	} else {
		kubectlArgs = append(kubectlArgs, "--all-namespaces")
	}

	if eventsWatch {
		kubectlArgs = append(kubectlArgs, "--watch")
	}

	// Sort by timestamp for better readability
	kubectlArgs = append(kubectlArgs, "--sort-by=.lastTimestamp")

	kubectlCmd := exec.Command("kubectl", kubectlArgs...)
	kubectlCmd.Stdout = os.Stdout
	kubectlCmd.Stderr = os.Stderr
	kubectlCmd.Stdin = os.Stdin

	if err := kubectlCmd.Run(); err != nil {
		return fmt.Errorf("failed to execute kubectl: %w", err)
	}

	return nil
}
