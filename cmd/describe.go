package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alessandro/kcsi/pkg/completion"
	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe Kubernetes resources",
	Long:  `Describe Kubernetes resources with smart autocompletion`,
}

var describePodCmd = &cobra.Command{
	Use:   "pod [pod-name]",
	Short: "Describe a pod",
	Long:  `Describe a specific pod with namespace and pod name autocompletion`,
	Args:  cobra.ExactArgs(1),
	RunE:  runDescribePod,
	ValidArgsFunction: completion.PodCompletion,
}

var describeNamespace string

func init() {
	rootCmd.AddCommand(describeCmd)
	describeCmd.AddCommand(describePodCmd)

	// Add namespace flag with autocompletion
	describePodCmd.Flags().StringVarP(&describeNamespace, "namespace", "n", "", "Kubernetes namespace")
	describePodCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)
}

func runDescribePod(cmd *cobra.Command, args []string) error {
	podName := args[0]

	// Build kubectl command
	kubectlArgs := []string{"describe", "pod", podName}
	
	if describeNamespace != "" {
		kubectlArgs = append(kubectlArgs, "-n", describeNamespace)
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
