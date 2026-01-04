<div align="center">

<img src="logo.png" alt="KCSI Logo" width="200"/>

# KCSI

**kubectl for humans** – Interactive Kubernetes CLI with cascading autocomplete and guided workflows

*Your friendly Kapibara buddy for Day-2 Kubernetes operations*

[![Documentation](https://img.shields.io/badge/docs-stanzinofree.github.io-blue?style=for-the-badge)](https://stanzinofree.github.io/kcsi/)
[![Cheatsheet](https://img.shields.io/badge/cheatsheet-quick%20reference-green?style=for-the-badge)](https://stanzinofree.github.io/kcsi/cheatsheet/)
[![Roadmap](https://img.shields.io/badge/roadmap-what's%20next-purple?style=for-the-badge)](https://stanzinofree.github.io/kcsi/roadmap/)
[![Buy Me A Coffee](https://img.shields.io/badge/Buy%20Me%20A%20Coffee-support-yellow?style=for-the-badge&logo=buy-me-a-coffee)](https://buymeacoffee.com/smilzao)

[![Go Report Card](https://goreportcard.com/badge/github.com/stanzinofree/kcsi)](https://goreportcard.com/report/github.com/stanzinofree/kcsi)
[![License](https://img.shields.io/github/license/stanzinofree/kcsi)](LICENSE)
[![Release](https://img.shields.io/github/v/release/stanzinofree/kcsi)](https://github.com/stanzinofree/kcsi/releases)
[![Build and Test](https://github.com/stanzinofree/kcsi/workflows/Build%20and%20Test/badge.svg)](https://github.com/stanzinofree/kcsi/actions/workflows/build.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_kcsi&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=stanzinofree_kcsi)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_kcsi&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=stanzinofree_kcsi)

</div>

---

## Why KCSI?

You know Kubernetes. You just don't remember the exact command syntax every time.

**KCSI eliminates context switching.** Instead of opening browser tabs or typing `kubectl --help` for the 47th time, you get:

- **Cascading TAB autocomplete** – select namespace → resource type → pod → container in one flow
- **Guardrails on destructive actions** – confirmation prompts before delete/drain/rollout restart
- **Day-2 ops shortcuts** – check events, debug pods, inspect PVCs, rollout status – without memorizing flags
- **kubectl muscle memory** – no paradigm shift, no TUI, just faster kubectl
- **Single cross-platform binary** – install once, works everywhere (macOS, Linux, Windows)

Perfect for DevOps engineers and sysadmins who touch Kubernetes intermittently but need to move fast when they do.

---

## Quick Start (30 seconds)

```bash
# Install (macOS/Linux)
curl -sSL https://raw.githubusercontent.com/stanzinofree/kcsi/main/install.sh | bash

# Or with Homebrew
brew install stanzinofree/tap/kcsi

# Check installation
kcsi version

# Try it
kcsi logs
# → TAB to select namespace
# → TAB to select pod
# → TAB to select container
# → logs stream instantly
```

That's it. You're ready.

---

## The 10 commands you'll actually use

<details>
<summary><strong>Click to expand the essentials</strong></summary>

```bash
# 1. Stream logs (with TAB autocomplete for namespace/pod/container)
kcsi logs

# 2. Exec into a pod
kcsi attach

# 3. Describe a resource
kcsi describe

# 4. Check recent events in a namespace
kcsi events

# 5. Port-forward to a service or pod
kcsi port-forward

# 6. Get resource status (pods, deployments, services, etc.)
kcsi get

# 7. Delete a resource (with confirmation prompt)
kcsi delete

# 8. Check rollout status
kcsi rollout status

# 9. Restart a deployment
kcsi rollout restart

# 10. Debug a pod (ephemeral container)
kcsi debug
```

</details>

**Every command supports TAB completion.** Start typing, press TAB, select from the list. No flags to remember.

---

## Features

What makes KCSI different:

- **Cascading TAB selection** – namespace → resource → pod → container. One smooth flow, no typing resource names.
- **Guardrails for safety** – Destructive operations (delete, drain, rollout restart) require explicit confirmation. No accidental production disasters.
- **Day-2 ops workflows built-in** – `kcsi events` for recent cluster activity, `kcsi check errors` to surface failing pods, `kcsi get pvc pods` for storage troubleshooting, `kcsi rollout status` for deployments, `kcsi dig` for DNS debugging.
- **Works with your kubectl muscle memory** – Same verbs (`get`, `describe`, `logs`, `attach`), same mental model. No learning curve for a new TUI or abstraction layer.
- **Cross-platform single binary** – Written in Go. Drop the binary anywhere (macOS, Linux, Windows). No dependencies, no runtime, no containers.

---

## Safety & Security notes

- **Confirmation prompts** on destructive actions (delete, drain, rollout restart)
- **Read-only by default** for most commands (logs, describe, get, events)
- **Respects your kubeconfig** – uses the same context and credentials as `kubectl`
- **No telemetry** – KCSI does not phone home or collect usage data
- **Open source** – audit the code yourself
- **Security warnings** when displaying decoded secrets – see [docs/SECURITY_SECRETS.md](docs/SECURITY_SECRETS.md)

---

## Documentation

- **Full docs**: [https://stanzinofree.github.io/kcsi/](https://stanzinofree.github.io/kcsi/)
- **Cheatsheet**: [https://stanzinofree.github.io/kcsi/cheatsheet/](https://stanzinofree.github.io/kcsi/cheatsheet/)
- **Roadmap**: [https://stanzinofree.github.io/kcsi/roadmap/](https://stanzinofree.github.io/kcsi/roadmap/)

---

## Support KCSI ☕️

KCSI is free and open source. If it saves you time, consider supporting:

- **[Buy Me a Coffee](https://buymeacoffee.com/smilzao)** – one-time support
- **[GitHub Sponsors](https://github.com/sponsors/stanzinofree)** – recurring sponsorship

**Workshops and customization packs available for teams.** Reach out via GitHub or sponsors page.

---

## Installation options (advanced)

<details>
<summary><strong>macOS</strong></summary>

```bash
# Homebrew
brew install stanzinofree/tap/kcsi

# Or download binary from GitHub releases
curl -L https://github.com/stanzinofree/kcsi/releases/latest/download/kcsi-darwin-amd64 -o kcsi
chmod +x kcsi
sudo mv kcsi /usr/local/bin/
```

</details>

<details>
<summary><strong>Linux</strong></summary>

```bash
# Install script
curl -sSL https://raw.githubusercontent.com/stanzinofree/kcsi/main/install.sh | bash

# Or manual download
curl -L https://github.com/stanzinofree/kcsi/releases/latest/download/kcsi-linux-amd64 -o kcsi
chmod +x kcsi
sudo mv kcsi /usr/local/bin/
```

</details>

<details>
<summary><strong>Windows</strong></summary>

Download the latest `.exe` from [GitHub Releases](https://github.com/stanzinofree/kcsi/releases) and add to PATH.

</details>

<details>
<summary><strong>Build from source</strong></summary>

```bash
git clone https://github.com/stanzinofree/kcsi.git
cd kcsi
go build -o kcsi
sudo mv kcsi /usr/local/bin/
```

</details>

---

## Setup Autocompletion

<details>
<summary><strong>Bash</strong></summary>

```bash
# Load in current session
source <(kcsi completion bash)

# Load for all sessions (Linux)
kcsi completion bash > /etc/bash_completion.d/kcsi

# Load for all sessions (macOS with Homebrew)
kcsi completion bash > $(brew --prefix)/etc/bash_completion.d/kcsi
```

</details>

<details>
<summary><strong>Zsh</strong></summary>

```bash
# Enable completion if not already enabled
echo "autoload -U compinit; compinit" >> ~/.zshrc

# Generate completion script
kcsi completion zsh > "${fpath[1]}/_kcsi"

# Restart your shell
```

</details>

<details>
<summary><strong>Fish</strong></summary>

```bash
# Load in current session
kcsi completion fish | source

# Load for all sessions
kcsi completion fish > ~/.config/fish/completions/kcsi.fish
```

</details>

<details>
<summary><strong>PowerShell (Windows)</strong></summary>

```powershell
# Generate completion script
kcsi completion powershell | Out-String | Invoke-Expression

# Load for all sessions - add to your PowerShell profile
# Find your profile location with: $PROFILE
# Then add this line to the profile:
kcsi completion powershell | Out-String | Invoke-Expression
```

</details>

---

## Usage Examples

<details>
<summary><strong>Get pods with namespace autocompletion</strong></summary>

```bash
# Type this and press TAB after -n to see all available namespaces
kcsi get pods -n <TAB>

# Example
kcsi get pods -n kube-system
```

</details>

<details>
<summary><strong>Stream logs with cascading autocompletion</strong></summary>

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
```

</details>

<details>
<summary><strong>Delete resources safely</strong></summary>

```bash
# Delete pod with confirmation prompt
kcsi delete pod -n <TAB>  # Shows namespaces
kcsi delete pod -n default <TAB>  # Shows pods in namespace
kcsi delete pod -n default my-pod
# Output: Are you sure you want to delete pod 'my-pod' in namespace 'default'? [y/N]:

# Delete with --force to skip confirmation (use with caution!)
kcsi delete pod -n default my-pod --force
```

</details>

<details>
<summary><strong>Monitor cluster events</strong></summary>

```bash
# Get recent events across all namespaces (sorted by timestamp)
kcsi events

# Get events in a specific namespace
kcsi events -n production

# Watch events in real-time
kcsi events -w
```

</details>

<details>
<summary><strong>Check for pod errors</strong></summary>

```bash
# Find all pods with issues (CrashLoopBackOff, Error, Pending, etc.)
kcsi check errors

# Output includes helpful diagnostics suggestions
```

</details>

<details>
<summary><strong>Debug pods with ephemeral containers</strong></summary>

```bash
# Attach ephemeral debug container to pod
kcsi debug -n production my-pod

# Features:
# - Automatic internet connectivity check
# - Smart image selection (netshoot → alpine → busybox)
# - Full networking and debugging toolkit
```

</details>

<details>
<summary><strong>Port forwarding</strong></summary>

```bash
# Forward local port 8080 to pod port 80
kcsi port-forward -n default my-pod 8080:80

# Features:
# - Root privilege check for ports < 1024
# - Port availability check before forwarding
```

</details>

<details>
<summary><strong>View and decode secrets</strong></summary>

```bash
# View all keys and values of a secret (decoded from base64)
kcsi get secrets decoded my-secret -n production

# Show only a specific key from a secret
kcsi get secrets show my-secret -n production -k api-key

# ⚠️ Security Note: See docs/SECURITY_SECRETS.md for security considerations
```

</details>

<details>
<summary><strong>Rollout management</strong></summary>

```bash
# Restart a deployment to trigger a new rollout
kcsi rollout restart deployment my-app -n production

# Check rollout status
kcsi rollout status deployment my-app -n production

# View rollout history
kcsi rollout history deployment my-app -n production

# Rollback to previous revision
kcsi rollout undo deployment my-app -n production

# Rollback to specific revision
kcsi rollout undo deployment my-app -n production --to-revision=3
```

</details>

<details>
<summary><strong>Apply configurations</strong></summary>

```bash
# Apply from a single file
kcsi apply -f deployment.yaml -n production

# Apply from a directory recursively
kcsi apply -f ./k8s-manifests --recursive -n production

# Apply from kustomize directory
kcsi apply -k ./overlays/production

# Dry-run to preview changes
kcsi apply -f deployment.yaml -n production --dry-run
```

</details>

<details>
<summary><strong>Edit resources with automatic backup</strong></summary>

```bash
# Edit a deployment with automatic backup
kcsi edit deployment my-app -n production

# Features:
# - Automatic backup to ~/.kcsi/backups/
# - Custom backup directory: --backup-dir
# - Skip backup: --no-backup
# - Custom editor: --editor or KUBE_EDITOR env var
```

</details>

---

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
task pr:push

# After PR is merged, create a tag and release
task tag VERSION=0.6.0 MESSAGE="Release notes"
```

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

For detailed development information, see the full documentation at [https://stanzinofree.github.io/kcsi/](https://stanzinofree.github.io/kcsi/).

---

## License

MIT License

---

<div align="center">

**Made with ❤️ by sysadmins, for sysadmins**

</div>
