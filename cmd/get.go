package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alessandro/kcsi/pkg/completion"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Kubernetes resources",
	Long:  `Get Kubernetes resources with smart autocompletion`,
}

var getPodsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Get pods in a namespace",
	Long:  `Get pods in a specific namespace with autocompletion support`,
	RunE:  runGetPods,
}

var namespace string

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getPodsCmd)

	// Add namespace flag with autocompletion
	getPodsCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Kubernetes namespace")
	getPodsCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)
}

func runGetPods(cmd *cobra.Command, args []string) error {
	// Build kubectl command
	kubectlArgs := []string{"get", "pods"}
	
	if namespace != "" {
		kubectlArgs = append(kubectlArgs, "-n", namespace)
	}

	// Execute kubectl
	kubectlCmd := exec.Command("kubectl", kubectlArgs...)
	kubectlCmd.Stdout = os.Stdout
	kubectlCmd.Stderr = os.Stderr
	kubectlCmd.Stdin = os.Stdin

	if err := kubectlCmd.Run(); err != nil {
		return fmt.Errorf("failed to execute kubectl: %w", err)
	}

	return nil
}
