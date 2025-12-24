package kubernetes

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
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

// GetNamespaces returns a list of all namespaces in the cluster
func GetNamespaces() ([]string, error) {
	output, err := ExecuteKubectl("get", "namespaces", "-o", "jsonpath={.items[*].metadata.name}")
	if err != nil {
		return nil, err
	}

	// Split the space-separated namespace names
	namespaces := strings.Fields(strings.TrimSpace(output))
	return namespaces, nil
}

// GetPods returns a list of pods in the specified namespace
func GetPods(namespace string) ([]string, error) {
	args := []string{"get", "pods", "-o", "jsonpath={.items[*].metadata.name}"}
	if namespace != "" {
		args = append(args, "-n", namespace)
	} else {
		args = append(args, "--all-namespaces")
	}

	output, err := ExecuteKubectl(args...)
	if err != nil {
		return nil, err
	}

	// Split the space-separated pod names
	pods := strings.Fields(strings.TrimSpace(output))
	return pods, nil
}
