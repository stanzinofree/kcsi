package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/alessandro/kcsi/pkg/completion"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Kubernetes resources",
	Long:  `Delete Kubernetes resources with confirmation prompts for safety`,
}

// Namespace flags for different delete commands
var (
	deletePodNamespace        string
	deleteServiceNamespace    string
	deleteDeploymentNamespace string
	deleteConfigMapNamespace  string
	deleteSecretNamespace     string
	deleteForce               bool
)

// askForConfirmation prompts the user for yes/no confirmation
func askForConfirmation(resourceType, resourceName, namespace string) bool {
	reader := bufio.NewReader(os.Stdin)
	
	nsInfo := ""
	if namespace != "" {
		nsInfo = fmt.Sprintf(" in namespace '%s'", namespace)
	}
	
	fmt.Printf("Are you sure you want to delete %s '%s'%s? [y/N]: ", resourceType, resourceName, nsInfo)
	
	response, err := reader.ReadString('\n')
	if err != nil {
		return false
	}
	
	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
}

// Generic kubectl delete command runner with confirmation
func runKubectlDelete(resourceType, namespace string, args []string, force bool) error {
	if len(args) == 0 {
		return fmt.Errorf("resource name is required")
	}

	resourceName := args[0]

	// Ask for confirmation unless --force is used
	if !force {
		if !askForConfirmation(resourceType, resourceName, namespace) {
			fmt.Println("Delete cancelled.")
			return nil
		}
	}

	kubectlArgs := []string{"delete", resourceType, resourceName}
	
	if namespace != "" {
		kubectlArgs = append(kubectlArgs, "-n", namespace)
	}

	fmt.Printf("Deleting %s '%s'...\n", resourceType, resourceName)

	kubectlCmd := exec.Command("kubectl", kubectlArgs...)
	kubectlCmd.Stdout = os.Stdout
	kubectlCmd.Stderr = os.Stderr
	kubectlCmd.Stdin = os.Stdin

	if err := kubectlCmd.Run(); err != nil {
		return fmt.Errorf("failed to execute kubectl: %w", err)
	}

	return nil
}

// Pod
var deletePodCmd = &cobra.Command{
	Use:   "pod [pod-name]",
	Short: "Delete a pod",
	Long:  `Delete a specific pod with confirmation prompt`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlDelete("pod", deletePodNamespace, args, deleteForce)
	},
	ValidArgsFunction: completion.PodCompletion,
}

// Service
var deleteServiceCmd = &cobra.Command{
	Use:     "service [service-name]",
	Aliases: []string{"svc", "services"},
	Short:   "Delete a service",
	Long:    `Delete a specific service with confirmation prompt`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlDelete("service", deleteServiceNamespace, args, deleteForce)
	},
	ValidArgsFunction: completion.ServiceCompletion,
}

// Deployment
var deleteDeploymentCmd = &cobra.Command{
	Use:     "deployment [deployment-name]",
	Aliases: []string{"deploy", "deployments"},
	Short:   "Delete a deployment",
	Long:    `Delete a specific deployment with confirmation prompt`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlDelete("deployment", deleteDeploymentNamespace, args, deleteForce)
	},
	ValidArgsFunction: completion.DeploymentCompletion,
}

// ConfigMap
var deleteConfigMapCmd = &cobra.Command{
	Use:     "configmap [configmap-name]",
	Aliases: []string{"cm", "configmaps"},
	Short:   "Delete a configmap",
	Long:    `Delete a specific configmap with confirmation prompt`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlDelete("configmap", deleteConfigMapNamespace, args, deleteForce)
	},
	ValidArgsFunction: completion.ConfigMapCompletion,
}

// Secret
var deleteSecretCmd = &cobra.Command{
	Use:     "secret [secret-name]",
	Aliases: []string{"secrets"},
	Short:   "Delete a secret",
	Long:    `Delete a specific secret with confirmation prompt`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runKubectlDelete("secret", deleteSecretNamespace, args, deleteForce)
	},
	ValidArgsFunction: completion.SecretCompletion,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	
	// Add all subcommands
	deleteCmd.AddCommand(deletePodCmd)
	deleteCmd.AddCommand(deleteServiceCmd)
	deleteCmd.AddCommand(deleteDeploymentCmd)
	deleteCmd.AddCommand(deleteConfigMapCmd)
	deleteCmd.AddCommand(deleteSecretCmd)

	// Add namespace flags with autocompletion for namespaced resources
	deletePodCmd.Flags().StringVarP(&deletePodNamespace, "namespace", "n", "", "Kubernetes namespace")
	deletePodCmd.Flags().BoolVarP(&deleteForce, "force", "f", false, "Skip confirmation prompt")
	deletePodCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)

	deleteServiceCmd.Flags().StringVarP(&deleteServiceNamespace, "namespace", "n", "", "Kubernetes namespace")
	deleteServiceCmd.Flags().BoolVarP(&deleteForce, "force", "f", false, "Skip confirmation prompt")
	deleteServiceCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)

	deleteDeploymentCmd.Flags().StringVarP(&deleteDeploymentNamespace, "namespace", "n", "", "Kubernetes namespace")
	deleteDeploymentCmd.Flags().BoolVarP(&deleteForce, "force", "f", false, "Skip confirmation prompt")
	deleteDeploymentCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)

	deleteConfigMapCmd.Flags().StringVarP(&deleteConfigMapNamespace, "namespace", "n", "", "Kubernetes namespace")
	deleteConfigMapCmd.Flags().BoolVarP(&deleteForce, "force", "f", false, "Skip confirmation prompt")
	deleteConfigMapCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)

	deleteSecretCmd.Flags().StringVarP(&deleteSecretNamespace, "namespace", "n", "", "Kubernetes namespace")
	deleteSecretCmd.Flags().BoolVarP(&deleteForce, "force", "f", false, "Skip confirmation prompt")
	deleteSecretCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)
}
