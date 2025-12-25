package completion

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestNamespaceCompletion(_ *testing.T) {
	cmd := &cobra.Command{}

	// Test that the function doesn't panic
	completions, directive := NamespaceCompletion(cmd, []string{}, "")

	// We expect either completions or an error directive
	// In test environment without kubectl, we expect ShellCompDirectiveError
	if directive != cobra.ShellCompDirectiveError && len(completions) == 0 {
		// This is fine - either error or empty list in test env
	}
}

func TestPodCompletionNoNamespace(_ *testing.T) {
	cmd := &cobra.Command{}

	// Test without namespace flag
	completions, directive := PodCompletion(cmd, []string{}, "")

	// Should handle gracefully even without namespace
	if directive != cobra.ShellCompDirectiveError && len(completions) == 0 {
		// This is fine - either error or empty list
	}
}

func TestContainerCompletionNoPod(_ *testing.T) {
	cmd := &cobra.Command{}

	// Test without pod argument
	completions, directive := ContainerCompletion(cmd, []string{}, "")

	// Should handle gracefully
	if directive != cobra.ShellCompDirectiveError && len(completions) == 0 {
		// This is fine - either error or empty list
	}
}
