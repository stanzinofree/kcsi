package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
)

var rolloutCmd = &cobra.Command{
	Use:   "rollout",
	Short: "Manage rollout of resources",
	Long:  "Manage the rollout of Kubernetes resources (restart, status, history, undo)",
}

var rolloutRestartCmd = &cobra.Command{
	Use:               "restart [resource-type] [name]",
	Short:             "Restart a resource",
	Long:              "Restart a deployment, daemonset, or statefulset by triggering a rollout",
	Args:              cobra.ExactArgs(2),
	ValidArgsFunction: resourceNameCompletion,
	RunE:              runRolloutRestart,
}

var rolloutStatusCmd = &cobra.Command{
	Use:               "status [resource-type] [name]",
	Short:             "Show rollout status",
	Long:              "Show the status of the rollout for a deployment, daemonset, or statefulset",
	Args:              cobra.ExactArgs(2),
	ValidArgsFunction: resourceNameCompletion,
	RunE:              runRolloutStatus,
}

var rolloutHistoryCmd = &cobra.Command{
	Use:               "history [resource-type] [name]",
	Short:             "View rollout history",
	Long:              "View previous rollout revisions and configurations",
	Args:              cobra.ExactArgs(2),
	ValidArgsFunction: resourceNameCompletion,
	RunE:              runRolloutHistory,
}

var rolloutUndoCmd = &cobra.Command{
	Use:               "undo [resource-type] [name]",
	Short:             "Undo a rollout",
	Long:              "Rollback to a previous revision",
	Args:              cobra.ExactArgs(2),
	ValidArgsFunction: resourceNameCompletion,
	RunE:              runRolloutUndo,
}

func init() {
	rootCmd.AddCommand(rolloutCmd)
	rolloutCmd.AddCommand(rolloutRestartCmd)
	rolloutCmd.AddCommand(rolloutStatusCmd)
	rolloutCmd.AddCommand(rolloutHistoryCmd)
	rolloutCmd.AddCommand(rolloutUndoCmd)

	// Add namespace flag to all rollout subcommands
	for _, cmd := range []*cobra.Command{rolloutRestartCmd, rolloutStatusCmd, rolloutHistoryCmd, rolloutUndoCmd} {
		cmd.Flags().StringP("namespace", "n", "", FlagDescNamespace)
		cmd.RegisterFlagCompletionFunc("namespace", func(cmd *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
			namespaces, err := kubernetes.GetNamespaces()
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}
			return namespaces, cobra.ShellCompDirectiveNoFileComp
		})
	}

	// Add revision flag to undo command
	rolloutUndoCmd.Flags().Int("to-revision", 0, "Revision to rollback to (0 means previous revision)")
}

func resourceNameCompletion(cmd *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
	if len(args) == 0 {
		// First argument: resource type
		return []string{"deployment", "daemonset", "statefulset"}, cobra.ShellCompDirectiveNoFileComp
	}

	if len(args) == 1 {
		// Second argument: resource name based on type
		namespace, _ := cmd.Flags().GetString("namespace")
		if namespace == "" {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		resourceType := args[0]
		var resources []string
		var err error

		switch resourceType {
		case "deployment", "deployments", "deploy":
			resources, err = kubernetes.GetDeployments(namespace)
		case "daemonset", "daemonsets", "ds":
			resources, err = kubernetes.GetDaemonSets(namespace)
		case "statefulset", "statefulsets", "sts":
			resources, err = kubernetes.GetStatefulSets(namespace)
		default:
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return resources, cobra.ShellCompDirectiveNoFileComp
	}

	return nil, cobra.ShellCompDirectiveNoFileComp
}

func runRolloutRestart(cmd *cobra.Command, args []string) error {
	namespace, _ := cmd.Flags().GetString("namespace")
	if namespace == "" {
		return fmt.Errorf("namespace is required (use -n flag)")
	}

	resourceType := args[0]
	resourceName := args[1]

	output, err := kubernetes.ExecuteKubectl("rollout", "restart", resourceType, resourceName, "-n", namespace)
	if err != nil {
		return fmt.Errorf("failed to restart rollout: %v", err)
	}

	fmt.Println(output)
	return nil
}

func runRolloutStatus(cmd *cobra.Command, args []string) error {
	namespace, _ := cmd.Flags().GetString("namespace")
	if namespace == "" {
		return fmt.Errorf("namespace is required (use -n flag)")
	}

	resourceType := args[0]
	resourceName := args[1]

	output, err := kubernetes.ExecuteKubectl("rollout", "status", resourceType, resourceName, "-n", namespace)
	if err != nil {
		return fmt.Errorf("failed to get rollout status: %v", err)
	}

	fmt.Println(output)
	return nil
}

func runRolloutHistory(cmd *cobra.Command, args []string) error {
	namespace, _ := cmd.Flags().GetString("namespace")
	if namespace == "" {
		return fmt.Errorf("namespace is required (use -n flag)")
	}

	resourceType := args[0]
	resourceName := args[1]

	output, err := kubernetes.ExecuteKubectl("rollout", "history", resourceType, resourceName, "-n", namespace)
	if err != nil {
		return fmt.Errorf("failed to get rollout history: %v", err)
	}

	fmt.Println(output)
	return nil
}

func runRolloutUndo(cmd *cobra.Command, args []string) error {
	namespace, _ := cmd.Flags().GetString("namespace")
	if namespace == "" {
		return fmt.Errorf("namespace is required (use -n flag)")
	}

	resourceType := args[0]
	resourceName := args[1]
	toRevision, _ := cmd.Flags().GetInt("to-revision")

	cmdArgs := []string{"rollout", "undo", resourceType, resourceName, "-n", namespace}
	if toRevision > 0 {
		cmdArgs = append(cmdArgs, fmt.Sprintf("--to-revision=%d", toRevision))
	}

	output, err := kubernetes.ExecuteKubectl(cmdArgs...)
	if err != nil {
		return fmt.Errorf("failed to undo rollout: %v", err)
	}

	fmt.Println(output)
	return nil
}
