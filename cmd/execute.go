package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/completion"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
)

var executeCmd = &cobra.Command{
	Use:   "execute [pod-name] -- [command...]",
	Short: "Execute a command in a pod",
	Long: `Execute a specific command in a pod.
Use -n to specify namespace first for better autocompletion.
Use -- to separate the pod name from the command to execute.

Examples:
  kcsi execute -n production my-pod -- ls -la
  kcsi execute -n production my-pod -c sidecar -- cat /etc/hosts
  kcsi execute -n default api-pod -- curl localhost:8080/health`,
	Args:              cobra.MinimumNArgs(1),
	RunE:              runExecute,
	ValidArgsFunction: completion.PodCompletion,
}

var (
	executeNamespace string
	executeContainer string
)

func runExecute(_ *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("command is required after pod name (use -- to separate)")
	}

	podName := args[0]
	command := args[1:]

	// Build kubectl command with namespace injection
	kubectlArgs := kubernetes.BuildNamespaceArgs([]string{"exec"}, executeNamespace)
	kubectlArgs = append(kubectlArgs, podName)

	if executeContainer != "" {
		kubectlArgs = append(kubectlArgs, "-c", executeContainer)
	}

	kubectlArgs = append(kubectlArgs, "--")
	kubectlArgs = append(kubectlArgs, command...)

	if err := kubernetes.ExecuteKubectlInteractive(kubectlArgs...); err != nil {
		// Provide helpful error message if pod not found
		effectiveNS := kubernetes.InjectDefaultNamespace(executeNamespace)
		if effectiveNS == "" {
			return fmt.Errorf("failed to execute command: %w\nHint: Did you forget to specify the namespace with -n?", err)
		}
		return fmt.Errorf("failed to execute command: %w", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(executeCmd)

	executeCmd.Flags().StringVarP(&executeNamespace, "namespace", "n", "", FlagDescNamespace)
	executeCmd.Flags().StringVarP(&executeContainer, "container", "c", "", "Container name (for multi-container pods)")

	executeCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)
	executeCmd.RegisterFlagCompletionFunc("container", completion.ContainerCompletion)
}
