package kubernetes

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// JSONPath constants for kubectl queries
const (
	jsonPathMetadataName   = "jsonpath={.items[*].metadata.name}"
	jsonPathContainerNames = "jsonpath={.spec.containers[*].name}"
	flagAllNamespaces      = "--all-namespaces"
)

// ExecuteKubectl runs a kubectl command and returns the output
func ExecuteKubectl(args ...string) (string, error) {
	cmd := exec.Command("kubectl", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("kubectl error: %v - %s", err, stderr.String())
	}

	return out.String(), nil
}

// ExecuteKubectlInteractive runs a kubectl command with stdin/stdout/stderr attached
func ExecuteKubectlInteractive(args ...string) error {
	cmd := exec.Command("kubectl", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("kubectl error: %v", err)
	}

	return nil
}

// GetNamespaces returns a list of all namespaces in the cluster
func GetNamespaces() ([]string, error) {
	output, err := ExecuteKubectl("get", "namespaces", "-o", jsonPathMetadataName)
	if err != nil {
		return nil, err
	}

	// Split the space-separated namespace names
	namespaces := strings.Fields(strings.TrimSpace(output))
	return namespaces, nil
}

// GetPods returns a list of pods in the specified namespace
func GetPods(namespace string) ([]string, error) {
	args := []string{"get", "pods", "-o", jsonPathMetadataName}
	if namespace != "" {
		args = append(args, "-n", namespace)
	} else {
		args = append(args, flagAllNamespaces)
	}

	output, err := ExecuteKubectl(args...)
	if err != nil {
		return nil, err
	}

	// Split the space-separated pod names
	pods := strings.Fields(strings.TrimSpace(output))
	return pods, nil
}

// GetContainers returns a list of container names in a specific pod
func GetContainers(namespace, podName string) ([]string, error) {
	if podName == "" {
		return nil, fmt.Errorf("pod name is required")
	}

	args := []string{"get", "pod", podName, "-o", jsonPathContainerNames}
	if namespace != "" {
		args = append(args, "-n", namespace)
	}

	output, err := ExecuteKubectl(args...)
	if err != nil {
		return nil, err
	}

	// Split the space-separated container names
	containers := strings.Fields(strings.TrimSpace(output))
	return containers, nil
}

// GetServices returns a list of services in the specified namespace
func GetServices(namespace string) ([]string, error) {
	args := []string{"get", "services", "-o", jsonPathMetadataName}
	if namespace != "" {
		args = append(args, "-n", namespace)
	} else {
		args = append(args, flagAllNamespaces)
	}

	output, err := ExecuteKubectl(args...)
	if err != nil {
		return nil, err
	}

	services := strings.Fields(strings.TrimSpace(output))
	return services, nil
}

// GetDeployments returns a list of deployments in the specified namespace
func GetDeployments(namespace string) ([]string, error) {
	args := []string{"get", "deployments", "-o", jsonPathMetadataName}
	if namespace != "" {
		args = append(args, "-n", namespace)
	} else {
		args = append(args, flagAllNamespaces)
	}

	output, err := ExecuteKubectl(args...)
	if err != nil {
		return nil, err
	}

	deployments := strings.Fields(strings.TrimSpace(output))
	return deployments, nil
}

// GetNodes returns a list of nodes in the cluster
func GetNodes() ([]string, error) {
	output, err := ExecuteKubectl("get", "nodes", "-o", jsonPathMetadataName)
	if err != nil {
		return nil, err
	}

	nodes := strings.Fields(strings.TrimSpace(output))
	return nodes, nil
}

// GetConfigMaps returns a list of configmaps in the specified namespace
func GetConfigMaps(namespace string) ([]string, error) {
	args := []string{"get", "configmaps", "-o", jsonPathMetadataName}
	if namespace != "" {
		args = append(args, "-n", namespace)
	} else {
		args = append(args, flagAllNamespaces)
	}

	output, err := ExecuteKubectl(args...)
	if err != nil {
		return nil, err
	}

	configmaps := strings.Fields(strings.TrimSpace(output))
	return configmaps, nil
}

// GetSecrets returns a list of secrets in the specified namespace
func GetSecrets(namespace string) ([]string, error) {
	args := []string{"get", "secrets", "-o", jsonPathMetadataName}
	if namespace != "" {
		args = append(args, "-n", namespace)
	} else {
		args = append(args, flagAllNamespaces)
	}

	output, err := ExecuteKubectl(args...)
	if err != nil {
		return nil, err
	}

	secrets := strings.Fields(strings.TrimSpace(output))
	return secrets, nil
}

// GetDaemonSets returns a list of daemonsets in the specified namespace
func GetDaemonSets(namespace string) ([]string, error) {
	args := []string{"get", "daemonsets", "-o", jsonPathMetadataName}
	if namespace != "" {
		args = append(args, "-n", namespace)
	} else {
		args = append(args, flagAllNamespaces)
	}

	output, err := ExecuteKubectl(args...)
	if err != nil {
		return nil, err
	}

	daemonsets := strings.Fields(strings.TrimSpace(output))
	return daemonsets, nil
}

// GetStatefulSets returns a list of statefulsets in the specified namespace
func GetStatefulSets(namespace string) ([]string, error) {
	args := []string{"get", "statefulsets", "-o", jsonPathMetadataName}
	if namespace != "" {
		args = append(args, "-n", namespace)
	} else {
		args = append(args, flagAllNamespaces)
	}

	output, err := ExecuteKubectl(args...)
	if err != nil {
		return nil, err
	}

	statefulsets := strings.Fields(strings.TrimSpace(output))
	return statefulsets, nil
}

// GetKubectlVersion returns the kubectl client version with timeout
func GetKubectlVersion() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Try new format first (kubectl 1.28+): kubectl version --client
	cmd := exec.CommandContext(ctx, "kubectl", "version", "--client")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		// If new format fails, try legacy format: kubectl version --client --short
		ctx2, cancel2 := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel2()

		cmd2 := exec.CommandContext(ctx2, "kubectl", "version", "--client", "--short")
		out.Reset()
		stderr.Reset()
		cmd2.Stdout = &out
		cmd2.Stderr = &stderr

		err2 := cmd2.Run()
		if err2 != nil {
			// Check if kubectl is not found
			if ctx2.Err() == context.DeadlineExceeded {
				return "", fmt.Errorf("timeout")
			}
			// Return stderr for debugging if available
			if stderr.Len() > 0 {
				return "", fmt.Errorf("error: %s", strings.TrimSpace(stderr.String()))
			}
			return "", fmt.Errorf("not found")
		}
	}

	// Parse output
	output := strings.TrimSpace(out.String())
	if output == "" {
		return "", fmt.Errorf("empty output")
	}

	// Handle different kubectl version output formats:
	// 1. New format (JSON): {"clientVersion": {"gitVersion": "v1.31.0", ...}}
	// 2. Old format: "Client Version: v1.28.0"
	// 3. Very old format: "v1.28.0"

	// Try to extract version from JSON format (kubectl 1.28+)
	if strings.Contains(output, "\"gitVersion\"") {
		// Extract version using simple string parsing
		if start := strings.Index(output, "\"gitVersion\":"); start != -1 {
			start += len("\"gitVersion\":")
			// Skip whitespace and opening quote
			for start < len(output) && (output[start] == ' ' || output[start] == '"') {
				start++
			}
			// Find closing quote
			end := start
			for end < len(output) && output[end] != '"' {
				end++
			}
			if end > start {
				return output[start:end], nil
			}
		}
	}

	// Handle old text formats
	if strings.Contains(output, "Client Version:") {
		// Extract just the first line with "Client Version:"
		lines := strings.Split(output, "\n")
		for _, line := range lines {
			if strings.Contains(line, "Client Version:") {
				parts := strings.Split(line, "Client Version:")
				if len(parts) > 1 {
					return strings.TrimSpace(parts[1]), nil
				}
			}
		}
	}

	// Return first line if it looks like a version (and only that line)
	lines := strings.Split(output, "\n")
	if len(lines) > 0 {
		firstLine := strings.TrimSpace(lines[0])
		if strings.HasPrefix(firstLine, "v") || strings.HasPrefix(firstLine, "Client Version:") {
			// Extract just the version part
			if strings.HasPrefix(firstLine, "Client Version:") {
				parts := strings.Split(firstLine, ":")
				if len(parts) > 1 {
					return strings.TrimSpace(parts[1]), nil
				}
			}
			return firstLine, nil
		}
	}

	// If all else fails, just return the first word that looks like a version
	fields := strings.Fields(output)
	for _, field := range fields {
		if strings.HasPrefix(field, "v") && len(field) > 1 {
			return field, nil
		}
	}

	return output, nil
}

// GetClusterInfo returns cluster info with timeout (for --cluster flag)
func GetClusterInfo() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "kubectl", "cluster-info")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("timeout (cluster unreachable)")
		}
		return "", fmt.Errorf("unreachable: %v", err)
	}

	return out.String(), nil
}

// GetCurrentContext returns the current kubeconfig context
func GetCurrentContext() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "kubectl", "config", "current-context")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "unknown", fmt.Errorf("unable to determine context")
	}

	return strings.TrimSpace(out.String()), nil
}

// GetCurrentNamespace returns the current namespace from the context
func GetCurrentNamespace() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "kubectl", "config", "view", "--minify", "--output", "jsonpath={..namespace}")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "default", nil // Default to "default" namespace if unable to determine
	}

	namespace := strings.TrimSpace(out.String())
	if namespace == "" {
		return "default", nil
	}

	return namespace, nil
}

// GetKubectlPath returns the path to kubectl binary (best effort)
func GetKubectlPath() string {
	// Try to find kubectl in PATH using exec.LookPath
	path, err := exec.LookPath("kubectl")
	if err != nil {
		return "not found"
	}
	return path
}
