package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/completion"
)

var executeCmd = &cobra.Command{
	Use:   "execute [pod-name] -- [command...]",
	Short: "Execute a command in a pod",
	Long: `Execute a specific command in a pod.
Use -- to separate the pod name from the command to execute.

Examples:
  kcsi execute my-pod -n production -- ls -la
  kcsi execute my-pod -c sidecar -- cat /etc/hosts`,
	Args: cobra.MinimumNArgs(1),
	RunE: runExecute,
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

	kubectlArgs := []string{"exec", podName}

	if executeNamespace != "" {
		kubectlArgs = append(kubectlArgs, "-n", executeNamespace)
	}

	if executeContainer != "" {
		kubectlArgs = append(kubectlArgs, "-c", executeContainer)
	}

	kubectlArgs = append(kubectlArgs, "--")
	kubectlArgs = append(kubectlArgs, command...)

	kubectlCmd := exec.Command("kubectl", kubectlArgs...)
	kubectlCmd.Stdout = os.Stdout
	kubectlCmd.Stderr = os.Stderr
	kubectlCmd.Stdin = os.Stdin

	if err := kubectlCmd.Run(); err != nil {
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
