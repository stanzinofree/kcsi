package cmd

const (
	flagAllNamespaces = "--all-namespaces"
	imageBusybox      = "busybox:latest"

	// Flag descriptions
	FlagDescNamespace   = "Kubernetes namespace"
	FlagDescSkipConfirm = "Skip confirmation prompt"
	FlagDescOutput      = "Output format (json, yaml, wide, etc.)"

	// Error messages
	ErrNamespaceRequired = "namespace is required (use -n flag)"
)
