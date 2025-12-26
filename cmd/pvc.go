package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
)

var pvcCmd = &cobra.Command{
	Use:     "pvc",
	Aliases: []string{"pvcs", "persistentvolumeclaim", "persistentvolumeclaims"},
	Short:   "PVC management commands",
	Long:    "Commands for managing and inspecting PersistentVolumeClaims",
}

var pvcPodsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Show PVCs with their associated pods",
	Long:  "Display all PVCs with the pods that are using them",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: runPVCPods,
}

var pvcUnboundCmd = &cobra.Command{
	Use:   "unbound",
	Short: "Show unbound PVCs",
	Long:  "Display PVCs that are not in Bound status (Pending, Lost, etc.)",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: runPVCUnbound,
}

func init() {
	getCmd.AddCommand(pvcCmd)
	pvcCmd.AddCommand(pvcPodsCmd)
	pvcCmd.AddCommand(pvcUnboundCmd)

	// Add namespace and output flags to both subcommands
	for _, cmd := range []*cobra.Command{pvcPodsCmd, pvcUnboundCmd} {
		cmd.Flags().StringP("namespace", "n", "", "Namespace to query (default: all namespaces)")
		cmd.Flags().StringP("output", "o", "", "Output format (wide, yaml, json)")
		cmd.RegisterFlagCompletionFunc("namespace", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			namespaces, err := kubernetes.GetNamespaces()
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}
			return namespaces, cobra.ShellCompDirectiveNoFileComp
		})
	}
}

func runPVCPods(cmd *cobra.Command, args []string) error {
	namespace, _ := cmd.Flags().GetString("namespace")
	outputFormat, _ := cmd.Flags().GetString("output")

	// If output format is specified, use kubectl directly
	if outputFormat != "" {
		kubectlArgs := []string{"get", "pvc"}
		if namespace != "" {
			kubectlArgs = append(kubectlArgs, "-n", namespace)
		} else {
			kubectlArgs = append(kubectlArgs, "--all-namespaces")
		}
		kubectlArgs = append(kubectlArgs, "-o", outputFormat)
		return kubernetes.ExecuteKubectlInteractive(kubectlArgs...)
	}

	// Get PVCs
	pvcArgs := []string{"get", "pvc", "-o", "custom-columns=NAMESPACE:metadata.namespace,NAME:metadata.name,STATUS:status.phase,VOLUME:spec.volumeName,CAPACITY:status.capacity.storage,STORAGECLASS:spec.storageClassName"}
	if namespace != "" {
		pvcArgs = append(pvcArgs, "-n", namespace)
	} else {
		pvcArgs = append(pvcArgs, "--all-namespaces")
	}

	pvcOutput, err := kubernetes.ExecuteKubectl(pvcArgs...)
	if err != nil {
		return fmt.Errorf("failed to get PVCs: %v", err)
	}

	// Get pods with volume info
	podArgs := []string{"get", "pods", "-o", "json"}
	if namespace != "" {
		podArgs = append(podArgs, "-n", namespace)
	} else {
		podArgs = append(podArgs, "--all-namespaces")
	}

	podsJSON, err := kubernetes.ExecuteKubectl(podArgs...)
	if err != nil {
		return fmt.Errorf("failed to get pods: %v", err)
	}

	// Build PVC to Pod mapping
	pvcToPods := buildPVCToPodMapping(podsJSON)

	// Display results
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "NAMESPACE\tPVC\tSTATUS\tCAPACITY\tSTORAGECLASS\tUSED BY PODS")
	fmt.Fprintln(w, "---------\t---\t------\t--------\t------------\t-------------")

	pvcLines := strings.Split(strings.TrimSpace(pvcOutput), "\n")
	for i, line := range pvcLines {
		if i == 0 || line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) >= 6 {
			ns := fields[0]
			name := fields[1]
			status := fields[2]
			capacity := fields[4]
			storageClass := fields[5]

			key := fmt.Sprintf("%s/%s", ns, name)
			pods := pvcToPods[key]
			if len(pods) == 0 {
				pods = []string{"-"}
			}

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
				ns, name, status, capacity, storageClass, strings.Join(pods, ", "))
		}
	}

	w.Flush()
	return nil
}

func runPVCUnbound(cmd *cobra.Command, args []string) error {
	namespace, _ := cmd.Flags().GetString("namespace")
	outputFormat, _ := cmd.Flags().GetString("output")

	// Get all PVCs
	pvcArgs := []string{"get", "pvc"}
	if namespace != "" {
		pvcArgs = append(pvcArgs, "-n", namespace)
	} else {
		pvcArgs = append(pvcArgs, "--all-namespaces")
	}

	// If output format specified, add it
	if outputFormat != "" {
		pvcArgs = append(pvcArgs, "-o", outputFormat)
		return kubernetes.ExecuteKubectlInteractive(pvcArgs...)
	}

	// Custom output to show only unbound PVCs
	pvcArgs = append(pvcArgs, "-o", "custom-columns=NAMESPACE:metadata.namespace,NAME:metadata.name,STATUS:status.phase,VOLUME:spec.volumeName,CAPACITY:spec.resources.requests.storage,STORAGECLASS:spec.storageClassName,AGE:metadata.creationTimestamp")

	output, err := kubernetes.ExecuteKubectl(pvcArgs...)
	if err != nil {
		return fmt.Errorf("failed to get PVCs: %v", err)
	}

	// Filter for non-Bound PVCs
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "NAMESPACE\tNAME\tSTATUS\tCAPACITY\tSTORAGECLASS\tAGE")
	fmt.Fprintln(w, "---------\t----\t------\t--------\t------------\t---")

	lines := strings.Split(strings.TrimSpace(output), "\n")
	unboundCount := 0

	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) >= 6 {
			status := fields[2]
			// Show anything that's NOT "Bound"
			if status != "Bound" {
				unboundCount++
				ns := fields[0]
				name := fields[1]
				capacity := fields[4]
				storageClass := fields[5]
				age := "N/A"
				if len(fields) > 6 {
					age = fields[6]
				}

				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
					ns, name, status, capacity, storageClass, age)
			}
		}
	}

	w.Flush()

	if unboundCount == 0 {
		fmt.Println()
		fmt.Println("✅ All PVCs are bound!")
	} else {
		fmt.Println()
		fmt.Printf("⚠️  Found %d unbound PVC(s)\n", unboundCount)
	}

	return nil
}

// buildPVCToPodMapping creates a map of PVC (namespace/name) to list of pod names
func buildPVCToPodMapping(podsJSON string) map[string][]string {
	pvcToPods := make(map[string][]string)

	// Simple JSON parsing - look for persistentVolumeClaim references
	lines := strings.Split(podsJSON, "\n")
	var currentNamespace, currentPod, currentPVC string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Get namespace
		if strings.Contains(line, `"namespace":`) {
			parts := strings.Split(line, `"`)
			if len(parts) >= 4 {
				currentNamespace = parts[3]
			}
		}

		// Get pod name
		if strings.Contains(line, `"name":`) && currentNamespace != "" {
			parts := strings.Split(line, `"`)
			if len(parts) >= 4 {
				currentPod = parts[3]
			}
		}

		// Get PVC claim name
		if strings.Contains(line, `"claimName":`) && currentNamespace != "" && currentPod != "" {
			parts := strings.Split(line, `"`)
			if len(parts) >= 4 {
				currentPVC = parts[3]
				key := fmt.Sprintf("%s/%s", currentNamespace, currentPVC)
				pvcToPods[key] = append(pvcToPods[key], currentPod)
			}
		}
	}

	return pvcToPods
}
