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
