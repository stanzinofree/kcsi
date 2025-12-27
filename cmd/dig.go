package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
)

var digCmd = &cobra.Command{
	Use:   "dig [namespace] [pod] [domain]",
	Short: "DNS debugging inside a pod",
	Long:  "Run DNS queries (dig command) inside a pod to debug DNS resolution issues",
	Args:  cobra.MinimumNArgs(2),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) == 0 {
			// First arg: namespace
			namespaces, err := kubernetes.GetNamespaces()
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}
			return namespaces, cobra.ShellCompDirectiveNoFileComp
		}
		if len(args) == 1 {
			// Second arg: pod in that namespace
			pods, err := kubernetes.GetPods(args[0])
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}
			return pods, cobra.ShellCompDirectiveNoFileComp
		}
		// Third arg: domain name - no completion
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: runDig,
}

func init() {
	rootCmd.AddCommand(digCmd)
	digCmd.Flags().StringP("container", "c", "", "Container name for multi-container pods")
	digCmd.RegisterFlagCompletionFunc("container", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) >= 2 {
			namespace := args[0]
			podName := args[1]
			containers, err := kubernetes.GetContainers(namespace, podName)
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}
			return containers, cobra.ShellCompDirectiveNoFileComp
		}
		return nil, cobra.ShellCompDirectiveNoFileComp
	})
}

func runDig(cmd *cobra.Command, args []string) error {
	namespace := args[0]
	podName := args[1]

	// Domain to query (optional)
	domain := ""
	if len(args) >= 3 {
		domain = args[2]
	}

	// Additional arguments
	additionalArgs := ""
	if len(args) > 3 {
		for i := 3; i < len(args); i++ {
			additionalArgs = fmt.Sprintf("%s %s", additionalArgs, args[i])
		}
	}

	// Check if container flag is set
	container, _ := cmd.Flags().GetString("container")

	// Try dig, nslookup, host in order with fallback
	dnsCommand := buildDNSCommandWithFallback(domain, additionalArgs)

	kubectlArgs := []string{"exec", "-n", namespace, "-it", podName}
	if container != "" {
		kubectlArgs = append(kubectlArgs, "-c", container)
	}
	kubectlArgs = append(kubectlArgs, "--", "sh", "-c", dnsCommand)

	return kubernetes.ExecuteKubectlInteractive(kubectlArgs...)
}

// buildDNSCommandWithFallback creates a command that tries dig, then nslookup, then host
func buildDNSCommandWithFallback(domain, additionalArgs string) string {
	dnsTools := buildDNSToolCommands(domain, additionalArgs)
	return createFallbackScript(dnsTools)
}

func buildDNSToolCommands(domain, additionalArgs string) map[string]string {
	tools := map[string]string{
		"dig":      "dig",
		"nslookup": "nslookup",
		"host":     "host",
	}

	if domain != "" {
		tools["dig"] = fmt.Sprintf("dig %s%s", domain, additionalArgs)
		tools["nslookup"] = fmt.Sprintf("nslookup %s", domain)
		tools["host"] = fmt.Sprintf("host %s", domain)
	}

	return tools
}

func createFallbackScript(dnsTools map[string]string) string {
	const errorMessage = `
  echo ""
  echo "╔═══════════════════════════════════════════════════════════════════════════╗"
  echo "║ ERROR: No DNS debugging tools found in this container                    ║"
  echo "╚═══════════════════════════════════════════════════════════════════════════╝"
  echo ""
  echo "This container doesn't have dig, nslookup, or host installed."
  echo ""
  echo "To install DNS tools in this pod, run one of these commands:"
  echo ""
  echo "  Debian/Ubuntu:  apt-get update && apt-get install -y dnsutils"
  echo "  Alpine Linux:   apk add --no-cache bind-tools"
  echo "  RHEL/CentOS:    yum install -y bind-utils"
  echo ""
  echo "Alternatively, use 'kubectl debug' to attach a debug container with tools:"
  echo "  kubectl debug -it <pod> --image=nicolaka/netshoot -n <namespace>"
  echo ""
  exit 0`

	return fmt.Sprintf(
		`if command -v dig >/dev/null 2>&1; then
  %s
elif command -v nslookup >/dev/null 2>&1; then
  echo "Note: 'dig' not found, using 'nslookup' instead"
  echo ""
  %s
elif command -v host >/dev/null 2>&1; then
  echo "Note: 'dig' and 'nslookup' not found, using 'host' instead"
  echo ""
  %s
else%s
fi`,
		dnsTools["dig"], dnsTools["nslookup"], dnsTools["host"], errorMessage,
	)
}
