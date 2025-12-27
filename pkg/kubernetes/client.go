package kubernetes

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
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
