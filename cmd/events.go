package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/completion"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
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

func runEvents(_ *cobra.Command, _ []string) error {
	kubectlArgs := []string{"get", "events"}

	// Use namespace injection, but fallback to --all-namespaces if no namespace is specified
	effectiveNS := kubernetes.InjectDefaultNamespace(eventsNamespace)
	if effectiveNS != "" {
		kubectlArgs = append(kubectlArgs, "-n", effectiveNS)
	} else {
		kubectlArgs = append(kubectlArgs, flagAllNamespaces)
	}

	if eventsWatch {
		kubectlArgs = append(kubectlArgs, "--watch")
	}

	// Sort by timestamp for better readability
	kubectlArgs = append(kubectlArgs, "--sort-by=.lastTimestamp")

	// Execute kubectl using the helper
	return kubernetes.ExecuteKubectlInteractive(kubectlArgs...)
}
