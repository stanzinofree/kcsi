# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.7.0] - 2026-01-08

### Added - Context Management for Multi-Cluster Operations

#### New Command
- **Context Management** - Manage multiple Kubernetes clusters without modifying system kubeconfig:
  - `kcsi context import <name> <kubeconfig-path>` - Import and copy kubeconfig to managed directory
  - `kcsi context add <name> <kubeconfig-path>` - Add reference to existing kubeconfig
  - `kcsi context list` (alias: `ls`) - List all configured contexts with current marker
  - `kcsi context use <name>` - Switch to a different cluster context
  - `kcsi context current` - Display currently active context
  - `kcsi context remove <name>` (aliases: `rm`, `delete`) - Remove a context and its files
  - `--description` / `-d` flag for adding context descriptions
  - Full autocompletion for context names

#### Infrastructure
- New `pkg/context` package for context management:
  - `Context` struct for representing cluster configurations
  - `Config` struct for managing contexts.yaml
  - Configuration stored in `~/.kcsi/contexts.yaml`
  - Imported kubeconfigs stored in `~/.kcsi/contexts/<name>/kube.config`
  - `InitializeKcsiDir()` creates directory structure on first use
  - `LoadConfig()` / `SaveConfig()` for configuration persistence
  - `AddContext()` / `ImportContext()` for adding clusters
  - `RemoveContext()` for cleanup with file deletion
  - `SetCurrentContext()` / `GetCurrentContext()` for active context tracking
  - `GetContextKubeconfigPath()` for path resolution
- Enhanced `pkg/kubernetes/client.go`:
  - `setKubeconfigEnv()` helper sets KUBECONFIG environment variable
  - All kubectl command executions now respect active kcsi context
  - Integrated with `ExecuteKubectl()`, `ExecuteKubectlInteractive()`
  - Applied to all CommandContext calls (version, cluster-info, config commands)
- Context isolation from system kubeconfig:
  - System `~/.kube/config` is never modified
  - Each kcsi context uses its own KUBECONFIG environment variable
  - Contexts are completely isolated from kubectl's default behavior

### Changed
- All kubectl operations now use active kcsi context when configured
- Version bumped to 0.7.0
- Updated README with Context Management section in Advanced features

### Technical Details
- Context configurations use YAML format with gopkg.in/yaml.v3
- Import operation copies kubeconfig files with 0600 permissions for security
- Add operation creates references without copying files
- Remove operation deletes both configuration and imported files
- Current context tracked in contexts.yaml for persistence across sessions
- All commands automatically detect and use active context via environment variables

### Migration Notes
- Existing kcsi installations continue to use system kubeconfig until contexts are configured
- First context import/add creates `~/.kcsi/` directory structure automatically
- No breaking changes to existing command syntax or behavior

## [0.6.0] - 2025-12-27

### Added - Phase 6: Advanced Operations & Roadmap Reorganization

#### New Commands
- **Rollout Management** - Complete rollout lifecycle management:
  - `kcsi rollout restart <type> <name>` - Trigger a rollout restart for deployments, daemonsets, statefulsets
  - `kcsi rollout status <type> <name>` - Check rollout status
  - `kcsi rollout history <type> <name>` - View rollout revision history
  - `kcsi rollout undo <type> <name>` - Rollback to previous revision
  - `--to-revision` flag for undo command to rollback to specific revision
  - Full autocompletion for resource types and names

- **Apply Configuration** - Apply Kubernetes manifests with validation:
  - `kcsi apply -f <file>` - Apply from file
  - `kcsi apply -f <dir> --recursive` - Apply from directory recursively
  - `kcsi apply -k <dir>` - Apply from kustomize directory
  - `--dry-run` and `--server-dry-run` flags for testing
  - `--validate` flag for pre-apply validation
  - `--force` flag for force apply
  - File extension validation with helpful warnings
  - Directory detection with recursive flag requirement

- **Edit with Automatic Backup** - Safe resource editing:
  - `kcsi edit <type> <name>` - Edit resource with automatic backup
  - Automatic backup to `~/.kcsi/backups/` with timestamp
  - `--backup-dir` flag for custom backup location
  - `--no-backup` flag to skip backup
  - `--editor` flag for custom editor (also supports `KUBE_EDITOR` env var)
  - Restore instructions displayed on edit failure
  - Backup filename includes resource type, name, namespace, and timestamp

#### Infrastructure
- Added `GetDaemonSets()` function to kubernetes client
- Added `GetStatefulSets()` function to kubernetes client
- Created `ErrNamespaceRequired` constant for consistent error messaging

#### Documentation
- **Roadmap Reorganization** - Restructured from phase-based to feature-based:
  - ✅ Core Features (9 items) - Essential kubectl operations
  - 🔧 Interactive & Debugging (6 items) - Troubleshooting tools  
  - 🚀 Advanced Features (7 items) - Power user capabilities
  - ⚙️ Quality & Developer Experience (7 items) - Code quality & UX
  - 📋 Planned Features organized by category:
    - ⚡ Performance (caching, parallel queries, fuzzy matching)
    - 🛠️ Utilities (context switching, backups, namespace defaults)
    - 🎨 Formatting & Visualization (pretty print, graphs, tables)
    - 🌐 Cluster Operations (cluster dump, health check, custom aliases)

### Changed
- Version bumped to 0.6.0
- Updated roadmap.html with new feature-based organization
- Enhanced CSS with improved section title contrast

### Fixed - Code Quality & Testing

#### SonarCloud Issues Resolved
- **Cognitive Complexity** in `cmd/apply.go`:
  - Reduced from 29 to ~8 (maximum allowed: 15)
  - Extracted 7 helper functions for better maintainability:
    - `addSourceArgs()` - handles kustomize vs filename logic
    - `addKustomizeArgs()` - processes kustomize directories
    - `addFilenameArgs()` - processes file/directory paths
    - `validateFilePath()` - validates file existence
    - `addDirectoryArgs()` - handles directory with recursive flag
    - `addFileArgs()` - handles single file with extension validation
    - `addApplyFlags()` - adds all optional flags

- **String Duplication** in `cmd/rollout.go`:
  - Created `ErrNamespaceRequired` constant in `cmd/constants.go`
  - Replaced 4 duplicate occurrences of "namespace is required (use -n flag)"

#### Testing Improvements
- **Dynamic Version Tests** - Tests no longer hardcoded to specific version:
  - Removed hardcoded version checks (previously "0.5.3")
  - Tests now validate semantic versioning pattern (contains dots)
  - Future-proof: tests pass regardless of current version number
  - Better maintainability for version bumps

#### Documentation Accessibility
- Fixed CSS contrast issue in roadmap planned features section
- Section titles now use solid color (#212529) instead of gradient with transparent fill
- Improved WCAG AA compliance

### Technical Details
- Rollout commands support deployment, daemonset, and statefulset resource types
- Apply command validates file paths and extensions before execution
- Edit command creates timestamped backups in format: `<type>-<name>-<namespace>-<timestamp>.<format>`
- All new commands maintain consistent flag patterns and autocompletion
- Code refactoring maintains 100% backward compatibility

## [0.5.3] - 2025-12-25

### Added - Interactive Pod Commands & Container Support
- `kcsi attach` command for interactive shell attachment to pods:
  - Automatic shell detection (tries bash → zsh → sh in order)
  - `-n/--namespace` flag with namespace autocompletion
  - `-c/--container` flag with container autocompletion for multi-container pods
  - Namespace-first syntax for proper shell autocompletion
  - Interactive TTY support for full shell experience
- `kcsi execute` command for running custom commands in pods:
  - Execute arbitrary commands in pod containers
  - `-n/--namespace` flag with namespace autocompletion
  - `-c/--container` flag with container autocompletion
  - Namespace-first syntax for better autocompletion workflow
  - Helpful error messages when namespace is missing
  - Full command passthrough to kubectl exec
- Container flag support extended to existing commands:
  - `kcsi logs -c <container>` for viewing specific container logs
  - `kcsi describe pod -c <container>` for container-specific information
  - Container autocompletion works across all container-aware commands

### Changed
- Improved command syntax to namespace-first pattern for better autocompletion UX
- Enhanced error messages with helpful hints (e.g., "try passing -n namespace")
- Version bumped to 0.5.3

### Fixed - Code Quality & Accessibility
- Fixed SonarCloud string literal duplication issues:
  - Created `cmd/constants.go` for shared flag descriptions
  - Created constants in `pkg/kubernetes/client.go` for JSONPath queries
  - Eliminated duplicate strings across codebase
- Fixed WCAG AA text contrast issues in documentation:
  - Updated text colors from #333 → #212529 and #666 → #495057
  - Improved footer design with modern gradient and better contrast
  - Enhanced accessibility compliance
- Fixed test function naming to comply with SonarCloud regex:
  - Renamed test functions to remove underscores
  - Now follows Go naming convention: `^[A-Z][a-zA-Z0-9]+$`

### Technical Details
- Attach command attempts multiple shells gracefully, falling back if one fails
- Execute command properly passes command arrays to kubectl exec
- Container autocompletion queries pod spec for container names
- Namespace-first argument ordering enables proper cascading autocompletion
- All new commands maintain consistent flag patterns with existing codebase

## [0.5.2] - 2025-12-25

### Added - Documentation Improvements
- Enhanced README with professional badges (version, Go version, license, platform, PRs welcome)
- Added Quick Start section with installation and first steps
- Added comprehensive Cheatsheet section:
  - Most common commands table
  - Resource aliases reference
  - Useful flags reference
  - Quick diagnostics examples
- Created beautiful HTML mini-guide (`docs/guide.html`):
  - Modern gradient design with responsive layout
  - Interactive feature cards with hover effects
  - Step-by-step installation guide
  - Comprehensive command reference tables
  - Practical examples section
  - Advanced features documentation
  - Mobile-friendly responsive design
- Improved README navigation with anchor links
- Added centered introduction and navigation menu

### Changed
- README now has a more professional and attractive appearance
- Better organization of content with visual hierarchy
- Version bumped to 0.5.2

## [0.5.1] - 2025-12-25

### Added - Version and Project Information System
- Centralized version manifest (`pkg/version/version.yaml`):
  - Single source of truth for version, author, and project metadata
  - Contains project spirit and philosophy
  - Includes license and repository information
- `pkg/version` package for reading and presenting version information:
  - `GetVersion()` returns simple version string
  - `GetDetailedVersion()` shows version, author, build info, Go version, OS/Arch
  - `GetAbout()` returns formatted project information with spirit/philosophy
  - Uses go:embed to include manifest in binary
- `kcsi about` command:
  - Displays project name, version, and description
  - Shows project spirit and key principles
  - Includes author, license, and repository information
  - Shows build details (Go version, OS/Arch)
  - Beautiful formatted output with borders
- Enhanced version flags:
  - `--version` / `-v` shows version and author name
  - `--version-detailed` shows comprehensive version information
  - Custom version template in root command
- Updated README with version/about command examples
- Added feature to README: "Centralized version and project information"

### Changed
- Root command now uses version package instead of hardcoded version
- Version bumped to 0.5.1

### Technical Details
- Manifest file uses YAML format for easy editing
- go:embed ensures manifest is included in compiled binary
- Version package provides clean API for version information
- --version-detailed flag handled in Execute() before Cobra processing

## [0.5.0] - 2025-12-24

### Added - Phase 5: Diagnostics & Output Control
- `kcsi events` command for cluster event monitoring:
  - `-n/--namespace` flag for filtering events by namespace
  - `-w/--watch` flag for real-time event streaming
  - Events automatically sorted by timestamp for readability
  - All namespaces shown by default when namespace not specified
- `kcsi check errors` command for finding problematic pods:
  - Scans all namespaces for pods not in Running or Completed state
  - Shows pods with issues: CrashLoopBackOff, Error, Pending, ImagePullBackOff, etc.
  - Displays helpful troubleshooting suggestions
  - Provides next-step commands for investigation (logs, describe)
  - Clean output with success message when no issues found
- `-o/--output` flag support for all `get` commands:
  - Added to: get pods, services, deployments, nodes, configmaps, secrets
  - Supports all kubectl output formats: wide, yaml, json, etc.
  - Enables viewing node placement with `-o wide`
  - Full kubectl output format compatibility

### Changed
- Updated README with Phase 5 completion status
- Added comprehensive examples for events, check errors, and output formats
- Enhanced roadmap with Phase 5 complete and reorganized future phases
- Version bumped to 0.5.0

### Technical Details
- Events command passes through to kubectl with proper flag mapping
- Check errors uses kubectl output parsing to filter pod states
- Output flag cleanly integrates with existing get command structure
- All new features maintain consistent autocompletion patterns

## [0.4.0] - 2025-12-24

### Added - Phase 4: Delete Operations with Safety
- `kcsi delete` commands for all major resource types:
  - `delete pod` with confirmation prompt and namespace autocompletion
  - `delete service` (aliases: svc, services) with confirmation
  - `delete deployment` (aliases: deploy, deployments) with confirmation
  - `delete configmap` (aliases: cm, configmaps) with confirmation
  - `delete secret` (alias: secrets) with confirmation
- Safety confirmation system:
  - Interactive prompt before deletion showing resource type, name, and namespace
  - Requires explicit 'y' or 'yes' response to proceed
  - `--force` / `-f` flag to skip confirmation for automation/scripts
  - Graceful cancellation with clear feedback
- `askForConfirmation()` helper function for user interaction
- Generic `runKubectlDelete()` function with confirmation handling
- All delete commands support cascading autocompletion (namespace → resource)

### Changed
- Updated README with delete command examples and safety feature documentation
- Enhanced roadmap with completed Phase 4 and reorganized future phases
- Version bumped to 0.4.0

### Technical Details
- Confirmation prompt uses bufio.Reader for stdin interaction
- Delete operations respect namespace scoping
- Consistent error handling across all delete commands
- Force flag properly bypasses confirmation while maintaining safety

## [0.3.0] - 2025-12-24

### Added - Phase 3: Extended Resource Support
- `kcsi get` commands for all major resource types:
  - `get services` (aliases: svc, service) with namespace autocompletion
  - `get deployments` (aliases: deploy, deployment) with namespace autocompletion
  - `get nodes` (aliases: no, node) with cluster-wide autocompletion
  - `get namespaces` (aliases: ns, namespace) for listing all namespaces
  - `get configmaps` (aliases: cm, configmap) with namespace autocompletion
  - `get secrets` (alias: secret) with namespace autocompletion
- `kcsi describe` commands for all resource types:
  - `describe service` with cascading autocompletion
  - `describe deployment` with cascading autocompletion
  - `describe node` with node name autocompletion
  - `describe configmap` with cascading autocompletion
  - `describe secret` with cascading autocompletion
- New kubernetes client functions:
  - `GetServices()` for retrieving service names
  - `GetDeployments()` for retrieving deployment names
  - `GetNodes()` for retrieving node names
  - `GetConfigMaps()` for retrieving configmap names
  - `GetSecrets()` for retrieving secret names
- New completion functions:
  - `ServiceCompletion()` for service name autocompletion
  - `DeploymentCompletion()` for deployment name autocompletion
  - `NodeCompletion()` for node name autocompletion
  - `ConfigMapCompletion()` for configmap name autocompletion
  - `SecretCompletion()` for secret name autocompletion
- Comprehensive usage examples in README for all new resource types

### Changed
- Refactored `cmd/get.go` for better modularity with generic `runKubectlGet()` function
- Refactored `cmd/describe.go` for better modularity with generic `runKubectlDescribe()` function
- Updated README with Phase 3 completion status
- Enhanced roadmap with completed Phase 3 and new Phase 4/5 planning
- Version bumped to 0.3.0

### Technical Details
- All resource retrieval uses consistent jsonpath queries for efficiency
- Aliases implemented using Cobra's Aliases feature
- Generic command runners reduce code duplication
- Consistent namespace flag handling across all commands

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

[Unreleased]: https://github.com/stanzinofree/kcsi/compare/v0.5.0...HEAD
[0.5.0]: https://github.com/stanzinofree/kcsi/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/stanzinofree/kcsi/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/stanzinofree/kcsi/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/stanzinofree/kcsi/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/stanzinofree/kcsi/releases/tag/v0.1.0
