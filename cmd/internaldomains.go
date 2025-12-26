package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/stanzinofree/kcsi/pkg/kubernetes"
)

var internalDomainsCmd = &cobra.Command{
	Use:     "internal-domains",
	Aliases: []string{"idomains", "idom"},
	Short:   "List internal Kubernetes FQDNs for services and pods",
	Long:    "Display all internal DNS names (FQDNs) for services and pods in the cluster",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
	RunE: runInternalDomains,
}

func init() {
	getCmd.AddCommand(internalDomainsCmd)
	internalDomainsCmd.Flags().StringP("namespace", "n", "", "Namespace to list internal domains from")
	internalDomainsCmd.RegisterFlagCompletionFunc("namespace", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		namespaces, err := kubernetes.GetNamespaces()
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return namespaces, cobra.ShellCompDirectiveNoFileComp
	})
}

func runInternalDomains(cmd *cobra.Command, args []string) error {
	namespace, _ := cmd.Flags().GetString("namespace")

	// Get services
	servicesArgs := []string{"get", "services", "-o", "custom-columns=TYPE:metadata.labels,NAME:metadata.name,NAMESPACE:metadata.namespace,CLUSTER-IP:spec.clusterIP,PORTS:spec.ports[*].port"}
	if namespace != "" {
		servicesArgs = append(servicesArgs, "-n", namespace)
	} else {
		servicesArgs = append(servicesArgs, "--all-namespaces")
	}

	servicesOutput, err := kubernetes.ExecuteKubectl(servicesArgs...)
	if err != nil {
		return fmt.Errorf("failed to get services: %v", err)
	}

	// Get pods
	podsArgs := []string{"get", "pods", "-o", "custom-columns=TYPE:metadata.labels,NAME:metadata.name,NAMESPACE:metadata.namespace,IP:status.podIP,STATUS:status.phase"}
	if namespace != "" {
		podsArgs = append(podsArgs, "-n", namespace)
	} else {
		podsArgs = append(podsArgs, "--all-namespaces")
	}

	podsOutput, err := kubernetes.ExecuteKubectl(podsArgs...)
	if err != nil {
		return fmt.Errorf("failed to get pods: %v", err)
	}

	// Display results
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "TYPE\tNAME\tNAMESPACE\tFQDN\tIP\tINFO")
	fmt.Fprintln(w, "----\t----\t---------\t----\t--\t----")

	// Process services
	serviceLines := strings.Split(strings.TrimSpace(servicesOutput), "\n")
	for i, line := range serviceLines {
		if i == 0 || line == "" {
			continue // Skip header and empty lines
		}
		fields := strings.Fields(line)
		if len(fields) >= 4 {
			name := fields[1]
			ns := fields[2]
			clusterIP := fields[3]
			ports := ""
			if len(fields) > 4 {
				ports = fields[4]
			}

			fqdn := fmt.Sprintf("%s.%s.svc.cluster.local", name, ns)
			fmt.Fprintf(w, "SERVICE\t%s\t%s\t%s\t%s\t%s\n", name, ns, fqdn, clusterIP, ports)
		}
	}

	// Process pods
	podLines := strings.Split(strings.TrimSpace(podsOutput), "\n")
	for i, line := range podLines {
		if i == 0 || line == "" {
			continue // Skip header and empty lines
		}
		fields := strings.Fields(line)
		if len(fields) >= 4 {
			name := fields[1]
			ns := fields[2]
			podIP := fields[3]
			status := ""
			if len(fields) > 4 {
				status = fields[4]
			}

			// Pod FQDN format: pod-ip-with-dashes.namespace.pod.cluster.local
			podIPDashes := strings.ReplaceAll(podIP, ".", "-")
			fqdn := fmt.Sprintf("%s.%s.pod.cluster.local", podIPDashes, ns)
			fmt.Fprintf(w, "POD\t%s\t%s\t%s\t%s\t%s\n", name, ns, fqdn, podIP, status)
		}
	}

	w.Flush()
	return nil
}
