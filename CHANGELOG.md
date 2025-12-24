# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2025-12-24

### Added - Phase 2: Expanded Commands
- `kcsi logs` command with full kubectl logs functionality
  - `-f/--follow` flag for following log output
  - `--tail` flag for limiting output lines
  - `-p/--previous` flag for previous container logs
  - `-c/--container` flag for multi-container pods
- `kcsi describe pod` command for describing pods
- Cascading autocompletion: namespace → pod → container
- Container autocompletion for multi-container pods
- `GetContainers()` function in kubernetes package
- `ContainerCompletion()` function in completion package
- Comprehensive usage examples in README for all new commands

### Added - Build and Development Tools
- Multi-platform build script (`build.sh`) for creating binaries
- Support for darwin/amd64, darwin/arm64, linux/amd64, linux/arm64, and linux/arm
- Build instructions in README for multi-platform compilation
- Taskfile.yml for simplified development workflow
- Task commands for build, clean, test, install, uninstall, and more
- Comprehensive documentation for Task usage in README
- Task command reference table in README

### Changed
- Updated README with Phase 2 completion status
- Enhanced roadmap with Phase 3 and Phase 4 planning
- Version bumped to 0.2.0

### Technical Details
- Used jsonpath queries for efficient container name retrieval
- Implemented ValidArgsFunction for positional argument autocompletion
- All kubectl flags are properly passed through to maintain full compatibility

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

[Unreleased]: https://github.com/alessandro/kcsi/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/alessandro/kcsi/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/alessandro/kcsi/releases/tag/v0.1.0
