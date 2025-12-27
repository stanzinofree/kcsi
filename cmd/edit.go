package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
)

var editCmd = &cobra.Command{
	Use:   "edit [resource-type] [name]",
	Short: "Edit a resource with automatic backup",
	Long:  "Edit a resource on the server using the default editor. A backup of the current state is automatically saved before editing.",
	Args:  cobra.MinimumNArgs(2),
	RunE:  runEdit,
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringP("namespace", "n", "", FlagDescNamespace)
	editCmd.Flags().StringP("output", "o", "yaml", "Output format for the backup (yaml or json)")
	editCmd.Flags().String("backup-dir", "", "Directory to save backups (defaults to ~/.kcsi/backups)")
	editCmd.Flags().Bool("no-backup", false, "Skip automatic backup before editing")
	editCmd.Flags().StringP("editor", "e", "", "Editor to use (defaults to KUBE_EDITOR or EDITOR environment variable)")

	editCmd.RegisterFlagCompletionFunc("namespace", func(cmd *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
		namespaces, err := kubernetes.GetNamespaces()
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return namespaces, cobra.ShellCompDirectiveNoFileComp
	})
}

func runEdit(cmd *cobra.Command, args []string) error {
	resourceType := args[0]
	resourceName := args[1]
	namespace, _ := cmd.Flags().GetString("namespace")
	outputFormat, _ := cmd.Flags().GetString("output")
	backupDir, _ := cmd.Flags().GetString("backup-dir")
	noBackup, _ := cmd.Flags().GetBool("no-backup")
	editor, _ := cmd.Flags().GetString("editor")

	if namespace == "" {
		return fmt.Errorf("namespace is required (use -n flag)")
	}

	// Create backup unless --no-backup is specified
	var backupPath string
	if !noBackup {
		var err error
		backupPath, err = createResourceBackup(resourceType, resourceName, namespace, outputFormat, backupDir)
		if err != nil {
			return fmt.Errorf("failed to create backup: %v", err)
		}
		fmt.Printf("✅ Backup saved to: %s\n", backupPath)
		fmt.Println()
	}

	// Build kubectl edit arguments
	args = []string{"edit", resourceType, resourceName, "-n", namespace}

	// Add output format
	if outputFormat != "" {
		args = append(args, "-o", outputFormat)
	}

	// Set editor if specified
	if editor != "" {
		os.Setenv("KUBE_EDITOR", editor)
	}

	// Execute kubectl edit interactively
	fmt.Printf("📝 Opening editor for %s/%s in namespace %s...\n", resourceType, resourceName, namespace)
	fmt.Println()

	err := kubernetes.ExecuteKubectlInteractive(args...)
	if err != nil {
		if !noBackup && backupPath != "" {
			fmt.Printf("\n⚠️  Edit failed. You can restore from backup: %s\n", backupPath)
		}
		return fmt.Errorf("failed to edit resource: %v", err)
	}

	fmt.Println()
	fmt.Println("✅ Resource updated successfully")
	if !noBackup && backupPath != "" {
		fmt.Printf("💾 Previous state backed up at: %s\n", backupPath)
	}

	return nil
}

func createResourceBackup(resourceType, resourceName, namespace, outputFormat, backupDir string) (string, error) {
	// Determine backup directory
	if backupDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get home directory: %v", err)
		}
		backupDir = filepath.Join(homeDir, ".kcsi", "backups")
	}

	// Create backup directory if it doesn't exist
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create backup directory: %v", err)
	}

	// Generate backup filename with timestamp
	timestamp := time.Now().Format("20060102-150405")
	filename := fmt.Sprintf("%s-%s-%s-%s.%s", resourceType, resourceName, namespace, timestamp, outputFormat)
	backupPath := filepath.Join(backupDir, filename)

	// Get current resource state
	output, err := kubernetes.ExecuteKubectl("get", resourceType, resourceName, "-n", namespace, "-o", outputFormat)
	if err != nil {
		return "", fmt.Errorf("failed to get resource state: %v", err)
	}

	// Write backup file
	if err := os.WriteFile(backupPath, []byte(output), 0644); err != nil {
		return "", fmt.Errorf("failed to write backup file: %v", err)
	}

	return backupPath, nil
}
