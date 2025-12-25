package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/stanzinofree/kcsi/pkg/completion"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Kubernetes resources",
	Long:  `Get Kubernetes resources with smart autocompletion`,
}

// Namespace and output flags for different get commands
var (
	getPodsNamespace        string
	getServicesNamespace    string
	getDeploymentsNamespace string
	getConfigMapsNamespace  string
	getSecretsNamespace     string
	getPodsOutput           string
	getServicesOutput       string
	getDeploymentsOutput    string
	getNodesOutput          string
	getConfigMapsOutput     string
	getSecretsOutput        string
)

// Generic kubectl get command runner
func runKubectlGet(resourceType, namespace, output string, args []string) error {
	kubectlArgs := []string{"get", resourceType}
	
	if namespace != "" {
		kubectlArgs = append(kubectlArgs, "-n", namespace)
	}

	if output != "" {
		kubectlArgs = append(kubectlArgs, "-o", output)
	}

	// Append any additional args passed
	kubectlArgs = append(kubectlArgs, args...)

	kubectlCmd := exec.Command("kubectl", kubectlArgs...)
	kubectlCmd.Stdout = os.Stdout
	kubectlCmd.Stderr = os.Stderr
	kubectlCmd.Stdin = os.Stdin

	if err := kubectlCmd.Run(); err != nil {
		return fmt.Errorf("failed to execute kubectl: %w", err)
	}

	return nil
}

// Pods command
var getPodsCmd = &cobra.Command{
	Use:   "pods [pod-name]",
	Short: "Get pods in a namespace",
	Long:  `Get pods in a specific namespace with autocompletion support`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlGet("pods", getPodsNamespace, getPodsOutput, args)
	},
	ValidArgsFunction: completion.PodCompletion,
}

// Namespaces command
var getNamespacesCmd = &cobra.Command{
	Use:     "namespaces",
	Aliases: []string{"ns", "namespace"},
	Short:   "Get namespaces",
	Long:    `Get all namespaces in the cluster`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlGet("namespaces", "", "", args)
	},
}

// Services command
var getServicesCmd = &cobra.Command{
	Use:     "services [service-name]",
	Aliases: []string{"svc", "service"},
	Short:   "Get services in a namespace",
	Long:    `Get services in a specific namespace with autocompletion support`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlGet("services", getServicesNamespace, getServicesOutput, args)
	},
	ValidArgsFunction: completion.ServiceCompletion,
}

// Deployments command
var getDeploymentsCmd = &cobra.Command{
	Use:     "deployments [deployment-name]",
	Aliases: []string{"deploy", "deployment"},
	Short:   "Get deployments in a namespace",
	Long:    `Get deployments in a specific namespace with autocompletion support`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlGet("deployments", getDeploymentsNamespace, getDeploymentsOutput, args)
	},
	ValidArgsFunction: completion.DeploymentCompletion,
}

// Nodes command
var getNodesCmd = &cobra.Command{
	Use:     "nodes [node-name]",
	Aliases: []string{"no", "node"},
	Short:   "Get nodes",
	Long:    `Get nodes in the cluster`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlGet("nodes", "", getNodesOutput, args)
	},
	ValidArgsFunction: completion.NodeCompletion,
}

// ConfigMaps command
var getConfigMapsCmd = &cobra.Command{
	Use:     "configmaps [configmap-name]",
	Aliases: []string{"cm", "configmap"},
	Short:   "Get configmaps in a namespace",
	Long:    `Get configmaps in a specific namespace with autocompletion support`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlGet("configmaps", getConfigMapsNamespace, getConfigMapsOutput, args)
	},
	ValidArgsFunction: completion.ConfigMapCompletion,
}

// Secrets command
var getSecretsCmd = &cobra.Command{
	Use:     "secrets [secret-name]",
	Aliases: []string{"secret"},
	Short:   "Get secrets in a namespace",
	Long:    `Get secrets in a specific namespace with autocompletion support`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlGet("secrets", getSecretsNamespace, getSecretsOutput, args)
	},
	ValidArgsFunction: completion.SecretCompletion,
}

func init() {
	rootCmd.AddCommand(getCmd)
	
	// Add all subcommands
	getCmd.AddCommand(getPodsCmd)
	getCmd.AddCommand(getNamespacesCmd)
	getCmd.AddCommand(getServicesCmd)
	getCmd.AddCommand(getDeploymentsCmd)
	getCmd.AddCommand(getNodesCmd)
	getCmd.AddCommand(getConfigMapsCmd)
	getCmd.AddCommand(getSecretsCmd)

	// Add namespace flags with autocompletion for namespaced resources
	getPodsCmd.Flags().StringVarP(&getPodsNamespace, "namespace", "n", "", "Kubernetes namespace")
	getPodsCmd.Flags().StringVarP(&getPodsOutput, "output", "o", "", "Output format (e.g., wide, yaml, json)")
	getPodsCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)

	getServicesCmd.Flags().StringVarP(&getServicesNamespace, "namespace", "n", "", "Kubernetes namespace")
	getServicesCmd.Flags().StringVarP(&getServicesOutput, "output", "o", "", "Output format (e.g., wide, yaml, json)")
	getServicesCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)

	getDeploymentsCmd.Flags().StringVarP(&getDeploymentsNamespace, "namespace", "n", "", "Kubernetes namespace")
	getDeploymentsCmd.Flags().StringVarP(&getDeploymentsOutput, "output", "o", "", "Output format (e.g., wide, yaml, json)")
	getDeploymentsCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)

	getNodesCmd.Flags().StringVarP(&getNodesOutput, "output", "o", "", "Output format (e.g., wide, yaml, json)")

	getConfigMapsCmd.Flags().StringVarP(&getConfigMapsNamespace, "namespace", "n", "", "Kubernetes namespace")
	getConfigMapsCmd.Flags().StringVarP(&getConfigMapsOutput, "output", "o", "", "Output format (e.g., wide, yaml, json)")
	getConfigMapsCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)

	getSecretsCmd.Flags().StringVarP(&getSecretsNamespace, "namespace", "n", "", "Kubernetes namespace")
	getSecretsCmd.Flags().StringVarP(&getSecretsOutput, "output", "o", "", "Output format (e.g., wide, yaml, json)")
	getSecretsCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)
}
