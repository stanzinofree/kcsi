package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a configuration to a resource",
	Long:  "Apply a configuration to a resource by file name or stdin. The resource name must be specified.",
	RunE:  runApply,
}

func init() {
	rootCmd.AddCommand(applyCmd)

	applyCmd.Flags().StringP("filename", "f", "", "Filename, directory, or URL to files to use to create the resource")
	applyCmd.Flags().StringP("namespace", "n", "", FlagDescNamespace)
	applyCmd.Flags().Bool("dry-run", false, "Run in dry-run mode (client or server)")
	applyCmd.Flags().Bool("server-dry-run", false, "Run in server dry-run mode")
	applyCmd.Flags().Bool("validate", true, "Validate the configuration before applying")
	applyCmd.Flags().Bool("force", false, "Force apply (delete and re-create the resource if necessary)")
	applyCmd.Flags().StringP("output", "o", "", FlagDescOutput)
	applyCmd.Flags().Bool("recursive", false, "Process the directory used in -f, --filename recursively")
	applyCmd.Flags().StringSliceP("kustomize", "k", []string{}, "Process a kustomization directory")

	applyCmd.RegisterFlagCompletionFunc("namespace", func(cmd *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
		namespaces, err := kubernetes.GetNamespaces()
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return namespaces, cobra.ShellCompDirectiveNoFileComp
	})

	applyCmd.RegisterFlagCompletionFunc("filename", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveDefault
	})
}

func runApply(cmd *cobra.Command, _ []string) error {
	filename, _ := cmd.Flags().GetString("filename")
	namespace, _ := cmd.Flags().GetString("namespace")
	dryRun, _ := cmd.Flags().GetBool("dry-run")
	serverDryRun, _ := cmd.Flags().GetBool("server-dry-run")
	validate, _ := cmd.Flags().GetBool("validate")
	force, _ := cmd.Flags().GetBool("force")
	output, _ := cmd.Flags().GetString("output")
	recursive, _ := cmd.Flags().GetBool("recursive")
	kustomize, _ := cmd.Flags().GetStringSlice("kustomize")

	args := []string{"apply"}

	// Handle filename or kustomize
	if err := addSourceArgs(&args, kustomize, filename, recursive); err != nil {
		return err
	}

	// Add optional flags
	addApplyFlags(&args, namespace, serverDryRun, dryRun, validate, force, output)

	// Execute kubectl apply
	result, err := kubernetes.ExecuteKubectl(args...)
	if err != nil {
		return fmt.Errorf("failed to apply configuration: %v", err)
	}

	fmt.Print(result)
	return nil
}

func addSourceArgs(args *[]string, kustomize []string, filename string, recursive bool) error {
	if len(kustomize) > 0 {
		return addKustomizeArgs(args, kustomize)
	}
	if filename != "" {
		return addFilenameArgs(args, filename, recursive)
	}
	return fmt.Errorf("must specify either --filename (-f) or --kustomize (-k)")
}

func addKustomizeArgs(args *[]string, kustomize []string) error {
	for _, k := range kustomize {
		*args = append(*args, "-k", k)
	}
	return nil
}

func addFilenameArgs(args *[]string, filename string, recursive bool) error {
	if err := validateFilePath(filename); err != nil {
		return err
	}

	fileInfo, err := os.Stat(filename)
	if err != nil {
		return fmt.Errorf("error checking file info: %v", err)
	}

	if fileInfo.IsDir() {
		return addDirectoryArgs(args, filename, recursive)
	}
	return addFileArgs(args, filename)
}

func validateFilePath(filename string) error {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file or directory '%s' does not exist", filename)
		}
		return fmt.Errorf("error checking file '%s': %v", filename, err)
	}
	return nil
}

func addDirectoryArgs(args *[]string, filename string, recursive bool) error {
	if !recursive {
		return fmt.Errorf("'%s' is a directory, use --recursive flag to process it", filename)
	}
	*args = append(*args, "-f", filename, "--recursive")
	return nil
}

func addFileArgs(args *[]string, filename string) error {
	ext := filepath.Ext(filename)
	if ext != ".yaml" && ext != ".yml" && ext != ".json" {
		fmt.Printf("⚠️  Warning: file '%s' does not have a typical Kubernetes manifest extension (.yaml, .yml, .json)\n", filename)
	}
	*args = append(*args, "-f", filename)
	return nil
}

func addApplyFlags(args *[]string, namespace string, serverDryRun, dryRun, validate, force bool, output string) {
	if namespace != "" {
		*args = append(*args, "-n", namespace)
	}
	if serverDryRun {
		*args = append(*args, "--dry-run=server")
	} else if dryRun {
		*args = append(*args, "--dry-run=client")
	}
	if !validate {
		*args = append(*args, "--validate=false")
	}
	if force {
		*args = append(*args, "--force")
	}
	if output != "" {
		*args = append(*args, "-o", output)
	}
}
