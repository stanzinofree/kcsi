package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/completion"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe Kubernetes resources",
	Long:  `Describe Kubernetes resources with smart autocompletion`,
}

// Namespace and container flags for different describe commands
var (
	describePodNamespace        string
	describePodContainer        string
	describeServiceNamespace    string
	describeDeploymentNamespace string
	describeConfigMapNamespace  string
	describeSecretNamespace     string
)

// Generic kubectl describe command runner
func runKubectlDescribe(resourceType, namespace, container string, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("resource name is required")
	}

	// Build kubectl command with namespace injection
	kubectlArgs := kubernetes.BuildNamespaceArgs([]string{"describe", resourceType}, namespace)
	kubectlArgs = append(kubectlArgs, args[0])

	if container != "" {
		kubectlArgs = append(kubectlArgs, "-c", container)
	}

	// Execute kubectl using the helper
	return kubernetes.ExecuteKubectlInteractive(kubectlArgs...)
}

// Pod
var describePodCmd = &cobra.Command{
	Use:   "pod [pod-name]",
	Short: "Describe a pod",
	Long:  `Describe a specific pod with namespace and pod name autocompletion`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlDescribe("pod", describePodNamespace, describePodContainer, args)
	},
	ValidArgsFunction: completion.PodCompletion,
}

// Service
var describeServiceCmd = &cobra.Command{
	Use:     "service [service-name]",
	Aliases: []string{"svc", "services"},
	Short:   "Describe a service",
	Long:    `Describe a specific service with namespace and service name autocompletion`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlDescribe("service", describeServiceNamespace, "", args)
	},
	ValidArgsFunction: completion.ServiceCompletion,
}

// Deployment
var describeDeploymentCmd = &cobra.Command{
	Use:     "deployment [deployment-name]",
	Aliases: []string{"deploy", "deployments"},
	Short:   "Describe a deployment",
	Long:    `Describe a specific deployment with namespace and deployment name autocompletion`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlDescribe("deployment", describeDeploymentNamespace, "", args)
	},
	ValidArgsFunction: completion.DeploymentCompletion,
}

// Node
var describeNodeCmd = &cobra.Command{
	Use:     "node [node-name]",
	Aliases: []string{"nodes"},
	Short:   "Describe a node",
	Long:    `Describe a specific node with node name autocompletion`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlDescribe("node", "", "", args)
	},
	ValidArgsFunction: completion.NodeCompletion,
}

// ConfigMap
var describeConfigMapCmd = &cobra.Command{
	Use:     "configmap [configmap-name]",
	Aliases: []string{"cm", "configmaps"},
	Short:   "Describe a configmap",
	Long:    `Describe a specific configmap with namespace and configmap name autocompletion`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlDescribe("configmap", describeConfigMapNamespace, "", args)
	},
	ValidArgsFunction: completion.ConfigMapCompletion,
}

// Secret
var describeSecretCmd = &cobra.Command{
	Use:     "secret [secret-name]",
	Aliases: []string{"secrets"},
	Short:   "Describe a secret",
	Long:    `Describe a specific secret with namespace and secret name autocompletion`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlDescribe("secret", describeSecretNamespace, "", args)
	},
	ValidArgsFunction: completion.SecretCompletion,
}

func init() {
	rootCmd.AddCommand(describeCmd)

	// Add all subcommands
	describeCmd.AddCommand(describePodCmd)
	describeCmd.AddCommand(describeServiceCmd)
	describeCmd.AddCommand(describeDeploymentCmd)
	describeCmd.AddCommand(describeNodeCmd)
	describeCmd.AddCommand(describeConfigMapCmd)
	describeCmd.AddCommand(describeSecretCmd)

	// Add namespace flags with autocompletion for namespaced resources
	describePodCmd.Flags().StringVarP(&describePodNamespace, "namespace", "n", "", FlagDescNamespace)
	describePodCmd.Flags().StringVarP(&describePodContainer, "container", "c", "", "Container name (for multi-container pods)")
	describePodCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)
	describePodCmd.RegisterFlagCompletionFunc("container", completion.ContainerCompletion)

	describeServiceCmd.Flags().StringVarP(&describeServiceNamespace, "namespace", "n", "", FlagDescNamespace)
	describeServiceCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)

	describeDeploymentCmd.Flags().StringVarP(&describeDeploymentNamespace, "namespace", "n", "", FlagDescNamespace)
	describeDeploymentCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)

	describeConfigMapCmd.Flags().StringVarP(&describeConfigMapNamespace, "namespace", "n", "", FlagDescNamespace)
	describeConfigMapCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)

	describeSecretCmd.Flags().StringVarP(&describeSecretNamespace, "namespace", "n", "", FlagDescNamespace)
	describeSecretCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)
}
