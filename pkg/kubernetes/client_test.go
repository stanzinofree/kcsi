package kubernetes

import (
	"testing"
)

func TestExecuteKubectlInvalidCommand(t *testing.T) {
	// Test that invalid commands return an error
	_, err := ExecuteKubectl("invalid-command-that-does-not-exist")
	if err == nil {
		t.Error("Expected error for invalid kubectl command, got nil")
	}
}

func TestGetContainersEmptyPodName(t *testing.T) {
	// Test that empty pod name returns an error
	_, err := GetContainers("default", "")
	if err == nil {
		t.Error("Expected error for empty pod name, got nil")
	}
	if err.Error() != "pod name is required" {
		t.Errorf("Expected 'pod name is required' error, got: %v", err)
	}
}
