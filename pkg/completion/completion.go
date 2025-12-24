package completion

import (
	"github.com/alessandro/kcsi/pkg/kubernetes"
	"github.com/spf13/cobra"
)

// NamespaceCompletion provides autocompletion for namespace flags
func NamespaceCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	namespaces, err := kubernetes.GetNamespaces()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	return namespaces, cobra.ShellCompDirectiveNoFileComp
}

// PodCompletion provides autocompletion for pod names
// It reads the namespace from the -n flag if provided
func PodCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	namespace, _ := cmd.Flags().GetString("namespace")
	
	pods, err := kubernetes.GetPods(namespace)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	return pods, cobra.ShellCompDirectiveNoFileComp
}

// ContainerCompletion provides autocompletion for container names within a pod
// It reads the namespace from the -n flag and the pod name from args[0] if provided
func ContainerCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// Need both namespace and pod name to get containers
	if len(args) == 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	namespace, _ := cmd.Flags().GetString("namespace")
	podName := args[0]
	
	containers, err := kubernetes.GetContainers(namespace, podName)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	return containers, cobra.ShellCompDirectiveNoFileComp
}
