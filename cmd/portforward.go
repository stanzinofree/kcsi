package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/completion"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
)

var (
	portForwardNamespace string
)

var portForwardCmd = &cobra.Command{
	Use:   "port-forward [pod-name] [local-port:remote-port]",
	Short: "Forward one or more local ports to a pod",
	Long: `Forward one or more local ports to a pod with intelligent autocompletion.

Examples:
  # Forward local port 8080 to pod port 80
  kcsi port-forward -n default my-pod 8080:80
  
  # Forward local port 80 to pod port 8080 (requires root for ports < 1024)
  sudo kcsi port-forward -n production web-server 80:8080`,
	Args:                  cobra.ExactArgs(2),
	ValidArgsFunction:     portForwardCompletion,
	DisableFlagsInUseLine: true,
	RunE:                  runPortForward,
}

func init() {
	rootCmd.AddCommand(portForwardCmd)
	portForwardCmd.Flags().StringVarP(&portForwardNamespace, "namespace", "n", "", FlagDescNamespace)
	portForwardCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)
}

func portForwardCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) == 0 {
		// First arg: pod name
		return completion.PodCompletion(cmd, args, toComplete)
	}
	// Second arg: port mapping (no completion)
	return nil, cobra.ShellCompDirectiveNoFileComp
}

func runPortForward(_ *cobra.Command, args []string) error {
	podName := args[0]
	portMapping := args[1]

	// Parse port mapping (format: localPort:remotePort)
	var localPort, remotePort int
	_, err := fmt.Sscanf(portMapping, "%d:%d", &localPort, &remotePort)
	if err != nil {
		return fmt.Errorf("invalid port format '%s', expected format: localPort:remotePort (e.g., 8080:80)", portMapping)
	}

	// Validate ports
	if localPort < 1 || localPort > 65535 {
		return fmt.Errorf("invalid local port %d, must be between 1 and 65535", localPort)
	}
	if remotePort < 1 || remotePort > 65535 {
		return fmt.Errorf("invalid remote port %d, must be between 1 and 65535", remotePort)
	}

	// Check if running as root for privileged ports
	if localPort < 1024 && os.Geteuid() != 0 {
		return fmt.Errorf("local port %d requires root privileges (ports < 1024)\nTry: sudo kcsi port-forward -n %s %s %s",
			localPort, portForwardNamespace, podName, portMapping)
	}

	// Check if local port is already in use
	if isPortInUse(localPort) {
		return fmt.Errorf("local port %d is already in use\nPlease choose a different port or stop the process using it", localPort)
	}

	// Build kubectl command
	kubectlArgs := []string{"port-forward"}
	if portForwardNamespace != "" {
		kubectlArgs = append(kubectlArgs, "-n", portForwardNamespace)
	}
	kubectlArgs = append(kubectlArgs, podName, portMapping)

	fmt.Printf("Forwarding from 127.0.0.1:%d -> %s:%d\n", localPort, podName, remotePort)
	fmt.Printf("Press Ctrl+C to stop port forwarding\n\n")

	return kubernetes.ExecuteKubectlInteractive(kubectlArgs...)
}

// isPortInUse checks if a local port is already in use
func isPortInUse(port int) bool {
	// Try to listen on the port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		// If we can't listen, the port is in use
		return true
	}
	// Port is free, close the listener
	listener.Close()
	return false
}
