package context

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	kcsiDir         = ".kcsi"
	contextsFile    = "contexts.yaml"
	contextsSubdir  = "contexts"
	kubeconfigName  = "kube.config"
	currentFileName = ".current"
)

// Context represents a kcsi context configuration
type Context struct {
	Name             string `yaml:"name"`
	KubeconfigPath   string `yaml:"kubeconfig_path"`
	Description      string `yaml:"description,omitempty"`
	DefaultNamespace string `yaml:"default_namespace,omitempty"`
}

// Config represents the contexts configuration file
type Config struct {
	Contexts       []Context `yaml:"contexts"`
	CurrentContext string    `yaml:"current_context,omitempty"`
}

// GetKcsiDir returns the kcsi configuration directory path
func GetKcsiDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(home, kcsiDir), nil
}

// GetContextsFilePath returns the path to contexts.yaml
func GetContextsFilePath() (string, error) {
	kcsiDir, err := GetKcsiDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(kcsiDir, contextsFile), nil
}

// GetContextDir returns the directory path for a specific context
func GetContextDir(name string) (string, error) {
	kcsiDir, err := GetKcsiDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(kcsiDir, contextsSubdir, name), nil
}

// GetContextKubeconfigPath returns the kubeconfig path for a specific context
func GetContextKubeconfigPath(name string) (string, error) {
	contextDir, err := GetContextDir(name)
	if err != nil {
		return "", err
	}
	return filepath.Join(contextDir, kubeconfigName), nil
}

// InitializeKcsiDir creates the kcsi directory structure if it doesn't exist
func InitializeKcsiDir() error {
	kcsiDir, err := GetKcsiDir()
	if err != nil {
		return err
	}

	// Create main kcsi directory
	if err := os.MkdirAll(kcsiDir, 0755); err != nil {
		return fmt.Errorf("failed to create kcsi directory: %w", err)
	}

	// Create contexts subdirectory
	contextsDir := filepath.Join(kcsiDir, contextsSubdir)
	if err := os.MkdirAll(contextsDir, 0755); err != nil {
		return fmt.Errorf("failed to create contexts directory: %w", err)
	}

	// Create contexts.yaml if it doesn't exist
	contextsFilePath, err := GetContextsFilePath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(contextsFilePath); os.IsNotExist(err) {
		config := Config{
			Contexts: []Context{},
		}
		if err := SaveConfig(&config); err != nil {
			return fmt.Errorf("failed to create contexts.yaml: %w", err)
		}
	}

	return nil
}

// LoadConfig loads the contexts configuration from contexts.yaml
func LoadConfig() (*Config, error) {
	contextsFilePath, err := GetContextsFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(contextsFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Return empty config if file doesn't exist
			return &Config{Contexts: []Context{}}, nil
		}
		return nil, fmt.Errorf("failed to read contexts.yaml: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse contexts.yaml: %w", err)
	}

	return &config, nil
}

// SaveConfig saves the contexts configuration to contexts.yaml
func SaveConfig(config *Config) error {
	contextsFilePath, err := GetContextsFilePath()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(contextsFilePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write contexts.yaml: %w", err)
	}

	return nil
}

// AddContext adds a new context to the configuration
func AddContext(name, kubeconfigPath, description string) error {
	if err := InitializeKcsiDir(); err != nil {
		return err
	}

	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// Check if context already exists
	for _, ctx := range config.Contexts {
		if ctx.Name == name {
			return fmt.Errorf("context '%s' already exists", name)
		}
	}

	// Add new context
	newContext := Context{
		Name:           name,
		KubeconfigPath: kubeconfigPath,
		Description:    description,
	}

	config.Contexts = append(config.Contexts, newContext)

	return SaveConfig(config)
}

// ImportContext imports a kubeconfig file into kcsi's managed directory
func ImportContext(name, sourceKubeconfigPath, description string) error {
	if err := InitializeKcsiDir(); err != nil {
		return err
	}

	// Check if source kubeconfig exists
	if _, err := os.Stat(sourceKubeconfigPath); os.IsNotExist(err) {
		return fmt.Errorf("kubeconfig file not found: %s", sourceKubeconfigPath)
	}

	// Create context directory
	contextDir, err := GetContextDir(name)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(contextDir, 0755); err != nil {
		return fmt.Errorf("failed to create context directory: %w", err)
	}

	// Copy kubeconfig file
	sourceData, err := os.ReadFile(sourceKubeconfigPath)
	if err != nil {
		return fmt.Errorf("failed to read source kubeconfig: %w", err)
	}

	destPath, err := GetContextKubeconfigPath(name)
	if err != nil {
		return err
	}

	if err := os.WriteFile(destPath, sourceData, 0600); err != nil {
		return fmt.Errorf("failed to write kubeconfig: %w", err)
	}

	// Add context to configuration
	return AddContext(name, destPath, description)
}

// RemoveContext removes a context from the configuration and deletes its files
func RemoveContext(name string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// Find and remove context
	found := false
	newContexts := []Context{}
	for _, ctx := range config.Contexts {
		if ctx.Name != name {
			newContexts = append(newContexts, ctx)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("context '%s' not found", name)
	}

	config.Contexts = newContexts

	// Clear current context if it was the removed one
	if config.CurrentContext == name {
		config.CurrentContext = ""
	}

	// Remove context directory
	contextDir, err := GetContextDir(name)
	if err == nil {
		os.RemoveAll(contextDir)
	}

	return SaveConfig(config)
}

// GetContext returns a specific context by name
func GetContext(name string) (*Context, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	for _, ctx := range config.Contexts {
		if ctx.Name == name {
			return &ctx, nil
		}
	}

	return nil, fmt.Errorf("context '%s' not found", name)
}

// ListContexts returns all available contexts
func ListContexts() ([]Context, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	return config.Contexts, nil
}

// SetCurrentContext sets the current active context
func SetCurrentContext(name string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// Verify context exists
	found := false
	for _, ctx := range config.Contexts {
		if ctx.Name == name {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("context '%s' not found", name)
	}

	config.CurrentContext = name
	return SaveConfig(config)
}

// GetCurrentContext returns the current active context
func GetCurrentContext() (*Context, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	if config.CurrentContext == "" {
		return nil, fmt.Errorf("no current context set")
	}

	return GetContext(config.CurrentContext)
}

// GetCurrentContextName returns the name of the current context
func GetCurrentContextName() (string, error) {
	config, err := LoadConfig()
	if err != nil {
		return "", err
	}

	return config.CurrentContext, nil
}

// SetDefaultNamespace sets the default namespace for a specific context
func SetDefaultNamespace(contextName, namespace string) error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// Find the context and update its default namespace
	found := false
	for i, ctx := range config.Contexts {
		if ctx.Name == contextName {
			config.Contexts[i].DefaultNamespace = namespace
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("context '%s' not found", contextName)
	}

	return SaveConfig(config)
}

// ClearDefaultNamespace removes the default namespace from a specific context
func ClearDefaultNamespace(contextName string) error {
	return SetDefaultNamespace(contextName, "")
}

// GetDefaultNamespace returns the default namespace for a specific context
func GetDefaultNamespace(contextName string) (string, error) {
	ctx, err := GetContext(contextName)
	if err != nil {
		return "", err
	}

	return ctx.DefaultNamespace, nil
}

// GetCurrentDefaultNamespace returns the default namespace for the current active context
func GetCurrentDefaultNamespace() (string, error) {
	ctx, err := GetCurrentContext()
	if err != nil {
		return "", err
	}

	return ctx.DefaultNamespace, nil
}
