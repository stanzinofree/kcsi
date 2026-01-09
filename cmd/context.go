package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/context"
)

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Manage kcsi contexts for different clusters",
	Long: `Manage kcsi contexts to easily switch between different Kubernetes clusters.
Contexts are stored separately from your system kubeconfig.`,
}

var contextAddCmd = &cobra.Command{
	Use:   "add <name> <kubeconfig-path>",
	Short: "Add a new context by referencing an existing kubeconfig",
	Long: `Add a new context by providing a name and path to an existing kubeconfig file.
The kubeconfig file will be referenced but not copied.`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		kubeconfigPath := args[1]

		// Resolve absolute path
		absPath, err := filepath.Abs(kubeconfigPath)
		if err != nil {
			return fmt.Errorf("failed to resolve path: %w", err)
		}

		// Check if file exists
		if _, err := os.Stat(absPath); os.IsNotExist(err) {
			return fmt.Errorf("kubeconfig file not found: %s", absPath)
		}

		description, _ := cmd.Flags().GetString("description")

		if err := context.AddContext(name, absPath, description); err != nil {
			return err
		}

		fmt.Printf("✓ Context '%s' added successfully\n", name)
		fmt.Printf("  Kubeconfig: %s\n", absPath)
		if description != "" {
			fmt.Printf("  Description: %s\n", description)
		}
		fmt.Println("\nUse 'kcsi context use " + name + "' to activate this context")

		return nil
	},
}

var contextImportCmd = &cobra.Command{
	Use:   "import <name> <kubeconfig-path>",
	Short: "Import a kubeconfig file into kcsi's managed directory",
	Long: `Import a kubeconfig file by copying it into kcsi's managed directory.
The file will be stored at ~/.kcsi/contexts/<name>/kube.config`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		kubeconfigPath := args[1]

		// Resolve absolute path
		absPath, err := filepath.Abs(kubeconfigPath)
		if err != nil {
			return fmt.Errorf("failed to resolve path: %w", err)
		}

		description, _ := cmd.Flags().GetString("description")

		if err := context.ImportContext(name, absPath, description); err != nil {
			return err
		}

		importedPath, _ := context.GetContextKubeconfigPath(name)
		fmt.Printf("✓ Context '%s' imported successfully\n", name)
		fmt.Printf("  Source: %s\n", absPath)
		fmt.Printf("  Imported to: %s\n", importedPath)
		if description != "" {
			fmt.Printf("  Description: %s\n", description)
		}
		fmt.Println("\nUse 'kcsi context use " + name + "' to activate this context")

		return nil
	},
}

var contextListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all available contexts",
	RunE: func(cmd *cobra.Command, args []string) error {
		contexts, err := context.ListContexts()
		if err != nil {
			return err
		}

		if len(contexts) == 0 {
			fmt.Println("No contexts configured yet.")
			fmt.Println("\nUse 'kcsi context import <name> <kubeconfig-path>' to add your first context")
			return nil
		}

		currentName, _ := context.GetCurrentContextName()

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "CURRENT\tNAME\tKUBECONFIG\tDEFAULT NS\tDESCRIPTION")

		for _, ctx := range contexts {
			current := ""
			if ctx.Name == currentName {
				current = "*"
			}
			description := ctx.Description
			if description == "" {
				description = "-"
			}
			defaultNS := ctx.DefaultNamespace
			if defaultNS == "" {
				defaultNS = "-"
			}
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", current, ctx.Name, ctx.KubeconfigPath, defaultNS, description)
		}

		w.Flush()
		return nil
	},
}

var contextUseCmd = &cobra.Command{
	Use:   "use <name>",
	Short: "Switch to a different context",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if err := context.SetCurrentContext(name); err != nil {
			return err
		}

		ctx, _ := context.GetContext(name)
		fmt.Printf("✓ Switched to context '%s'\n", name)
		if ctx != nil && ctx.Description != "" {
			fmt.Printf("  %s\n", ctx.Description)
		}

		return nil
	},
}

var contextCurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "Display the current context",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := context.GetCurrentContext()
		if err != nil {
			if err.Error() == "no current context set" {
				fmt.Println("No context currently active")
				fmt.Println("\nUse 'kcsi context use <name>' to activate a context")
				return nil
			}
			return err
		}

		fmt.Printf("Current context: %s\n", ctx.Name)
		fmt.Printf("Kubeconfig: %s\n", ctx.KubeconfigPath)
		if ctx.Description != "" {
			fmt.Printf("Description: %s\n", ctx.Description)
		}

		return nil
	},
}

var contextRemoveCmd = &cobra.Command{
	Use:     "remove <name>",
	Aliases: []string{"rm", "delete"},
	Short:   "Remove a context",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if err := context.RemoveContext(name); err != nil {
			return err
		}

		fmt.Printf("✓ Context '%s' removed successfully\n", name)
		return nil
	},
}

var contextSetNamespaceCmd = &cobra.Command{
	Use:   "set-namespace <namespace>",
	Short: "Set default namespace for current context",
	Long: `Set a default namespace for the current active context.
When a default namespace is set, all kcsi commands will use it automatically 
if no -n/--namespace flag is explicitly provided.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		namespace := args[0]

		// Get current context name
		currentName, err := context.GetCurrentContextName()
		if err != nil {
			return err
		}
		if currentName == "" {
			return fmt.Errorf("no active context. Use 'kcsi context use <name>' first")
		}

		// Set default namespace
		if err := context.SetDefaultNamespace(currentName, namespace); err != nil {
			return err
		}

		fmt.Printf("✓ Default namespace set to '%s' for context '%s'\n", namespace, currentName)
		fmt.Println("\nAll kcsi commands will now use this namespace by default.")
		fmt.Println("You can still override it with -n/--namespace flag.")

		return nil
	},
}

var contextClearNamespaceCmd = &cobra.Command{
	Use:   "clear-namespace",
	Short: "Clear default namespace for current context",
	Long: `Remove the default namespace setting from the current active context.
After clearing, kcsi commands will use kubectl's default behavior (usually 'default' namespace).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get current context name
		currentName, err := context.GetCurrentContextName()
		if err != nil {
			return err
		}
		if currentName == "" {
			return fmt.Errorf("no active context. Use 'kcsi context use <name>' first")
		}

		// Clear default namespace
		if err := context.ClearDefaultNamespace(currentName); err != nil {
			return err
		}

		fmt.Printf("✓ Default namespace cleared for context '%s'\n", currentName)
		fmt.Println("\nkcsi commands will now use kubectl's default namespace behavior.")

		return nil
	},
}

var contextGetNamespaceCmd = &cobra.Command{
	Use:   "get-namespace",
	Short: "Show default namespace for current context",
	Long:  `Display the default namespace configured for the current active context.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get current context
		ctx, err := context.GetCurrentContext()
		if err != nil {
			if err.Error() == "no current context set" {
				fmt.Println("No context currently active")
				fmt.Println("\nUse 'kcsi context use <name>' to activate a context")
				return nil
			}
			return err
		}

		if ctx.DefaultNamespace == "" {
			fmt.Printf("No default namespace set for context '%s'\n", ctx.Name)
			fmt.Println("\nUse 'kcsi context set-namespace <namespace>' to set one")
		} else {
			fmt.Printf("Default namespace for context '%s': %s\n", ctx.Name, ctx.DefaultNamespace)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(contextCmd)

	// Add subcommands
	contextCmd.AddCommand(contextAddCmd)
	contextCmd.AddCommand(contextImportCmd)
	contextCmd.AddCommand(contextListCmd)
	contextCmd.AddCommand(contextUseCmd)
	contextCmd.AddCommand(contextCurrentCmd)
	contextCmd.AddCommand(contextRemoveCmd)
	contextCmd.AddCommand(contextSetNamespaceCmd)
	contextCmd.AddCommand(contextClearNamespaceCmd)
	contextCmd.AddCommand(contextGetNamespaceCmd)

	// Add flags
	contextAddCmd.Flags().StringP("description", "d", "", "Description of the context")
	contextImportCmd.Flags().StringP("description", "d", "", "Description of the context")
}
