package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
)

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Display resource usage statistics",
	Long:  "Display resource usage statistics for pods or nodes",
}

var topPodsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Display resource usage of pods",
	Long:  "Display CPU and memory usage of pods in a namespace",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: runTopPods,
}

var topNodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "Display resource usage of nodes",
	Long:  "Display CPU and memory usage of nodes in the cluster",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) == 0 {
			nodes, err := kubernetes.GetNodes()
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}
			return nodes, cobra.ShellCompDirectiveNoFileComp
		}
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: runTopNodes,
}

func init() {
	rootCmd.AddCommand(topCmd)
	topCmd.AddCommand(topPodsCmd)
	topCmd.AddCommand(topNodesCmd)

	// Add namespace flag to pods subcommand
	topPodsCmd.Flags().StringP("namespace", "n", "", "Namespace for the pods")
	topPodsCmd.RegisterFlagCompletionFunc("namespace", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		namespaces, err := kubernetes.GetNamespaces()
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return namespaces, cobra.ShellCompDirectiveNoFileComp
	})
}

func runTopPods(cmd *cobra.Command, args []string) error {
	namespace, _ := cmd.Flags().GetString("namespace")

	kubectlArgs := []string{"top", "pods"}

	if namespace != "" {
		kubectlArgs = append(kubectlArgs, "-n", namespace)
	} else {
		kubectlArgs = append(kubectlArgs, "--all-namespaces")
	}

	return kubernetes.ExecuteKubectlInteractive(kubectlArgs...)
}

func runTopNodes(cmd *cobra.Command, args []string) error {
	kubectlArgs := []string{"top", "nodes"}

	if len(args) > 0 {
		kubectlArgs = append(kubectlArgs, args[0])
	}

	return kubernetes.ExecuteKubectlInteractive(kubectlArgs...)
}
