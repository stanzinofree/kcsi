# kcsi - Kubectl Smart Interactive Wrapper

[![Version](https://img.shields.io/badge/version-0.5.3-blue.svg)](https://github.com/stanzinofree/kcsi/releases)
[![Go Version](https://img.shields.io/badge/go-1.23+-00ADD8.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey.svg)](#installation)
[![Documentation](https://img.shields.io/badge/docs-GitHub%20Pages-blue.svg)](https://stanzinofree.github.io/kcsi/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

### CI/CD & Quality

[![Build and Test](https://github.com/stanzinofree/kcsi/workflows/Build%20and%20Test/badge.svg)](https://github.com/stanzinofree/kcsi/actions/workflows/build.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_kcsi&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=stanzinofree_kcsi)

### Security & Code Quality
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_kcsi&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=stanzinofree_kcsi)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_kcsi&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=stanzinofree_kcsi)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_kcsi&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=stanzinofree_kcsi)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_kcsi&metric=bugs)](https://sonarcloud.io/summary/new_code?id=stanzinofree_kcsi)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_kcsi&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=stanzinofree_kcsi)
[![Duplicated Lines](https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_kcsi&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=stanzinofree_kcsi)

### Dependency Management

[![Dependabot Status](https://img.shields.io/badge/Dependabot-enabled-brightgreen.svg?logo=dependabot)](https://github.com/stanzinofree/kcsi/network/updates)

<p align="center">
  <i>A kubectl wrapper with intelligent autocompletion for namespaces, pods, and other Kubernetes resources.</i>
</p>

<p align="center">
  <a href="#features">Features</a> •
  <a href="#installation">Installation</a> •
  <a href="#quick-start">Quick Start</a> •
  <a href="#usage">Usage</a> •
  <a href="https://stanzinofree.github.io/kcsi/cheatsheet.html">📖 Cheatsheet</a> •
  <a href="https://stanzinofree.github.io/kcsi/roadmap.html">🗺️ Roadmap</a>
</p>

---

## Features

- Smart autocompletion for Kubernetes namespaces
- Interactive resource selection
- Simplified kubectl commands with tab completion
- Centralized version and project information
- Built with Go and Cobra

## Quick Start

```bash
# Install kcsi
curl -sL https://raw.githubusercontent.com/stanzinofree/kcsi/main/install.sh | bash

# Set up autocompletion (bash example)
source <(kcsi completion bash)

# Start using kcsi!
kcsi get pods -n <TAB>  # Press TAB to see all namespaces
```

## Cheatsheet

### Quick Reference

For a complete interactive cheatsheet with search functionality, see **[📖 Full Cheatsheet](https://stanzinofree.github.io/kcsi/cheatsheet.html)**.

**Most Common Commands:**

```bash
# List pods with namespace autocomplete
kcsi get pods -n <TAB>

# Follow pod logs with autocomplete
kcsi logs -f -n <ns> <pod>

# Describe resources with cascading autocomplete
kcsi describe pod -n <ns> <pod>

# Delete resources with confirmation prompt
kcsi delete pod -n <ns> <pod>

# Find all problematic pods
kcsi check errors

# Watch cluster events in real-time
kcsi events -w
```

**Resource Aliases:** `svc`, `deploy`, `ns`, `no`, `cm`

**Useful Flags:** `-n` (namespace), `-o` (output format), `-f` (follow/force), `--tail`, `-w` (watch)

> 💡 **Tip:** Use the [interactive cheatsheet](https://stanzinofree.github.io/kcsi/cheatsheet.html) to quickly search all available commands and options.

## Current Status

**Version:** 0.5.3 - Resource Monitoring & DNS Debugging

Currently implemented:

**Get Commands (with -o/--output flag):**
- `kcsi get pods -o wide` - List pods with node information
- `kcsi get services -o wide` (alias: `svc`) - List services with extended info
- `kcsi get deployments -o wide` (alias: `deploy`) - List deployments
- `kcsi get namespaces` (alias: `ns`) - List namespaces
- `kcsi get nodes -o wide` (alias: `no`) - List nodes with details
- `kcsi get configmaps` (alias: `cm`) - List configmaps
- `kcsi get secrets` - List secrets
- `kcsi get internal-domains` (aliases: `idomains`, `idom`) - List all internal Kubernetes FQDNs
  - Shows services: `service.namespace.svc.cluster.local`
  - Shows pods: `pod-ip.namespace.pod.cluster.local`
  - Displays resource type, name, namespace, FQDN, IP, and additional info
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

**Interactive & Execution Commands:**
- `kcsi attach` - Attach to a running pod with automatic shell detection
- `kcsi execute` - Execute custom commands in pods
- `kcsi port-forward` - Forward local ports to pods with validation
  - Root privilege check for ports < 1024
  - Port availability check before forwarding

**Resource Usage & Debugging:**
- `kcsi top pods` - Display CPU and memory usage for pods
- `kcsi top nodes` - Display CPU and memory usage for nodes
- `kcsi dig [namespace] [pod] [domain]` - DNS debugging inside pods
  - Namespace-first autocompletion
  - Container selection with `-c` flag

**Other Commands:**
- `kcsi logs` - Get pod logs with full kubectl flags support (-f, --tail, -p, -c)
- Container autocompletion for multi-container pods
- Cascading autocompletion: namespace → resource → container
- `kcsi about` - Show project information, philosophy, and author details
- `kcsi --version` - Show version and author
- `kcsi --version-detailed` - Show detailed version info (Go version, OS/Arch)

## Installation

### Prerequisites

- Go 1.23 or higher
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
- Windows (x86_64): `kcsi-windows-amd64.exe`
- Windows (ARM64): `kcsi-windows-arm64.exe`

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

#### macOS / Linux

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

#### Windows

**PowerShell (Administrator):**
```powershell
# Using pre-built binary (example for Windows x86_64)
Copy-Item bin\kcsi-windows-amd64.exe C:\Windows\System32\kcsi.exe

# Or add to a custom directory and update PATH
New-Item -ItemType Directory -Force -Path C:\Tools
Copy-Item bin\kcsi-windows-amd64.exe C:\Tools\kcsi.exe
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";C:\Tools", "Machine")
```

**Command Prompt (Administrator):**
```cmd
# Using pre-built binary
copy bin\kcsi-windows-amd64.exe C:\Windows\System32\kcsi.exe
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

### PowerShell (Windows)

```powershell
# Generate completion script
kcsi completion powershell | Out-String | Invoke-Expression

# Load for all sessions - add to your PowerShell profile
# Find your profile location with: $PROFILE
# Then add this line to the profile:
kcsi completion powershell | Out-String | Invoke-Expression
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

### Port forwarding

```bash
# Forward local port 8080 to pod port 80
kcsi port-forward -n default my-pod 8080:80

# Forward local port 3000 to pod port 8080
kcsi port-forward -n production web-app 3000:8080

# Forward privileged port (requires sudo for ports < 1024)
sudo kcsi port-forward -n production nginx-pod 80:8080

# Features:
# - Validates port numbers (1-65535)
# - Checks if running as root for privileged ports (< 1024)
# - Checks if local port is already in use
# - Interactive session (Ctrl+C to stop)
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

For a detailed roadmap with progress tracking and visual indicators, see **[🗺️ Full Roadmap](https://stanzinofree.github.io/kcsi/roadmap.html)**.

**Current Status:** 5 phases completed (38 features delivered), 2 phases planned (9 features upcoming)

**Recently Completed:**
- ✅ Internal domains listing with `get internal-domains` (shows all Kubernetes FQDNs)
- ✅ Port-forward with root privilege check and port availability validation
- ✅ Resource usage monitoring with `top` command (pods and nodes)
- ✅ DNS debugging with `dig` command inside pods
- ✅ Phase 5: Diagnostics & Output Control (events, error checking, output formats)

**Next Up:**
- 🔄 Phase 6: Additional Commands (exec, port-forward, apply, edit, rollout, top)
- 📋 Phase 7: Enhancements (caching, configuration, fuzzy matching)

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
| **Build Tasks** | |
| `task build` | Build for current platform |
| `task build:all` | Build for all platforms |
| `task build:darwin` | Build for macOS (Intel + ARM) |
| `task build:linux` | Build for Linux (amd64, arm64, arm) |
| `task clean` | Clean build artifacts |
| **Development Tasks** | |
| `task install` | Build and install to /usr/local/bin |
| `task uninstall` | Uninstall from /usr/local/bin |
| `task run -- <args>` | Build and run with arguments |
| `task test` | Run tests |
| `task fmt` | Format code |
| `task vet` | Run go vet |
| `task check` | Run fmt, vet, and test |
| `task dev` | Development mode (build + install) |
| **PR Workflow Tasks** | |
| `task pr BRANCH=fix/name TITLE='...' DESC='...'` | Create and push new PR |
| `task pr:push` | Push updates to current PR (interactive) |
| `task pr:update MESSAGE='...'` | Update PR with commit message |
| `task branch NAME=fix/name` | Create new branch from main |
| `task sync` | Sync main branch with remote |
| **Release Tasks** | |
| `task tag VERSION=0.5.4 MESSAGE='...'` | Create and push git tag |
| `task release` | Prepare release build (all platforms) |
| `task release:github VERSION=0.5.4` | Create GitHub release with binaries |
| **Other Tasks** | |
| `task completion:all` | Generate all completion scripts |

## Contributing

Contributions are welcome! Here's how to propose changes via Pull Requests:

### Making a Pull Request

1. **Create a branch** from updated main:
   ```bash
   git checkout main
   git pull origin main
   git checkout -b fix/your-feature-name
   ```

2. **Make your changes** and commit:
   ```bash
   git add .
   git commit -m "fix: description of your changes"
   ```

3. **Push your branch**:
   ```bash
   git push origin fix/your-feature-name
   ```

4. **Create the Pull Request**:
   ```bash
   gh pr create --title "Your PR Title" --body "Detailed description" --base main
   ```

### Using Task for PR Workflow

We provide convenient Task commands for the PR workflow:

```bash
# Create and push a PR
task pr BRANCH=fix/my-feature TITLE="Fix something" DESC="Detailed description"

# Push updates to existing PR
task pr-push

# After PR is merged, create a tag and release
task tag VERSION=0.5.4 MESSAGE="Release notes"
task release VERSION=0.5.4
```

See [Available Task Commands](#available-task-commands) for the full list.

### Commit Message Convention

- `feat:` New features
- `fix:` Bug fixes
- `docs:` Documentation changes
- `chore:` Maintenance tasks
- `test:` Test changes

### Repository Rules

- All PRs require passing CI/CD checks (tests, CodeQL, SonarCloud)
- Direct pushes to `main` are blocked
- PRs must be reviewed and merged via GitHub interface

## License

MIT License (to be added)

## Acknowledgments

Built with:
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [kubectl](https://kubernetes.io/docs/reference/kubectl/) - Kubernetes command-line tool
