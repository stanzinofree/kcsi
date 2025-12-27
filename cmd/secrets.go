package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
)

var secretsCmd = &cobra.Command{
	Use:     "secrets",
	Aliases: []string{"secret"},
	Short:   "Secrets management commands",
	Long:    "Commands for viewing and managing Kubernetes secrets",
}

var secretsDecodedCmd = &cobra.Command{
	Use:               "decoded [name]",
	Short:             "Show decoded secret data",
	Long:              "Display all keys and values of a secret with base64 decoding applied",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: secretNameCompletion,
	RunE:              runSecretsDecoded,
}

var secretsShowCmd = &cobra.Command{
	Use:               "show [name]",
	Short:             "Show specific secret key value",
	Long:              "Display the decoded value of a specific key in a secret",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: secretNameCompletion,
	RunE:              runSecretsShow,
}

func secretNameCompletion(cmd *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
	if len(args) == 0 {
		namespace, _ := cmd.Flags().GetString("namespace")
		if namespace == "" {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		secrets, err := kubernetes.GetSecrets(namespace)
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return secrets, cobra.ShellCompDirectiveNoFileComp
	}
	return nil, cobra.ShellCompDirectiveNoFileComp
}

func init() {
	getCmd.AddCommand(secretsCmd)
	secretsCmd.AddCommand(secretsDecodedCmd)
	secretsCmd.AddCommand(secretsShowCmd)

	// Add namespace flag to both subcommands
	for _, cmd := range []*cobra.Command{secretsDecodedCmd, secretsShowCmd} {
		cmd.Flags().StringP("namespace", "n", "", FlagDescNamespace)
		cmd.RegisterFlagCompletionFunc("namespace", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			namespaces, err := kubernetes.GetNamespaces()
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}
			return namespaces, cobra.ShellCompDirectiveNoFileComp
		})
	}

	// Add key flag to show command
	secretsShowCmd.Flags().StringP("key", "k", "", "Secret key to display")
	secretsShowCmd.MarkFlagRequired("key")
}

func runSecretsDecoded(cmd *cobra.Command, args []string) error {
	namespace, secretName, err := validateSecretsArgs(cmd, args)
	if err != nil {
		return err
	}

	printSecurityWarning(true)

	data, err := fetchSecretData(secretName, namespace)
	if err != nil {
		return err
	}

	displayAllSecretKeys(secretName, namespace, data)
	return nil
}

func runSecretsShow(cmd *cobra.Command, args []string) error {
	namespace, secretName, err := validateSecretsArgs(cmd, args)
	if err != nil {
		return err
	}

	key, _ := cmd.Flags().GetString("key")
	printSecurityWarning(false)

	data, err := fetchSecretData(secretName, namespace)
	if err != nil {
		return err
	}

	value, err := getSecretKeyValue(data, key)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", value)
	return nil
}

func validateSecretsArgs(cmd *cobra.Command, args []string) (namespace, secretName string, err error) {
	namespace, _ = cmd.Flags().GetString("namespace")
	if namespace == "" {
		return "", "", fmt.Errorf("namespace is required (use -n flag)")
	}
	secretName = args[0]
	return namespace, secretName, nil
}

func printSecurityWarning(verbose bool) {
	if verbose {
		fmt.Println("⚠️  Warning: Secret values will be displayed in plain text")
		fmt.Println("   Make sure your terminal is not being shared or recorded")
		fmt.Println()
	} else {
		fmt.Fprintln(os.Stderr, "⚠️  Displaying secret in plain text")
	}
}

func fetchSecretData(secretName, namespace string) (map[string]interface{}, error) {
	output, err := kubernetes.ExecuteKubectl("get", "secret", secretName, "-n", namespace, "-o", "json")
	if err != nil {
		return nil, fmt.Errorf("failed to get secret: %v", err)
	}

	var secret map[string]interface{}
	if err := json.Unmarshal([]byte(output), &secret); err != nil {
		return nil, fmt.Errorf("failed to parse secret JSON: %v", err)
	}

	data, ok := secret["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("secret has no data field")
	}

	return data, nil
}

func getSecretKeyValue(data map[string]interface{}, key string) (string, error) {
	encodedValue, ok := data[key].(string)
	if !ok {
		return "", fmt.Errorf("key '%s' not found in secret", key)
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(encodedValue)
	if err != nil {
		return "", fmt.Errorf("failed to decode value: %v", err)
	}

	return string(decodedBytes), nil
}

func displayAllSecretKeys(secretName, namespace string, data map[string]interface{}) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, "Secret: %s (namespace: %s)\n\n", secretName, namespace)
	fmt.Fprintln(w, "KEY\tVALUE")
	fmt.Fprintln(w, "---\t-----")

	for key, value := range data {
		encodedValue, ok := value.(string)
		if !ok {
			fmt.Fprintf(w, "%s\t<invalid format>\n", key)
			continue
		}

		decodedBytes, err := base64.StdEncoding.DecodeString(encodedValue)
		if err != nil {
			fmt.Fprintf(w, "%s\t<decode error: %v>\n", key, err)
			continue
		}

		fmt.Fprintf(w, "%s\t%s\n", key, string(decodedBytes))
	}

	w.Flush()
}
