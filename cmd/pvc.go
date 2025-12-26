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

func runPVCPods(cmd *cobra.Command, _ []string) error {
	namespace, _ := cmd.Flags().GetString("namespace")
	outputFormat, _ := cmd.Flags().GetString("output")

	if outputFormat != "" {
		return executeKubectlWithFormat("pvc", namespace, outputFormat)
	}

	pvcOutput, err := fetchPVCsForDisplay(namespace)
	if err != nil {
		return err
	}

	podsJSON, err := fetchPodsJSONForPVC(namespace)
	if err != nil {
		return err
	}

	pvcToPods := buildPVCToPodMapping(podsJSON)
	displayPVCsWithPods(pvcOutput, pvcToPods)
	return nil
}

func executeKubectlWithFormat(resource, namespace, outputFormat string) error {
	kubectlArgs := []string{"get", resource}
	if namespace != "" {
		kubectlArgs = append(kubectlArgs, "-n", namespace)
	} else {
		kubectlArgs = append(kubectlArgs, flagAllNamespaces)
	}
	kubectlArgs = append(kubectlArgs, "-o", outputFormat)
	return kubernetes.ExecuteKubectlInteractive(kubectlArgs...)
}

func fetchPVCsForDisplay(namespace string) (string, error) {
	pvcArgs := []string{"get", "pvc", "-o", "custom-columns=NAMESPACE:metadata.namespace,NAME:metadata.name,STATUS:status.phase,VOLUME:spec.volumeName,CAPACITY:status.capacity.storage,STORAGECLASS:spec.storageClassName"}
	if namespace != "" {
		pvcArgs = append(pvcArgs, "-n", namespace)
	} else {
		pvcArgs = append(pvcArgs, flagAllNamespaces)
	}

	output, err := kubernetes.ExecuteKubectl(pvcArgs...)
	if err != nil {
		return "", fmt.Errorf("failed to get PVCs: %v", err)
	}
	return output, nil
}

func fetchPodsJSONForPVC(namespace string) (string, error) {
	podArgs := []string{"get", "pods", "-o", "json"}
	if namespace != "" {
		podArgs = append(podArgs, "-n", namespace)
	} else {
		podArgs = append(podArgs, flagAllNamespaces)
	}

	output, err := kubernetes.ExecuteKubectl(podArgs...)
	if err != nil {
		return "", fmt.Errorf("failed to get pods: %v", err)
	}
	return output, nil
}

func displayPVCsWithPods(pvcOutput string, pvcToPods map[string][]string) {
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
}

func runPVCUnbound(cmd *cobra.Command, _ []string) error {
	namespace, _ := cmd.Flags().GetString("namespace")
	outputFormat, _ := cmd.Flags().GetString("output")

	if outputFormat != "" {
		return executeKubectlWithFormat("pvc", namespace, outputFormat)
	}

	output, err := fetchUnboundPVCs(namespace)
	if err != nil {
		return err
	}

	displayUnboundPVCs(output)
	return nil
}

func fetchUnboundPVCs(namespace string) (string, error) {
	pvcArgs := []string{"get", "pvc"}
	if namespace != "" {
		pvcArgs = append(pvcArgs, "-n", namespace)
	} else {
		pvcArgs = append(pvcArgs, flagAllNamespaces)
	}

	pvcArgs = append(pvcArgs, "-o", "custom-columns=NAMESPACE:metadata.namespace,NAME:metadata.name,STATUS:status.phase,VOLUME:spec.volumeName,CAPACITY:spec.resources.requests.storage,STORAGECLASS:spec.storageClassName,AGE:metadata.creationTimestamp")

	output, err := kubernetes.ExecuteKubectl(pvcArgs...)
	if err != nil {
		return "", fmt.Errorf("failed to get PVCs: %v", err)
	}
	return output, nil
}

func displayUnboundPVCs(output string) {
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
