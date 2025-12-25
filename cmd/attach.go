package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/completion"
)

var attachCmd = &cobra.Command{
	Use:   "attach [pod-name]",
	Short: "Attach to a pod with an interactive shell",
	Long: `Attach to a pod and start an interactive shell session.
Use -n to specify namespace first for better autocompletion.
Tries bash, zsh, and sh in order to find an available shell.

Examples:
  kcsi attach -n production my-pod
  kcsi attach -n production my-pod -c sidecar`,
	Args: cobra.ExactArgs(1),
	RunE: runAttach,
	ValidArgsFunction: completion.PodCompletion,
}

var (
	attachNamespace string
	attachContainer string
)

func runAttach(_ *cobra.Command, args []string) error {
	podName := args[0]

	// Shells to try in order of preference
	shells := []string{"bash", "zsh", "sh"}

	for _, shell := range shells {
		kubectlArgs := []string{"exec", "-it"}

		if attachNamespace != "" {
			kubectlArgs = append(kubectlArgs, "-n", attachNamespace)
		}

		kubectlArgs = append(kubectlArgs, podName)

		if attachContainer != "" {
			kubectlArgs = append(kubectlArgs, "-c", attachContainer)
		}

		kubectlArgs = append(kubectlArgs, "--", shell)

		fmt.Printf("Trying to attach with %s...\n", shell)

		kubectlCmd := exec.Command("kubectl", kubectlArgs...)
		kubectlCmd.Stdout = os.Stdout
		kubectlCmd.Stderr = os.Stderr
		kubectlCmd.Stdin = os.Stdin

		err := kubectlCmd.Run()
		if err == nil {
			// Successfully attached
			return nil
		}

		// If this shell didn't work, try the next one
		fmt.Printf("%s not available, trying next shell...\n", shell)
	}

	// Provide helpful error message
	if attachNamespace == "" {
		return fmt.Errorf("no interactive shell found in pod %s\nHint: Did you forget to specify the namespace with -n?", podName)
	}
	return fmt.Errorf("no interactive shell found in pod %s", podName)
}

func init() {
	rootCmd.AddCommand(attachCmd)

	attachCmd.Flags().StringVarP(&attachNamespace, "namespace", "n", "", FlagDescNamespace)
	attachCmd.Flags().StringVarP(&attachContainer, "container", "c", "", "Container name (for multi-container pods)")

	attachCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)
	attachCmd.RegisterFlagCompletionFunc("container", completion.ContainerCompletion)
}
