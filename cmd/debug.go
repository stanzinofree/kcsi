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
	debugCmd.Flags().BoolP("fast", "f", false, "Use lightweight busybox image for faster startup")
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
	fast, _ := cmd.Flags().GetBool("fast")

	// If fast mode, use busybox
	if fast && image == "" {
		image = imageBusybox
		fmt.Println("⚡ Fast mode: using busybox (lightweight, limited tools)")
		fmt.Println()
	} else if image == "" {
		fmt.Println("🔍 Checking internet connectivity from cluster...")
		image = selectDebugImage(namespace)
	}

	fmt.Printf("🐛 Creating debug session for pod '%s' in namespace '%s'\n", podName, namespace)
	fmt.Printf("📦 Using debug image: %s\n", image)

	// Show image size info
	if strings.Contains(image, "netshoot") {
		fmt.Println("⏳ Note: netshoot is ~400MB, first pull may take 1-2 minutes...")
		fmt.Println("   Tip: Use -f for fast mode or -i busybox for faster startup")
	}
	fmt.Println()

	// Try ephemeral container first (lightweight), fall back to copy if needed
	kubectlArgs := []string{"debug", "-it", podName, "-n", namespace,
		"--image=" + image,
		"--target=" + getPrimaryContainer(namespace, podName, container)}

	fmt.Println("🚀 Attaching ephemeral debug container...")
	fmt.Println("   (The pod will NOT be modified, debug container is temporary)")
	fmt.Println()

	return kubernetes.ExecuteKubectlInteractive(kubectlArgs...)
}

// getPrimaryContainer returns the target container name (user-specified or first container)
func getPrimaryContainer(namespace, podName, userContainer string) string {
	if userContainer != "" {
		return userContainer
	}

	// Get first container from pod
	containers, err := kubernetes.GetContainers(namespace, podName)
	if err != nil || len(containers) == 0 {
		return ""
	}

	return containers[0]
}

// selectDebugImage checks internet connectivity and returns appropriate debug image
func selectDebugImage(_ string) string {
	// Try to check internet connectivity by testing if we can resolve DNS
	// We'll run a quick test pod or check existing pods' connectivity

	fmt.Println("  Testing internet connectivity...")

	// Try to get a node and check if it can pull images
	nodesOutput, err := kubernetes.ExecuteKubectl("get", "nodes", "-o", "jsonpath={.items[0].metadata.name}")
	if err != nil {
		fmt.Println("  ⚠️  Could not check nodes, using safe fallback image")
		return imageBusybox
	}

	nodeName := strings.TrimSpace(nodesOutput)

	// Check if common debug images are already present by looking at existing pods
	imagesOutput, err := kubernetes.ExecuteKubectl("get", "pods", flagAllNamespaces,
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
		"--image="+imageBusybox, "--rm", "-i", "--restart=Never",
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

	return imageBusybox
}
