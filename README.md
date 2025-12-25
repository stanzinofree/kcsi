# kcsi - Kubectl Smart Interactive Wrapper

A kubectl wrapper with intelligent autocompletion for namespaces, pods, and other Kubernetes resources.

## Features

- Smart autocompletion for Kubernetes namespaces
- Interactive resource selection
- Simplified kubectl commands with tab completion
- Centralized version and project information
- Built with Go and Cobra

## Current Status

**Version:** 0.5.1 - Version & Project Information

Currently implemented:

**Get Commands (with -o/--output flag):**
- `kcsi get pods -o wide` - List pods with node information
- `kcsi get services -o wide` (alias: `svc`) - List services with extended info
- `kcsi get deployments -o wide` (alias: `deploy`) - List deployments
- `kcsi get namespaces` (alias: `ns`) - List namespaces
- `kcsi get nodes -o wide` (alias: `no`) - List nodes with details
- `kcsi get configmaps` (alias: `cm`) - List configmaps
- `kcsi get secrets` - List secrets
- All get commands support `-o` for output formats: wide, yaml, json, etc.

**Describe Commands:**
- `kcsi describe pod` - Describe a specific pod
- `kcsi describe service` - Describe a service
- `kcsi describe deployment` - Describe a deployment
- `kcsi describe node` - Describe a node
- `kcsi describe configmap` - Describe a configmap
- `kcsi describe secret` - Describe a secret

**Delete Commands (with safety confirmation):**
- `kcsi delete pod` - Delete a pod with confirmation prompt
- `kcsi delete service` - Delete a service with confirmation
- `kcsi delete deployment` - Delete a deployment with confirmation
- `kcsi delete configmap` - Delete a configmap with confirmation
- `kcsi delete secret` - Delete a secret with confirmation
- All delete commands support `--force` flag to skip confirmation

**Diagnostics & Monitoring:**
- `kcsi events` - Get cluster events sorted by timestamp
- `kcsi events -w` - Watch events in real-time
- `kcsi check errors` - Find all pods with issues (not Running/Completed)
- Helpful diagnostics suggestions for troubleshooting

**Other Commands:**
- `kcsi logs` - Get pod logs with full kubectl flags support (-f, --tail, -p, -c)
- Container autocompletion for multi-container pods
- Cascading autocompletion: namespace → resource → container
- `kcsi about` - Show project information, philosophy, and author details
- `kcsi --version` - Show version and author
- `kcsi --version-detailed` - Show detailed version info (Go version, OS/Arch)

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

### Get other resources

```bash
# Get services
kcsi get services -n <TAB>  # Shows namespaces
kcsi get svc -n default <TAB>  # Shows services in default namespace

# Get deployments
kcsi get deployments -n <TAB>
kcsi get deploy -n kube-system

# Get nodes (cluster-wide)
kcsi get nodes
kcsi get no  # Short alias

# Get namespaces
kcsi get namespaces
kcsi get ns  # Short alias

# Get configmaps and secrets
kcsi get configmaps -n default
kcsi get cm -n default  # Short alias
kcsi get secrets -n kube-system
```

### Describe resources

```bash
# Describe service
kcsi describe service -n <TAB>  # Shows namespaces
kcsi describe svc -n default <TAB>  # Shows services
kcsi describe service -n default my-service

# Describe deployment
kcsi describe deployment -n production <TAB>
kcsi describe deploy -n production my-app

# Describe node
kcsi describe node <TAB>  # Shows all nodes
kcsi describe node worker-1

# Describe configmap or secret
kcsi describe configmap -n default <TAB>
kcsi describe cm -n default my-config
kcsi describe secret -n default my-secret
```

### Delete resources safely

```bash
# Delete pod with confirmation prompt
kcsi delete pod -n <TAB>  # Shows namespaces
kcsi delete pod -n default <TAB>  # Shows pods in namespace
kcsi delete pod -n default my-pod
# Output: Are you sure you want to delete pod 'my-pod' in namespace 'default'? [y/N]:

# Delete with --force to skip confirmation (use with caution!)
kcsi delete pod -n default my-pod --force
kcsi delete pod -n default my-pod -f  # Short form

# Delete other resources
kcsi delete service -n default my-service
kcsi delete deployment -n production my-app
kcsi delete configmap -n default my-config
kcsi delete secret -n default my-secret

# All delete commands have autocompletion
kcsi delete svc -n <TAB>  # Namespace autocomplete
kcsi delete deploy -n prod <TAB>  # Deployment autocomplete
```

**Safety Features:**
- Confirmation prompt shows resource type, name, and namespace
- Requires explicit 'y' or 'yes' to proceed
- Use `--force` or `-f` flag to skip confirmation (for scripts/automation)
- All deletes support cascading autocompletion

### Get resources with output formats

```bash
# Get pods with wide output (shows node, IP, etc.)
kcsi get pods -n production -o wide

# Get services in yaml format
kcsi get services -n default -o yaml

# Get deployments as JSON
kcsi get deploy -n kube-system -o json

# Get nodes with extended information
kcsi get nodes -o wide
```

### Monitor cluster events

```bash
# Get recent events across all namespaces (sorted by timestamp)
kcsi events

# Get events in a specific namespace
kcsi events -n production

# Watch events in real-time
kcsi events -w
kcsi events -n kube-system -w
```

### Check for pod errors

```bash
# Find all pods with issues (CrashLoopBackOff, Error, Pending, etc.)
kcsi check errors
# or
kcsi check err

# Output example:
# Checking for pods with errors across all namespaces...
# (Excluding: Running, Completed)
#
# NAMESPACE     NAME                    READY   STATUS             RESTARTS   AGE
# production    api-server-xxx          0/1     CrashLoopBackOff   5          10m
# staging       worker-yyy              0/1     ImagePullBackOff   0          5m
#
# ⚠ Found pods with issues. Common states to investigate:
#   - CrashLoopBackOff: Pod is repeatedly crashing
#   - Error: Pod encountered an error
#   - Pending: Pod cannot be scheduled
#   - ImagePullBackOff: Cannot pull container image
#
# Use 'kcsi logs -n <namespace> <pod>' to investigate further
# Use 'kcsi describe pod -n <namespace> <pod>' for detailed information
```

### Show version and project information

```bash
# Show version and author
kcsi --version
# Output: kcsi version 0.5.0
#         Author: Alessandro

# Show detailed version information
kcsi --version-detailed
# Output: Kubectl Smart Interactive (kcsi)
#         Version: 0.5.0
#         Author: Alessandro Middei
#         Go Version: go1.25.5
#         OS/Arch: darwin/arm64

# Show project information and philosophy
kcsi about
# Displays project spirit, key principles, author info, and license
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

### Phase 3: Extended Resource Support ✅
- [x] `get services` command with autocompletion
- [x] `get deployments` command with autocompletion
- [x] `get nodes` command with autocompletion
- [x] `get namespaces` command
- [x] `get configmaps` command with autocompletion
- [x] `get secrets` command with autocompletion
- [x] `describe` commands for all resource types
- [x] Aliases support (svc, deploy, cm, ns, no)

### Phase 4: Delete Operations ✅
- [x] `delete pod` command with confirmation prompt
- [x] `delete service` command with confirmation
- [x] `delete deployment` command with confirmation
- [x] `delete configmap` command with confirmation
- [x] `delete secret` command with confirmation
- [x] `--force` flag to skip confirmation for automation
- [x] Safety prompts showing resource type, name, and namespace

### Phase 5: Diagnostics & Output Control ✅
- [x] `-o/--output` flag for all get commands (wide, yaml, json)
- [x] `events` command with namespace filtering and watch mode
- [x] `check errors` command to find problematic pods
- [x] Helpful troubleshooting suggestions

### Phase 6: Additional Commands (Next)
- [ ] `exec` command with interactive pod selection
- [ ] `port-forward` command
- [ ] `apply` and `edit` commands
- [ ] `rollout` commands (status, restart, undo)
- [ ] `top` commands for resource usage

### Phase 7: Enhancements
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
