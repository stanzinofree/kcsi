package completion

import (
	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
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

// ServiceCompletion provides autocompletion for service names
func ServiceCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	namespace, _ := cmd.Flags().GetString("namespace")

	services, err := kubernetes.GetServices(namespace)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	return services, cobra.ShellCompDirectiveNoFileComp
}

// DeploymentCompletion provides autocompletion for deployment names
func DeploymentCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	namespace, _ := cmd.Flags().GetString("namespace")

	deployments, err := kubernetes.GetDeployments(namespace)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	return deployments, cobra.ShellCompDirectiveNoFileComp
}

// NodeCompletion provides autocompletion for node names
func NodeCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	nodes, err := kubernetes.GetNodes()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	return nodes, cobra.ShellCompDirectiveNoFileComp
}

// ConfigMapCompletion provides autocompletion for configmap names
func ConfigMapCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	namespace, _ := cmd.Flags().GetString("namespace")

	configmaps, err := kubernetes.GetConfigMaps(namespace)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	return configmaps, cobra.ShellCompDirectiveNoFileComp
}

// SecretCompletion provides autocompletion for secret names
func SecretCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	namespace, _ := cmd.Flags().GetString("namespace")

	secrets, err := kubernetes.GetSecrets(namespace)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}
	return secrets, cobra.ShellCompDirectiveNoFileComp
}
