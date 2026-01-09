# KCSI - Kubectl for Humans

A kubectl wrapper that makes Kubernetes operations faster, safer, and more intuitive through intelligent TAB autocompletion and safety guardrails. KCSI eliminates context switching by providing cascading autocompletion (namespace → resource → pod → container) and confirmation prompts for destructive operations.

## Tech Stack

| Layer | Technology | Version | Purpose |
|-------|------------|---------|---------|
| Language | Go | 1.25.5 | Chosen for cross-platform compilation, single binary distribution, and excellent CLI tooling |
| CLI Framework | Cobra | 1.10.2 | Industry-standard CLI framework with built-in autocompletion support |
| YAML Parsing | gopkg.in/yaml.v3 | 3.0.1 | For context configuration management and version manifest |
| Build Tool | Taskfile | 3 | Modern task runner with dependency management, replacing Makefiles |
| Testing | Go standard | built-in | go test with race detection enabled in CI |
| CI/CD | GitHub Actions | - | Automated testing, builds, and releases with matrix builds for multiple platforms |

## Quick Start

```bash
# Prerequisites
- Go 1.25.5+ (for building from source)
- kubectl installed and configured
- Active Kubernetes cluster access

# Clone and build
git clone https://github.com/stanzinofree/kcsi.git
cd kcsi
task build

# Install locally
task install

# Verify installation
kcsi version

# Enable shell completion (choose your shell)
# Bash
echo 'source <(kcsi completion bash)' >> ~/.bashrc
# Zsh
echo 'source <(kcsi completion zsh)' >> ~/.zshrc
# Fish
kcsi completion fish > ~/.config/fish/completions/kcsi.fish

# Quick test - try TAB autocompletion
kcsi logs -n <TAB>  # Shows namespaces
kcsi get pods -n default <TAB>  # Shows pods in default namespace
```

## Project Structure

```
kcsi/
├── cmd/                      # Command implementations (Cobra commands)
│   ├── root.go              # Root command and CLI setup
│   ├── logs.go              # Log streaming command
│   ├── get.go               # Get resources (pods, services, etc.)
│   ├── delete.go            # Delete with confirmation prompts
│   ├── context.go           # Multi-cluster context management
│   ├── debug.go             # Ephemeral debug containers
│   ├── diag.go              # Diagnostics report generation
│   ├── events.go            # Cluster event monitoring
│   ├── apply.go             # Apply manifests
│   ├── rollout.go           # Deployment rollouts
│   ├── secrets.go           # Secret management
│   ├── pvc.go               # PVC operations
│   ├── portforward.go       # Port forwarding
│   └── [25 command files]   # One file per major command
│
├── pkg/                      # Reusable packages
│   ├── completion/          # Shell autocompletion logic
│   │   └── completion.go    # Functions for namespace, pod, container completion
│   ├── context/             # Multi-cluster context management
│   │   └── context.go       # Context storage, switching, and isolation
│   ├── kubernetes/          # Kubernetes client operations
│   │   └── client.go        # kubectl command execution with context awareness
│   └── version/             # Version information
│       ├── version.go       # Runtime version getters
│       └── version.yaml     # Single source of truth for version info
│
├── docs/                     # Documentation website (GitHub Pages)
│   └── assets/              # CSS/JS for docs
├── scripts/                  # Utility scripts
│   └── git-hooks/           # Git hooks for development
├── .github/                  # GitHub configuration
│   └── workflows/           # CI/CD pipelines
│       ├── build.yml        # Build and test on push/PR
│       ├── release.yml      # Release automation
│       └── pages.yml        # Documentation deployment
│
├── main.go                   # Entry point (calls cmd.Execute())
├── Taskfile.yml             # Task automation definitions
├── go.mod                   # Go module dependencies
└── build.sh                 # Cross-platform build script
```

## Architecture Overview

KCSI follows a clean **wrapper architecture** where it acts as an intelligent proxy to kubectl:

```
┌─────────────────────────────────────────────────────────┐
│                     User Shell                          │
└─────────────────────────┬───────────────────────────────┘
                          │
                          │ kcsi logs -n prod <TAB>
                          ▼
┌─────────────────────────────────────────────────────────┐
│                  KCSI CLI (Cobra)                       │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  │
│  │  Commands    │  │  Completion  │  │   Context    │  │
│  │  (cmd/)      │◀─│  (pkg/)      │◀─│  (pkg/)      │  │
│  └──────────────┘  └──────────────┘  └──────────────┘  │
└─────────────────────────┬───────────────────────────────┘
                          │
                          │ ExecuteKubectl("get pods -n prod")
                          │ with KUBECONFIG env set
                          ▼
┌─────────────────────────────────────────────────────────┐
│                   kubectl binary                        │
└─────────────────────────┬───────────────────────────────┘
                          │
                          ▼
┌─────────────────────────────────────────────────────────┐
│              Kubernetes API Server                      │
└─────────────────────────────────────────────────────────┘
```

**Key Design Principles:**

1. **Wrapper Pattern**: KCSI doesn't reimplement Kubernetes logic. It constructs kubectl commands and executes them, ensuring compatibility with all kubectl versions.

2. **Context Isolation**: Multi-cluster support via isolated context management in `~/.kcsi/contexts/`, never modifying system kubeconfig.

3. **Progressive Enhancement**: Commands work exactly like kubectl but with added autocompletion and safety features. Users can drop `kcsi` and use `kubectl` at any time.

4. **Command-per-File**: Each major command (logs, get, delete, etc.) lives in its own file under `cmd/`, making the codebase easy to navigate and extend.

### Key Modules

| Module | Location | Purpose |
|--------|----------|---------|
| Command Layer | cmd/*.go | Cobra command definitions with flags, validation, and execution |
| Completion Engine | pkg/completion/ | Shell autocompletion functions querying K8s resources in real-time |
| Context Manager | pkg/context/ | Multi-cluster context storage and switching without touching system kubeconfig |
| Kubernetes Client | pkg/kubernetes/ | kubectl command execution with context-aware KUBECONFIG environment injection |
| Version System | pkg/version/ | Single source of truth (version.yaml) with runtime getters for build info |

## Development Guidelines

### Code Style

**File Naming:**
- Command files: lowercase with descriptive names (`logs.go`, `portforward.go`, `context.go`)
- Package files: lowercase, descriptive (`completion.go`, `client.go`, `context.go`)
- Test files: `*_test.go` suffix, co-located with source

**Code Naming:**
- Exported functions: PascalCase (`GetNamespaces`, `ExecuteKubectl`)
- Unexported functions: camelCase (`runLogs`, `askForConfirmation`, `setKubeconfigEnv`)
- Variables: camelCase (`namespace`, `podName`, `kubectlArgs`)
- Constants: SCREAMING_SNAKE_CASE or camelCase depending on scope (`jsonPathMetadataName`, `kcsiDir`)
- Boolean variables: descriptive predicates (`logsFollow`, `deletePodForce`, `diagCluster`)

**Cobra Command Structure Pattern:**
```go
var exampleCmd = &cobra.Command{
    Use:   "example [args]",
    Short: "Brief one-line description",
    Long:  `Detailed multi-line description with examples`,
    Args:  cobra.ExactArgs(1),
    RunE:  runExample,
    ValidArgsFunction: completion.SomeCompletion,
}

func init() {
    rootCmd.AddCommand(exampleCmd)
    exampleCmd.Flags().StringVarP(&exampleNamespace, "namespace", "n", "", FlagDescNamespace)
    exampleCmd.RegisterFlagCompletionFunc("namespace", completion.NamespaceCompletion)
}

func runExample(_ *cobra.Command, args []string) error {
    // Implementation
    return nil
}
```

### Import Order

Follow Go's standard import grouping (enforced by gofmt):

1. Standard library imports
2. External dependencies (github.com/spf13/cobra, etc.)
3. Internal package imports (github.com/stanzinofree/kcsi/pkg/*)

Example:
```go
import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "gopkg.in/yaml.v3"

    "github.com/stanzinofree/kcsi/pkg/completion"
    "github.com/stanzinofree/kcsi/pkg/kubernetes"
)
```

### Error Handling

- Use `fmt.Errorf` with `%w` verb for error wrapping to preserve context
- Return errors from RunE functions rather than calling os.Exit
- Provide actionable error messages with context

Example:
```go
if err != nil {
    return fmt.Errorf("failed to execute kubectl: %w", err)
}
```

### Testing Conventions

- Test files co-located with source files (`completion_test.go` next to `completion.go`)
- Use Go's standard testing package
- Tests run with `-race` flag in CI to detect race conditions
- Focus on unit tests for completion and kubernetes packages

### Commit Message Conventions

Follow Conventional Commits format (required for changelog generation):

- `feat:` New features or enhancements
- `fix:` Bug fixes
- `docs:` Documentation changes
- `chore:` Maintenance tasks (dependency updates, version bumps)
- `test:` Test changes
- `refactor:` Code restructuring without behavior change

Examples:
```
feat: add context management for multi-cluster operations
fix: correct kubectl version parsing for 1.31+
docs: update README with context command examples
chore: bump version to 0.7.0
```

## Available Commands

| Command | Description | Common Flags |
|---------|-------------|-------------|
| `task build` | Build kcsi binary for current platform | - |
| `task build:all` | Build for all platforms (uses build.sh) | - |
| `task install` | Install kcsi to /usr/local/bin (requires sudo) | - |
| `task run -- <args>` | Build and run with arguments | Any kcsi flags |
| `task test` | Run all tests with verbose output | - |
| `task fmt` | Format all Go code with gofmt | - |
| `task vet` | Run go vet static analysis | - |
| `task lint` | Run golangci-lint (if installed) | - |
| `task check` | Run fmt + vet + test in sequence | - |
| `task clean` | Remove build artifacts | - |
| `task dev` | Quick development cycle: build + install | - |
| `task pr` | Create and push PR (requires BRANCH, TITLE, DESC) | - |
| `task tag` | Create git tag (requires VERSION, MESSAGE) | - |

### Development Workflow Example

```bash
# Make changes to code
vim cmd/logs.go

# Format and check
task fmt
task vet

# Build and test locally
task build
./kcsi logs -n default <TAB>

# Run tests
task test

# Install locally for integration testing
task install
kcsi version

# Create PR (example)
task pr BRANCH=fix/log-timestamps TITLE="Fix log timestamp parsing" DESC="Fixes issue #123"
```

## Environment Variables

| Variable | Required | Description | Example |
|----------|----------|-------------|---------|
| `KUBECONFIG` | No | Path to kubeconfig (overridden by kcsi context if active) | `~/.kube/config` |
| `KUBE_EDITOR` | No | Editor for `kcsi edit` command | `vim`, `nano`, `code` |
| `EDITOR` | No | Fallback editor if KUBE_EDITOR not set | `vim` |

**Note:** When using `kcsi context`, the KUBECONFIG environment variable is automatically set per-command to the active context's kubeconfig path. This allows seamless multi-cluster switching without modifying your system kubeconfig.

## Configuration Files

### Version Manifest (Single Source of Truth)

**Location:** `pkg/version/version.yaml`

Contains all version and metadata information used by build scripts and runtime:

```yaml
version: 0.6.4
name: kcsi
fullName: Kubectl Cli Super Intuitive
description: A kubectl wrapper with intelligent autocompletion
author: Alessandro Middei
license: MIT
repository: https://github.com/stanzinofree/kcsi
```

### Context Configuration

**Location:** `~/.kcsi/contexts.yaml` (auto-created on first `kcsi context` command)

Stores multi-cluster context configuration:

```yaml
contexts:
  - name: prod
    kubeconfig_path: ~/.kcsi/contexts/prod/kube.config
    description: Production cluster
  - name: dev
    kubeconfig_path: ~/.kcsi/contexts/dev/kube.config
    description: Development cluster
current_context: prod
```

**Context Directory Structure:**
```
~/.kcsi/
├── contexts.yaml                    # Context registry
└── contexts/                        # Per-context storage
    ├── prod/
    │   └── kube.config             # Imported kubeconfig
    └── dev/
        └── kube.config
```

## Testing

- **Unit Tests:** Located in `pkg/*/` directories
  - `pkg/completion/completion_test.go` - Tests completion functions
  - `pkg/kubernetes/client_test.go` - Tests kubectl execution helpers
  - `pkg/version/version_test.go` - Tests version info retrieval

- **Test Execution:**
  ```bash
  task test                    # Run all tests
  go test -v ./...            # Direct go test invocation
  go test -race -v ./...      # With race detection (used in CI)
  ```

- **Coverage:** No explicit coverage target, but all core packages have test coverage
- **CI Testing:** GitHub Actions runs tests on push/PR with Go 'stable' version

## Deployment

### Release Process

1. **Update version:**
   ```bash
   # Edit pkg/version/version.yaml
   vim pkg/version/version.yaml
   # Change version: 0.6.4 → 0.7.0
   ```

2. **Update CHANGELOG.md:**
   ```bash
   vim CHANGELOG.md
   # Add new version section with changes
   ```

3. **Create PR and merge:**
   ```bash
   task pr BRANCH=release/v0.7.0 TITLE="chore: bump version to 0.7.0" DESC="Release v0.7.0"
   # Wait for CI, get approval, merge
   ```

4. **Tag release (after merge to main):**
   ```bash
   git checkout main && git pull
   task tag VERSION=0.7.0 MESSAGE="Release v0.7.0 with context management"
   ```

5. **GitHub Actions automatically:**
   - Builds binaries for all platforms
   - Creates GitHub Release with artifacts
   - Updates documentation site

### Manual Release (if needed)

```bash
task build:all                               # Build all platforms
task release:github VERSION=0.7.0           # Create GitHub release
```

### Build Artifacts

The build process generates binaries for:
- **macOS:** darwin-amd64, darwin-arm64
- **Linux:** linux-amd64, linux-arm64, linux-arm
- **Windows:** windows-amd64.exe, windows-arm64.exe

Artifacts are stored in `bin/` directory and uploaded to GitHub Releases.

## Key Implementation Patterns

### 1. Cascading Autocompletion

Autocompletion functions in `pkg/completion/` query the Kubernetes API in real-time:

```go
func PodCompletion(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
    namespace, _ := cmd.Flags().GetString("namespace")
    pods, err := kubernetes.GetPods(namespace)
    // Returns pod names for TAB completion
}
```

This creates the cascading effect: namespace → resource → container.

### 2. Context-Aware kubectl Execution

All kubectl commands are wrapped in `pkg/kubernetes/client.go` with automatic context injection:

```go
func setKubeconfigEnv(cmd *exec.Cmd) {
    ctx, err := kcsicontext.GetCurrentContext()
    if err == nil && ctx != nil {
        cmd.Env = append(os.Environ(), fmt.Sprintf("KUBECONFIG=%s", ctx.KubeconfigPath))
    }
}
```

Every kubectl execution checks for an active kcsi context and sets KUBECONFIG accordingly.

### 3. Safety Guardrails

Delete operations require explicit confirmation:

```go
func askForConfirmation(resourceType, resourceName, namespace string) bool {
    fmt.Printf("Are you sure you want to delete %s '%s' in namespace '%s'? [y/N]: ",
               resourceType, resourceName, namespace)
    // Read user input, return true only for "y" or "yes"
}
```

Use `--force` flag to skip confirmation prompts in scripts.

### 4. Progressive Enhancement

Commands construct kubectl arguments and execute them directly:

```go
kubectlArgs := []string{"logs"}
if namespace != "" {
    kubectlArgs = append(kubectlArgs, "-n", namespace)
}
kubectlArgs = append(kubectlArgs, podName)

cmd := exec.Command("kubectl", kubectlArgs...)
cmd.Stdout = os.Stdout  // Stream directly to user
cmd.Run()
```

This ensures 100% kubectl compatibility while adding KCSI features on top.

## Additional Resources

- **Documentation:** https://stanzinofree.github.io/kcsi/
- **Cheatsheet:** https://stanzinofree.github.io/kcsi/cheatsheet.html
- **Roadmap:** https://stanzinofree.github.io/kcsi/roadmap.html
- **GitHub Repository:** https://github.com/stanzinofree/kcsi
- **Issue Tracker:** https://github.com/stanzinofree/kcsi/issues
- **License:** MIT (see LICENSE file)
- **Security:** docs/SECURITY_SECRETS.md for secret handling guidance

## Contributing

Contributions are welcome! See @README.md for PR workflow and repository rules.

**Quick Contribution Guide:**
1. Fork the repository
2. Create a feature branch from main
3. Make changes following code style guidelines
4. Run `task check` to validate
5. Submit PR with descriptive title and body
6. Wait for CI checks and review

**Repository Protection:**
- Direct pushes to main are blocked
- All PRs require passing CI (tests, CodeQL, SonarCloud)
- PRs must be reviewed before merge


## Skill Usage Guide

When working on tasks involving these technologies, invoke the corresponding skill:

| Skill | Invoke When |
|-------|-------------|
| yaml | YAML parsing and serialization with gopkg.in/yaml.v3 |
| shell-scripting | Shell completion scripts for bash, zsh, fish, and PowerShell |
| cobra | Cobra CLI framework, commands, flags, and shell completion |
| taskfile | Taskfile task automation, build orchestration, and dependencies |
| github-actions | GitHub Actions CI/CD workflows, matrix builds, and automation |
| kubernetes | Kubernetes API, kubectl integration, and cluster operations |
| go | Go language, goroutines, standard library, and CLI patterns |
