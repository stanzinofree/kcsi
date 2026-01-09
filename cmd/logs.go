package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/completion"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
)

var logsCmd = &cobra.Command{
	Use:               "logs [pod-name]",
	Short:             "Get logs from a pod",
	Long:              `Get logs from a specific pod with namespace and pod name autocompletion`,
	Args:              cobra.ExactArgs(1),
	RunE:              runLogs,
	ValidArgsFunction: completion.PodCompletion,
}

var (
	logsNamespace string
	logsFollow    bool
	logsPrevious  bool
	logsTail      int64
	logsContainer string
)

func init() {
	rootCmd.AddCommand(logsCmd)

	// Add flags with autocompletion
	logsCmd.Flags().StringVarP(&logsNamespace, "namespace", "n", "", FlagDescNamespace)
	logsCmd.Flags().BoolVarP(&logsFollow, "follow", "f", false, "Follow log output")
	logsCmd.Flags().BoolVarP(&logsPrevious, "previous", "p", false, "Print the logs for the previous instance of the container")
	logsCmd.Flags().Int64Var(&logsTail, "tail", -1, "Lines of recent log file to display (default: all)")
	logsCmd.Flags().StringVarP(&logsContainer, "container", "c", "", "Container name (for multi-container pods)")

	logsCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)
	logsCmd.RegisterFlagCompletionFunc("container", completion.ContainerCompletion)
}

func runLogs(_ *cobra.Command, args []string) error {
	podName := args[0]

	// Build kubectl command with namespace injection
	kubectlArgs := kubernetes.BuildNamespaceArgs([]string{"logs"}, logsNamespace)
	kubectlArgs = append(kubectlArgs, podName)

	if logsFollow {
		kubectlArgs = append(kubectlArgs, "-f")
	}

	if logsPrevious {
		kubectlArgs = append(kubectlArgs, "-p")
	}

	if logsTail >= 0 {
		kubectlArgs = append(kubectlArgs, fmt.Sprintf("--tail=%d", logsTail))
	}

	if logsContainer != "" {
		kubectlArgs = append(kubectlArgs, "-c", logsContainer)
	}

	// Execute kubectl using the helper
	return kubernetes.ExecuteKubectlInteractive(kubectlArgs...)
}
