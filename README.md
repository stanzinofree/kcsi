# kcsi - Kubectl Smart Interactive Wrapper

A kubectl wrapper with intelligent autocompletion for namespaces, pods, and other Kubernetes resources.

## Features

- Smart autocompletion for Kubernetes namespaces
- Interactive resource selection
- Simplified kubectl commands with tab completion
- Built with Go and Cobra

## Current Status

**Version:** 0.1.0 - Proof of Concept

Currently implemented:
- `kcsi get pods -n <namespace>` with namespace autocompletion

## Installation

### Prerequisites

- Go 1.21 or higher
- kubectl installed and configured
- Access to a Kubernetes cluster

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
```bash
# Build for all supported platforms
./build.sh

# Binaries will be created in the bin/ directory
```

### Install

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

## Roadmap

### Phase 1: Proof of Concept (Current)
- [x] Basic CLI structure with Cobra
- [x] `get pods` command with namespace flag
- [x] Namespace autocompletion
- [x] Completion script generation

### Phase 2: Expand Commands
- [ ] `logs` command with pod and namespace autocompletion
- [ ] `exec` command with interactive pod selection
- [ ] `describe` command with resource autocompletion
- [ ] Cascading autocompletion (namespace → pod)

### Phase 3: Enhancements
- [ ] Cache for faster autocompletion
- [ ] Default context/namespace configuration
- [ ] Custom aliases
- [ ] Fuzzy matching for resources

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
└── main.go          # Entry point
```

### Building

```bash
# Quick build for current platform
go build -o kcsi

# Multi-platform build
./build.sh
```

### Testing

```bash
# Run
./kcsi get pods -n <namespace>

# Test autocompletion (after setting up completion scripts)
./kcsi get pods -n <TAB>

# Test specific platform binary
./bin/kcsi-linux-arm64 --version
```

## Contributing

This is currently a personal project in early development. Feedback and suggestions are welcome!

## License

MIT License (to be added)

## Acknowledgments

Built with:
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [kubectl](https://kubernetes.io/docs/reference/kubectl/) - Kubernetes command-line tool
