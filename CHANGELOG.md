# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Multi-platform build script (`build.sh`) for creating binaries
- Support for darwin/amd64, darwin/arm64, linux/amd64, linux/arm64, and linux/arm
- Build instructions in README for multi-platform compilation

## [0.1.0] - 2025-12-24

### Added
- Initial project setup with Go modules
- Basic CLI structure using Cobra framework
- `kcsi get pods` command with `-n/--namespace` flag
- Namespace autocompletion support using kubectl
- Completion script generation for bash, zsh, fish, and powershell
- Project structure with separated concerns:
  - `cmd/` for Cobra commands
  - `pkg/kubernetes/` for kubectl wrapper functions
  - `pkg/completion/` for autocompletion logic
- README with installation and usage instructions
- CHANGELOG for tracking project history

### Technical Details
- Used `kubectl get namespaces -o jsonpath` for efficient namespace retrieval
- Implemented completion functions using Cobra's `RegisterFlagCompletionFunc`
- Direct kubectl command passthrough for `get pods` to maintain kubectl's output format

### Known Limitations
- No caching yet, each autocomplete queries the cluster
- Only supports `get pods` command in this POC
- Requires kubectl to be installed and configured

[Unreleased]: https://github.com/alessandro/kcsi/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/alessandro/kcsi/releases/tag/v0.1.0
