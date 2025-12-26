package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
)

var debugCmd = &cobra.Command{
	Use:   "debug [namespace] [pod]",
	Short: "Attach ephemeral debug container to a pod",
	Long: `Attach an ephemeral debug container with networking tools to a running pod.
Automatically checks internet connectivity and selects appropriate debug image.

Common debug images:
  - nicolaka/netshoot (full toolkit, requires internet)
  - busybox (minimal, widely cached)
  - alpine (lightweight with package manager)`,
	Args: cobra.MinimumNArgs(2),
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
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: runDebug,
}

func init() {
	rootCmd.AddCommand(debugCmd)
	debugCmd.Flags().StringP("image", "i", "", "Debug image to use (default: auto-detect based on connectivity)")
	debugCmd.Flags().StringP("container", "c", "", "Target container name (for multi-container pods)")
	debugCmd.RegisterFlagCompletionFunc("container", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
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

func runDebug(cmd *cobra.Command, args []string) error {
	namespace := args[0]
	podName := args[1]

	// Get user-specified image or auto-detect
	image, _ := cmd.Flags().GetString("image")
	container, _ := cmd.Flags().GetString("container")

	if image == "" {
		fmt.Println("🔍 Checking internet connectivity from cluster...")
		image = selectDebugImage(namespace)
	}

	fmt.Printf("🐛 Creating debug session for pod '%s' in namespace '%s'\n", podName, namespace)
	fmt.Printf("📦 Using debug image: %s\n", image)
	fmt.Println()

	// Build kubectl debug command with proper flags for interactive shell
	kubectlArgs := []string{"debug", "-it", podName, "-n", namespace, "--image=" + image}

	// Add target container if specified
	if container != "" {
		kubectlArgs = append(kubectlArgs, "--target="+container)
	}

	// Use --share-processes and provide explicit command
	// This ensures we get an interactive shell
	kubectlArgs = append(kubectlArgs, "--share-processes", "--")
	
	// Detect and run the best available shell
	kubectlArgs = append(kubectlArgs, "sh", "-c", 
		"if command -v bash >/dev/null 2>&1; then exec bash; elif command -v zsh >/dev/null 2>&1; then exec zsh; else exec sh; fi")

	return kubernetes.ExecuteKubectlInteractive(kubectlArgs...)
}

// selectDebugImage checks internet connectivity and returns appropriate debug image
func selectDebugImage(namespace string) string {
	// Try to check internet connectivity by testing if we can resolve DNS
	// We'll run a quick test pod or check existing pods' connectivity
	
	fmt.Println("  Testing internet connectivity...")
	
	// Try to get a node and check if it can pull images
	nodesOutput, err := kubernetes.ExecuteKubectl("get", "nodes", "-o", "jsonpath={.items[0].metadata.name}")
	if err != nil {
		fmt.Println("  ⚠️  Could not check nodes, using safe fallback image")
		return "busybox:latest"
	}

	nodeName := strings.TrimSpace(nodesOutput)
	
	// Check if common debug images are already present by looking at existing pods
	imagesOutput, err := kubernetes.ExecuteKubectl("get", "pods", "--all-namespaces", 
		"-o", "jsonpath={.items[*].spec.containers[*].image}")
	
	if err == nil {
		images := strings.ToLower(imagesOutput)
		
		// If netshoot is already used in the cluster, it's likely available
		if strings.Contains(images, "nicolaka/netshoot") {
			fmt.Println("  ✅ Found netshoot image in cluster (full debugging toolkit)")
			return "nicolaka/netshoot:latest"
		}
		
		// Check for alpine
		if strings.Contains(images, "alpine") {
			fmt.Println("  ✅ Found alpine image in cluster (lightweight with package manager)")
			return "alpine:latest"
		}
	}

	// Try to check internet by testing if we can reach a public registry
	// Create a temporary test using kubectl run with --rm
	fmt.Println("  Testing public registry access...")
	testOutput, err := kubernetes.ExecuteKubectl("run", "kcsi-connectivity-test", 
		"--image=busybox:latest", "--rm", "-i", "--restart=Never", 
		"--command", "--", "echo", "connected")
	
	if err == nil && strings.Contains(testOutput, "connected") {
		fmt.Println("  ✅ Internet connectivity confirmed (using full debug toolkit)")
		// Clean up test pod if it exists
		kubernetes.ExecuteKubectl("delete", "pod", "kcsi-connectivity-test", "--ignore-not-found=true")
		return "nicolaka/netshoot:latest"
	}

	// Fallback: use busybox (minimal but widely cached)
	fmt.Printf("  ⚠️  Limited connectivity detected on node '%s'\n", nodeName)
	fmt.Println("  📦 Using busybox (minimal tools, widely cached)")
	fmt.Println()
	fmt.Println("  Tip: If you need more tools, specify an image with -i flag:")
	fmt.Println("       kcsi debug <ns> <pod> -i nicolaka/netshoot")
	fmt.Println()
	
	return "busybox:latest"
}
