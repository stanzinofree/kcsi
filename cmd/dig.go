package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
)

var digCmd = &cobra.Command{
	Use:   "dig [namespace] [pod] [domain]",
	Short: "DNS debugging inside a pod",
	Long:  "Run DNS queries (dig command) inside a pod to debug DNS resolution issues",
	Args:  cobra.MinimumNArgs(2),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) == 0 {
			// First arg: namespace
			namespaces, err := kubernetes.GetNamespaces()
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}
			return namespaces, cobra.ShellCompDirectiveNoFileComp
		}
		if len(args) == 1 {
			// Second arg: pod in that namespace
			pods, err := kubernetes.GetPods(args[0])
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}
			return pods, cobra.ShellCompDirectiveNoFileComp
		}
		// Third arg: domain name - no completion
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: runDig,
}

func init() {
	rootCmd.AddCommand(digCmd)
	digCmd.Flags().StringP("container", "c", "", "Container name for multi-container pods")
	digCmd.RegisterFlagCompletionFunc("container", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) >= 2 {
			namespace := args[0]
			podName := args[1]
			containers, err := kubernetes.GetContainers(namespace, podName)
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}
			return containers, cobra.ShellCompDirectiveNoFileComp
		}
		return nil, cobra.ShellCompDirectiveNoFileComp
	})
}

func runDig(cmd *cobra.Command, args []string) error {
	namespace := args[0]
	podName := args[1]
	
	// Build the dig command
	digCommand := "dig"
	
	// If domain is specified, add it
	if len(args) >= 3 {
		digCommand = fmt.Sprintf("dig %s", args[2])
	}
	
	// Add any additional dig arguments
	if len(args) > 3 {
		for i := 3; i < len(args); i++ {
			digCommand = fmt.Sprintf("%s %s", digCommand, args[i])
		}
	}

	kubectlArgs := []string{"exec", "-n", namespace, "-it", podName, "--"}
	
	// Check if container flag is set
	container, _ := cmd.Flags().GetString("container")
	if container != "" {
		kubectlArgs = []string{"exec", "-n", namespace, "-it", podName, "-c", container, "--"}
	}
	
	// Add the dig command
	kubectlArgs = append(kubectlArgs, "sh", "-c", digCommand)

	return kubernetes.ExecuteKubectlInteractive(kubectlArgs...)
}
