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

	cmd := exec.CommandContext(ctx, "kubectl", "version", "--client", "--short")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		// Check if kubectl is not found
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("timeout")
		}
		return "", fmt.Errorf("not found or error: %v", err)
	}

	// Parse output - format: "Client Version: v1.28.0" or just "v1.28.0"
	output := strings.TrimSpace(out.String())
	if output == "" {
		return "", fmt.Errorf("empty output")
	}

	// Handle both old and new kubectl version output formats
	if strings.Contains(output, "Client Version:") {
		parts := strings.Split(output, "Client Version:")
		if len(parts) > 1 {
			return strings.TrimSpace(parts[1]), nil
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
