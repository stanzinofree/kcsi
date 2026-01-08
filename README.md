<div align="center">

<img src="logo.png" alt="KCSI Logo" width="140" />

# KCSI

**kubectl for humans** – Cascading TAB + guardrails for day-2 ops.

*Your friendly Kapibara buddy making Kubernetes feel less scary*

[![Documentation](https://img.shields.io/badge/docs-read%20here-blue?style=for-the-badge)](https://stanzinofree.github.io/kcsi/)
[![Cheatsheet](https://img.shields.io/badge/cheatsheet-quick%20ref-green?style=for-the-badge)](https://stanzinofree.github.io/kcsi/cheatsheet.html)
[![Roadmap](https://img.shields.io/badge/roadmap-what's%20next-purple?style=for-the-badge)](https://stanzinofree.github.io/kcsi/roadmap.html)
[![Buy Me A Coffee](https://img.shields.io/badge/☕_Buy_Me_A_Coffee-support-yellow?style=for-the-badge)](https://buymeacoffee.com/smilzao)

[![Go Report Card](https://goreportcard.com/badge/github.com/stanzinofree/kcsi)](https://goreportcard.com/report/github.com/stanzinofree/kcsi)
[![License](https://img.shields.io/github/license/stanzinofree/kcsi)](LICENSE)
[![Release](https://img.shields.io/github/v/release/stanzinofree/kcsi)](https://github.com/stanzinofree/kcsi/releases)
[![Build and Test](https://github.com/stanzinofree/kcsi/actions/workflows/build.yml/badge.svg)](https://github.com/stanzinofree/kcsi/actions/workflows/build.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_kcsi&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=stanzinofree_kcsi)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_kcsi&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=stanzinofree_kcsi)

</div>

---

## Why KCSI?

You know Kubernetes. You just don't remember the exact command syntax every time.

**KCSI eliminates the context switching.** Instead of opening browser tabs or typing `kubectl --help` for the 47th time, you get:

- **Cascading TAB autocomplete** – namespace → resource type → pod → container in one smooth flow
- **Guardrails on destructive ops** – confirmation prompts before delete/drain/rollout restart
- **Day-2 workflows built-in** – check events, debug pods, inspect PVCs, rollout status – no flags to memorize
- **Your kubectl muscle memory works** – same verbs, same mental model, just faster
- **Single cross-platform binary** – drop it anywhere (macOS, Linux, Windows) and go

Perfect for sysadmins, DevOps engineers, and anyone who touches Kubernetes intermittently but needs to move fast when they do.

---

## Demo (12s)

See KCSI in action: install → TAB autocomplete → get pods

[![asciicast](https://asciinema.org/a/MDTbe6ahLXfv80YleTkXyPG4o.svg)](https://asciinema.org/a/MDTbe6ahLXfv80YleTkXyPG4o)

---

## Installation

### Quick install (macOS/Linux)

```bash
curl -sSL https://raw.githubusercontent.com/stanzinofree/kcsi/main/install.sh | bash
```

Verify:
```bash
kcsi version
```

### Platform binaries

**macOS Intel**
```bash
curl -L https://github.com/stanzinofree/kcsi/releases/latest/download/kcsi-darwin-amd64 -o kcsi
chmod +x kcsi && sudo mv kcsi /usr/local/bin/
```

**macOS Apple Silicon**
```bash
curl -L https://github.com/stanzinofree/kcsi/releases/latest/download/kcsi-darwin-arm64 -o kcsi
chmod +x kcsi && sudo mv kcsi /usr/local/bin/
```

**Linux** (x64)
```bash
curl -L https://github.com/stanzinofree/kcsi/releases/latest/download/kcsi-linux-amd64 -o kcsi
chmod +x kcsi && sudo mv kcsi /usr/local/bin/
```

**Windows**

Download the `.exe` from [GitHub Releases](https://github.com/stanzinofree/kcsi/releases) and add to PATH.

### Build from source

```bash
git clone https://github.com/stanzinofree/kcsi.git
cd kcsi
go build -o kcsi
sudo mv kcsi /usr/local/bin/
```

### Enable completion

**Bash**

Try once (current shell):
```bash
source <(kcsi completion bash)
```

Persist:
```bash
echo 'source <(kcsi completion bash)' >> ~/.bashrc
```

**Zsh**

Try once (current shell):
```bash
source <(kcsi completion zsh)
```

Persist:
```bash
echo 'source <(kcsi completion zsh)' >> ~/.zshrc
```

**Fish**
```fish
kcsi completion fish > ~/.config/fish/completions/kcsi.fish
```

**PowerShell (Windows)**
```powershell
# Add to your PowerShell profile ($PROFILE)
kcsi completion powershell | Out-String | Invoke-Expression
```

---

## Quick Start (30 seconds)

```bash
# Verify installation
kcsi version

# Stream logs – press TAB to select namespace/pod/container
kcsi logs

# No flags. No typing resource names. Just flow.
```

That's it. You're ready.

---

## The 10 commands you'll actually use

```bash
kcsi logs              # Stream logs
kcsi attach            # Exec into a pod
kcsi describe          # Describe a resource
kcsi events            # Check recent events
kcsi port-forward      # Port-forward to a service or pod
kcsi get               # Get resource status
kcsi delete            # Delete a resource (with confirmation)
kcsi rollout status    # Check rollout status
kcsi rollout restart   # Restart a deployment
kcsi debug             # Debug a pod (ephemeral container)
kcsi diag              # Generate diagnostics for issue reporting
```

**Every command supports TAB completion.** Start typing, press TAB, select from the list.

---

## Features

- **Cascading TAB selection** – namespace → resource → pod → container. One smooth flow, no typing.
- **Guardrails for safety** – Destructive ops require explicit confirmation. No accidental production disasters.
- **Day-2 ops workflows** – `kcsi events` for cluster activity, `kcsi check errors` to surface failing pods, `kcsi get pvc pods` for storage troubleshooting, `kcsi dig` for DNS debugging.
- **kubectl muscle memory compatible** – Same verbs (`get`, `describe`, `logs`, `attach`), same mental model. Zero learning curve.
- **Cross-platform single binary** – Written in Go. No dependencies, no runtime, no containers.

### 🌐 New in v0.7.0 — Multi-Cluster Context Management

Manage multiple Kubernetes clusters without modifying your system kubeconfig:

- **Isolated context storage** in `~/.kcsi/contexts/` — your `~/.kube/config` stays untouched
- **Import or reference kubeconfigs** with `kcsi context import` / `kcsi context add`
- **Switch contexts instantly** with `kcsi context use` — all kcsi commands respect the active context
- **Full command set**: `import`, `add`, `list`, `use`, `current`, `remove`

[See full context management documentation in Advanced section below](#advanced) | [View v0.7.0 changelog](https://github.com/stanzinofree/kcsi/blob/main/CHANGELOG.md#070---2026-01-08)

---

## Safety & Security notes

- **Confirmation prompts** on destructive actions (delete, drain, rollout restart)
- **Read-only by default** for most commands (logs, describe, get, events)
- **Respects your kubeconfig** – uses the same context and credentials as `kubectl`
- **No telemetry by default** – KCSI does not include analytics or tracking. Review the source if you need strict compliance guarantees.
- **Open source** – audit the code yourself
- **Security warnings** when displaying decoded secrets – see [docs/SECURITY_SECRETS.md](docs/SECURITY_SECRETS.md)

---

## Documentation

- **Full docs**: [https://stanzinofree.github.io/kcsi/](https://stanzinofree.github.io/kcsi/)
- **Cheatsheet**: [https://stanzinofree.github.io/kcsi/cheatsheet.html](https://stanzinofree.github.io/kcsi/cheatsheet.html)
- **Roadmap**: [https://stanzinofree.github.io/kcsi/roadmap.html](https://stanzinofree.github.io/kcsi/roadmap.html)
- **Support**: [https://stanzinofree.github.io/kcsi/support.html](https://stanzinofree.github.io/kcsi/support.html)

---

## Support KCSI ☕️

KCSI is free and open source. If it saves you time, consider supporting:

- **[Buy Me a Coffee](https://buymeacoffee.com/smilzao)** – one-time support
- **[GitHub Sponsors](https://github.com/sponsors/stanzinofree)** – recurring sponsorship

**For teams:**  
60-minute onboarding workshop + custom command/alias pack + guardrail/preset suggestions available. Annual sponsors can prioritize feature requests and triage. Reach out via GitHub or sponsors page.

---

## Diagnostics

When reporting issues or requesting support, use the `kcsi diag` command to generate a comprehensive diagnostics report:

```bash
kcsi diag
```

**What it includes:**
- KCSI version information
- System details (OS, architecture, Go version)
- Kubernetes context and namespace
- kubectl version and availability
- Terminal environment

**Safety:** The command does NOT print secrets, tokens, or kubeconfig contents. Only safe metadata is included.

**Options:**
```bash
kcsi diag --cluster    # Include cluster reachability check
kcsi diag --strict     # Exit with error if any check fails
```

**For support:**
1. Run `kcsi diag` and copy the output
2. Open a GitHub issue at https://github.com/stanzinofree/kcsi/issues
3. Paste the diagnostics output in your issue
4. Sponsors can request the `sponsor-priority` label for faster triage

---

## Advanced

<details>
<summary><strong>Logs & exec</strong></summary>

**Get pods with namespace autocomplete**
```bash
kcsi get pods -n <TAB>
kcsi get pods -n kube-system
```

**Stream logs with cascading autocomplete**
```bash
kcsi logs -n <TAB>
kcsi logs -n kube-system <TAB>
kcsi logs -f -n kube-system my-pod
kcsi logs --tail 100 -n kube-system my-pod
kcsi logs -n kube-system my-pod -c <TAB>
```

**Monitor cluster events**
```bash
kcsi events
kcsi events -n production
kcsi events -w
```

**Check for pod errors**
```bash
kcsi check errors
```

</details>

<details>
<summary><strong>Safe delete & confirmations</strong></summary>

**Delete resources with confirmation**
```bash
kcsi delete pod -n <TAB>
kcsi delete pod -n default <TAB>
kcsi delete pod -n default my-pod
# Output: Are you sure you want to delete pod 'my-pod' in namespace 'default'? [y/N]:
```

**Force delete (skip confirmation)**
```bash
kcsi delete pod -n default my-pod --force
```

</details>

<details>
<summary><strong>Debug & port-forward</strong></summary>

**Debug pods with ephemeral containers**
```bash
kcsi debug -n production my-pod
# Features:
# - Automatic internet connectivity check
# - Smart image selection (netshoot → alpine → busybox)
# - Full networking and debugging toolkit
```

**Port forwarding**
```bash
kcsi port-forward -n default my-pod 8080:80
# Features:
# - Root privilege check for ports < 1024
# - Port availability check before forwarding
```

</details>

<details>
<summary><strong>Secrets & rollout</strong></summary>

**View and decode secrets**
```bash
kcsi get secrets decoded my-secret -n production
kcsi get secrets show my-secret -n production -k api-key
# ⚠️ Security Note: See docs/SECURITY_SECRETS.md
```

**Rollout management**
```bash
kcsi rollout restart deployment my-app -n production
kcsi rollout status deployment my-app -n production
kcsi rollout history deployment my-app -n production
kcsi rollout undo deployment my-app -n production
kcsi rollout undo deployment my-app -n production --to-revision=3
```

</details>

<details>
<summary><strong>Apply & edit</strong></summary>

**Apply configurations**
```bash
kcsi apply -f deployment.yaml -n production
kcsi apply -f ./k8s-manifests --recursive -n production
kcsi apply -k ./overlays/production
kcsi apply -f deployment.yaml -n production --dry-run
```

**Edit resources with automatic backup**
```bash
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
