package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
	"github.com/stanzinofree/kcsi/pkg/version"
)

var (
	diagCluster bool
	diagStrict  bool
)

var diagCmd = &cobra.Command{
	Use:   "diag",
	Short: "Generate diagnostics report for GitHub issues",
	Long: `Generate a copy/paste-friendly diagnostics report for GitHub issues.

This command produces a safe diagnostic output containing:
- KCSI version information
- System details (OS, Go version)
- Kubernetes context information
- kubectl toolchain checks

The output is designed to be shared publicly without exposing secrets.`,
	RunE: runDiag,
}

func init() {
	rootCmd.AddCommand(diagCmd)
	diagCmd.Flags().BoolVar(&diagCluster, "cluster", false, "Include cluster reachability check (requires active cluster connection)")
	diagCmd.Flags().BoolVar(&diagStrict, "strict", false, "Exit with non-zero code if any check fails")
}

func runDiag(_ *cobra.Command, _ []string) error {
	var sb strings.Builder
	hasErrors := false

	// Header
	sb.WriteString("==============================================\n")
	sb.WriteString("KCSI DIAGNOSTICS\n")
	sb.WriteString("==============================================\n\n")

	// Timestamp
	sb.WriteString(fmt.Sprintf("Generated: %s\n\n", time.Now().Format("2006-01-02 15:04:05 MST")))

	// KCSI Version Info
	sb.WriteString("----------------------------------------------\n")
	sb.WriteString("KCSI Version\n")
	sb.WriteString("----------------------------------------------\n")
	manifest := version.GetManifest()
	sb.WriteString(fmt.Sprintf("  Version:    %s\n", version.GetVersion()))
	sb.WriteString(fmt.Sprintf("  Build Date: %s\n", getBuildDate(manifest.BuildDate)))
	sb.WriteString(fmt.Sprintf("  Git Commit: %s\n", "unknown")) // Will be set during build
	sb.WriteString("\n")

	// System Information
	sb.WriteString("----------------------------------------------\n")
	sb.WriteString("System Information\n")
	sb.WriteString("----------------------------------------------\n")
	sb.WriteString(fmt.Sprintf("  OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH))
	sb.WriteString(fmt.Sprintf("  Go Version: %s\n", runtime.Version()))

	// Terminal variables (safe subset)
	if term := os.Getenv("TERM"); term != "" {
		sb.WriteString(fmt.Sprintf("  TERM:       %s\n", term))
	}
	if colorTerm := os.Getenv("COLORTERM"); colorTerm != "" {
		sb.WriteString(fmt.Sprintf("  COLORTERM:  %s\n", colorTerm))
	}

	// Shell (best-effort)
	if shell := getShell(); shell != "" {
		sb.WriteString(fmt.Sprintf("  Shell:      %s\n", shell))
	}
	sb.WriteString("\n")

	// Kubernetes Context (safe)
	sb.WriteString("----------------------------------------------\n")
	sb.WriteString("Kubernetes Context\n")
	sb.WriteString("----------------------------------------------\n")

	// Kubeconfig path
	kubeconfigPath := getKubeconfigPath()
	sb.WriteString(fmt.Sprintf("  Kubeconfig: %s\n", kubeconfigPath))

	// Current context
	currentContext, err := kubernetes.GetCurrentContext()
	if err != nil {
		sb.WriteString(fmt.Sprintf("  Context:    error: %v\n", err))
		hasErrors = true
	} else {
		sb.WriteString(fmt.Sprintf("  Context:    %s\n", currentContext))
	}

	// Current namespace
	currentNamespace, err := kubernetes.GetCurrentNamespace()
	if err != nil {
		sb.WriteString(fmt.Sprintf("  Namespace:  error: %v\n", err))
		hasErrors = true
	} else {
		sb.WriteString(fmt.Sprintf("  Namespace:  %s\n", currentNamespace))
	}
	sb.WriteString("\n")

	// Toolchain Checks
	sb.WriteString("----------------------------------------------\n")
	sb.WriteString("Toolchain Checks\n")
	sb.WriteString("----------------------------------------------\n")

	// kubectl version check
	kubectlVersion, err := kubernetes.GetKubectlVersion()
	if err != nil {
		sb.WriteString(fmt.Sprintf("  kubectl:    not found or error: %v\n", err))
		hasErrors = true
	} else {
		sb.WriteString(fmt.Sprintf("  kubectl:    %s\n", kubectlVersion))
	}

	// Cluster reachability (only if --cluster flag is set)
	if diagCluster {
		clusterInfo, err := kubernetes.GetClusterInfo()
		if err != nil {
			sb.WriteString(fmt.Sprintf("  Cluster:    unreachable or error: %v\n", err))
			hasErrors = true
		} else {
			sb.WriteString(fmt.Sprintf("  Cluster:    reachable\n"))
			// Show first line of cluster-info (usually the control plane endpoint)
			lines := strings.Split(strings.TrimSpace(clusterInfo), "\n")
			if len(lines) > 0 {
				sb.WriteString(fmt.Sprintf("              %s\n", lines[0]))
			}
		}
	} else {
		sb.WriteString("  Cluster:    not checked (use --cluster to enable)\n")
	}
	sb.WriteString("\n")

	// KCSI Config (safe)
	sb.WriteString("----------------------------------------------\n")
	sb.WriteString("KCSI Configuration\n")
	sb.WriteString("----------------------------------------------\n")

	// Check for config directories (best-effort)
	homeDir, err := os.UserHomeDir()
	if err == nil {
		kcsiConfigDir := fmt.Sprintf("%s/.kcsi", homeDir)
		if _, err := os.Stat(kcsiConfigDir); err == nil {
			sb.WriteString(fmt.Sprintf("  Config Dir: %s (exists)\n", kcsiConfigDir))
		} else {
			sb.WriteString(fmt.Sprintf("  Config Dir: %s (not found)\n", kcsiConfigDir))
		}
	} else {
		sb.WriteString("  Config Dir: unable to determine home directory\n")
	}
	sb.WriteString("\n")

	// Footer - Safety Warning
	sb.WriteString("==============================================\n")
	sb.WriteString("SAFETY & NEXT STEPS\n")
	sb.WriteString("==============================================\n\n")
	sb.WriteString("⚠️  Safety: This command does not print secrets, but\n")
	sb.WriteString("    review the output before sharing logs publicly.\n\n")
	sb.WriteString("📝  Next Steps:\n")
	sb.WriteString("    1. Open an issue: https://github.com/stanzinofree/kcsi/issues\n")
	sb.WriteString("    2. Paste this output in the issue description\n")
	sb.WriteString("    3. Sponsors can request label: sponsor-priority\n\n")
	sb.WriteString("💰  Sponsor KCSI:\n")
	sb.WriteString("    - GitHub Sponsors: https://github.com/sponsors/stanzinofree\n")
	sb.WriteString("    - Buy Me a Coffee: https://buymeacoffee.com/smilzao\n")
	sb.WriteString("\n")

	// Print the complete diagnostics report
	fmt.Print(sb.String())

	// Exit with error if --strict is set and there were errors
	if diagStrict && hasErrors {
		return fmt.Errorf("diagnostics completed with errors (--strict mode)")
	}

	return nil
}

// getBuildDate returns the build date or "unknown" if not set
func getBuildDate(buildDate string) string {
	if buildDate == "" {
		return "unknown"
	}
	return buildDate
}

// getKubeconfigPath returns the kubeconfig path being used
func getKubeconfigPath() string {
	if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
		return kubeconfig
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "~/.kube/config (unable to expand ~)"
	}

	return fmt.Sprintf("%s/.kube/config (default)", homeDir)
}

// getShell returns the current shell (best-effort)
func getShell() string {
	// Try SHELL environment variable first
	if shell := os.Getenv("SHELL"); shell != "" {
		return shell
	}

	// Try to detect shell on Windows
	if runtime.GOOS == "windows" {
		// Check for PowerShell
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		cmd := exec.CommandContext(ctx, "powershell", "-Command", "$PSVersionTable.PSVersion.ToString()")
		if output, err := cmd.Output(); err == nil {
			return fmt.Sprintf("PowerShell %s", strings.TrimSpace(string(output)))
		}

		return "cmd.exe (assumed)"
	}

	return ""
}
