package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check cluster health and issues",
	Long:  `Run various health checks on the cluster`,
}

var checkErrorsCmd = &cobra.Command{
	Use:     "errors",
	Aliases: []string{"err", "error"},
	Short:   "Find pods with errors",
	Long:    `Find all pods that are not in Running or Completed state (errors, crashloops, pending, etc.)`,
	RunE:    runCheckErrors,
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.AddCommand(checkErrorsCmd)
}

func runCheckErrors(_ *cobra.Command, _ []string) error {
	fmt.Println("Checking for pods with errors across all namespaces...")
	fmt.Println("(Excluding: Running, Completed)")
	fmt.Println()

	// Get all pods with wide output
	kubectlArgs := []string{"get", "pods", flagAllNamespaces, "-o", "wide"}

	kubectlCmd := exec.Command("kubectl", kubectlArgs...)
	output, err := kubectlCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to execute kubectl: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) == 0 {
		fmt.Println("No pods found.")
		return nil
	}

	// Print header
	header := lines[0]
	fmt.Println(header)

	// Filter and print non-running/non-completed pods
	foundIssues := false
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		// Check if line contains Running or Completed (case insensitive)
		lowerLine := strings.ToLower(line)
		if strings.Contains(lowerLine, "running") || strings.Contains(lowerLine, "completed") {
			continue
		}

		// Print problematic pod
		fmt.Println(line)
		foundIssues = true
	}

	if !foundIssues {
		fmt.Println()
		fmt.Println("✓ No problematic pods found! All pods are Running or Completed.")
	} else {
		fmt.Println()
		fmt.Println("⚠ Found pods with issues. Common states to investigate:")
		fmt.Println("  - CrashLoopBackOff: Pod is repeatedly crashing")
		fmt.Println("  - Error: Pod encountered an error")
		fmt.Println("  - Pending: Pod cannot be scheduled")
		fmt.Println("  - ImagePullBackOff: Cannot pull container image")
		fmt.Println("  - CreateContainerError: Container creation failed")
		fmt.Println()
		fmt.Println("Use 'kcsi logs -n <namespace> <pod>' to investigate further")
		fmt.Println("Use 'kcsi describe pod -n <namespace> <pod>' for detailed information")
	}

	return nil
}
