# kcsi - Kubectl Smart Interactive Wrapper

A kubectl wrapper with intelligent autocompletion for namespaces, pods, and other Kubernetes resources.

## Features

- Smart autocompletion for Kubernetes namespaces
- Interactive resource selection
- Simplified kubectl commands with tab completion
- Built with Go and Cobra

## Current Status

**Version:** 0.2.0 - Phase 2 Complete

Currently implemented:
- `kcsi get pods -n <namespace>` with namespace autocompletion
- `kcsi describe pod -n <namespace> <pod>` with cascading autocompletion
- `kcsi logs -n <namespace> <pod>` with cascading autocompletion and full flags support
- Container autocompletion for multi-container pods

## Installation

### Prerequisites

- Go 1.21 or higher
- kubectl installed and configured
- Access to a Kubernetes cluster
- (Optional) [Task](https://taskfile.dev) for simplified build commands

### Download Pre-built Binaries

Pre-built binaries are available for multiple platforms in the `bin/` directory after running the build script:

- macOS (Intel): `kcsi-darwin-amd64`
- macOS (Apple Silicon): `kcsi-darwin-arm64`
- Linux (x86_64): `kcsi-linux-amd64`
- Linux (ARM64): `kcsi-linux-arm64`
- Linux (ARM): `kcsi-linux-arm`

### Build from source

#### Quick build (current platform)
```bash
git clone <repository-url>
cd kcsi
go build -o kcsi
```

#### Multi-platform build

Using Task (recommended):
```bash
# See all available tasks
task

# Build for all platforms
task build:all

# Build for specific platform
task build:linux-arm64

# Clean and release build
task release
```

Using build script:
```bash
# Build for all supported platforms
./build.sh

# Binaries will be created in the bin/ directory
```

### Install

Using Task:
```bash
# Build and install in one step
task install

# Uninstall
task uninstall
```

Manual installation:
```bash
# Using pre-built binary (example for Linux ARM64)
sudo cp bin/kcsi-linux-arm64 /usr/local/bin/kcsi
sudo chmod +x /usr/local/bin/kcsi

# Or using locally built binary
sudo mv kcsi /usr/local/bin/
```

## Setup Autocompletion

### Bash

```bash
# Load in current session
source <(kcsi completion bash)

# Load for all sessions (Linux)
kcsi completion bash > /etc/bash_completion.d/kcsi

# Load for all sessions (macOS with Homebrew)
kcsi completion bash > $(brew --prefix)/etc/bash_completion.d/kcsi
```

### Zsh

```bash
# Enable completion if not already enabled
echo "autoload -U compinit; compinit" >> ~/.zshrc

# Generate completion script
kcsi completion zsh > "${fpath[1]}/_kcsi"

# Restart your shell
```

### Fish

```bash
# Load in current session
kcsi completion fish | source

# Load for all sessions
kcsi completion fish > ~/.config/fish/completions/kcsi.fish
```

## Usage

### Get pods with namespace autocompletion

```bash
# Type this and press TAB after -n to see all available namespaces
kcsi get pods -n <TAB>

# Example
kcsi get pods -n kube-system
```

This is equivalent to:
```bash
kubectl get pods -n kube-system
```

### Describe pod with cascading autocompletion

```bash
# First select namespace, then pod name
kcsi describe pod -n <TAB>  # Shows namespaces
kcsi describe pod -n kube-system <TAB>  # Shows pods in kube-system

# Example
kcsi describe pod -n kube-system coredns-123456
```

This is equivalent to:
```bash
kubectl describe pod coredns-123456 -n kube-system
```

### Get logs with cascading autocompletion

```bash
# Basic usage with namespace and pod autocompletion
kcsi logs -n <TAB>  # Shows namespaces
kcsi logs -n kube-system <TAB>  # Shows pods in kube-system

# Follow logs
kcsi logs -f -n kube-system my-pod

# Get last 100 lines
kcsi logs --tail 100 -n kube-system my-pod

# Get logs from specific container (if pod has multiple containers)
kcsi logs -n kube-system my-pod -c <TAB>  # Shows containers in the pod
kcsi logs -n kube-system my-pod -c my-container

# Get previous container logs
kcsi logs -p -n kube-system my-pod
```

This is equivalent to:
```bash
kubectl logs -f my-pod -n kube-system
kubectl logs --tail 100 my-pod -n kube-system
kubectl logs my-pod -c my-container -n kube-system
kubectl logs -p my-pod -n kube-system
```

## Roadmap

### Phase 1: Proof of Concept ✅
- [x] Basic CLI structure with Cobra
- [x] `get pods` command with namespace flag
- [x] Namespace autocompletion
- [x] Completion script generation

### Phase 2: Expand Commands ✅
- [x] `logs` command with pod and namespace autocompletion
- [x] `describe pod` command with resource autocompletion
- [x] Cascading autocompletion (namespace → pod)
- [x] Container autocompletion for multi-container pods
- [x] Full `logs` flags support (-f, --tail, -p, -c)

### Phase 3: Additional Commands (Next)
- [ ] `exec` command with interactive pod selection
- [ ] `delete pod` command with confirmation
- [ ] `get services` command
- [ ] `get deployments` command
- [ ] `port-forward` command

### Phase 4: Enhancements
- [ ] Cache for faster autocompletion
- [ ] Default context/namespace configuration
- [ ] Custom aliases
- [ ] Fuzzy matching for resources
- [ ] Configuration file support

## Development

### Project Structure

```
kcsi/
├── cmd/              # Cobra commands
│   ├── root.go      # Root command setup
│   ├── get.go       # Get command implementation
│   └── completion.go # Completion script generation
├── pkg/
│   ├── kubernetes/  # Kubernetes client wrapper
│   └── completion/  # Autocompletion logic
├── Taskfile.yml      # Task runner configuration
├── build.sh          # Multi-platform build script
└── main.go          # Entry point
```

### Common Development Tasks

Using Task (recommended):
```bash
# List all available tasks
task

# Build for current platform
task build

# Build for all platforms
task build:all

# Run with arguments
task run -- get pods -n default

# Clean build artifacts
task clean

# Run tests
task test

# Format and vet code
task check

# Development workflow (build + install)
task dev

# Generate completion scripts
task completion:all

# Prepare release
task release
```

### Building

```bash
# Quick build for current platform
go build -o kcsi

# Using Task
task build

# Multi-platform build
./build.sh
# or
task build:all
```

### Testing

```bash
# Run
./kcsi get pods -n <namespace>

# Using Task
task run -- get pods -n default

# Test autocompletion (after setting up completion scripts)
./kcsi get pods -n <TAB>

# Test specific platform binary
./bin/kcsi-linux-arm64 --version
```

### Available Task Commands

| Command | Description |
|---------|-------------|
| `task` | Show all available tasks |
| `task build` | Build for current platform |
| `task build:all` | Build for all platforms |
| `task build:linux-arm64` | Build for specific platform |
| `task clean` | Clean build artifacts |
| `task install` | Build and install to /usr/local/bin |
| `task uninstall` | Uninstall from /usr/local/bin |
| `task run -- <args>` | Build and run with arguments |
| `task test` | Run tests |
| `task fmt` | Format code |
| `task vet` | Run go vet |
| `task check` | Run fmt, vet, and test |
| `task completion:all` | Generate all completion scripts |
| `task dev` | Development mode (build + install) |
| `task release` | Prepare release build |

## Contributing

This is currently a personal project in early development. Feedback and suggestions are welcome!

## License

MIT License (to be added)

## Acknowledgments

Built with:
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [kubectl](https://kubernetes.io/docs/reference/kubectl/) - Kubernetes command-line tool
