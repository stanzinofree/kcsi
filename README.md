<p align="center">
  <img src="logo.png" alt="KCSI Logo" width="140" />
</p>
<h1 align="center">KCSI — kubectl for humans 🐹</h1>
<p align="center">
  <b>Fast day-2 Kubernetes ops when you don’t live in Kubernetes all day.</b><br/>
  Cascading autocomplete, guided selection, and guardrails — so you stop memorizing commands and start shipping fixes.
</p>
<p align="center">
  <a href="https://stanzinofree.github.io/kcsi/"><img alt="Docs" src="https://img.shields.io/badge/docs-online-brightgreen"></a>
  <a href="https://stanzinofree.github.io/kcsi/cheatsheet/"><img alt="Cheatsheet" src="https://img.shields.io/badge/cheatsheet-interactive-blue"></a>
  <a href="https://stanzinofree.github.io/kcsi/roadmap/"><img alt="Roadmap" src="https://img.shields.io/badge/roadmap-public-orange"></a>
  <a href="https://buymeacoffee.com/smilzao"><img alt="Buy Me a Coffee" src="https://img.shields.io/badge/buy%20me%20a%20coffee-support-yellow"></a>
</p>
<p align="center">
  <a href="https://goreportcard.com/report/github.com/stanzinofree/kcsi"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/stanzinofree/kcsi"></a>
  <a href="LICENSE"><img alt="License" src="https://img.shields.io/github/license/stanzinofree/kcsi"></a>
  <a href="https://github.com/stanzinofree/kcsi/releases"><img alt="Release" src="https://img.shields.io/github/v/release/stanzinofree/kcsi"></a>
  <a href="https://github.com/stanzinofree/kcsi/actions"><img alt="Build and Test" src="https://github.com/stanzinofree/kcsi/actions/workflows/build.yml/badge.svg"></a>
  <a href="https://sonarcloud.io/summary/new_code?id=stanzinofree_kcsi"><img alt="Quality Gate Status" src="https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_kcsi&metric=alert_status"></a>
  <a href="https://sonarcloud.io/summary/new_code?id=stanzinofree_kcsi"><img alt="Security Rating" src="https://sonarcloud.io/api/project_badges/measure?project=stanzinofree_kcsi&metric=security_rating"></a>
</p>
<p align="center">
  <i>Your Kapibara buddy in the terminal: calm, practical, always ready to troubleshoot.</i>
</p>
------

## **Table of contents**

- [Why KCSI?](#why-kcsi)![Attachment.tiff](Attachment.tiff)
- [Installation](#installation)![Attachment.tiff](Attachment.tiff)
- [Quick Start (30 seconds)](#quick-start-30-seconds)![Attachment.tiff](Attachment.tiff)
- [The 10 commands you’ll actually use](#the-10-commands-youll-actually-use)![Attachment.tiff](Attachment.tiff)
- [Features](#features)![Attachment.tiff](Attachment.tiff)
- [Safety & Security notes](#safety--security-notes)![Attachment.tiff](Attachment.tiff)
- [Documentation](#documentation)![Attachment.tiff](Attachment.tiff)
- [Support KCSI ☕️](#support-kcsi-️)![Attachment.tiff](Attachment.tiff)
- [Advanced](#advanced)![Attachment.tiff](Attachment.tiff)
- [Contributing](#contributing)![Attachment.tiff](Attachment.tiff)
- [License](#license)![Attachment.tiff](Attachment.tiff)

------

## **Why KCSI?**

You know Kubernetes. You just don’t remember the exact command syntax every time.

KCSI reduces context switching when you operate clusters intermittently. Instead of opening docs, copy/pasting commands, then manually fixing namespaces/pods/containers, you get:
- Cascading TAB autocomplete: namespace → resource → pod → container
- Guardrails on destructive actions: confirmation prompts before “oops” commands
- Day-2 ops shortcuts: events, debug pods, inspect PVCs, rollout status, error checks
- kubectl muscle memory: same verbs, same mental model — just faster
- Cross-platform single binary: macOS, Linux, Windows

Perfect for sysadmins/DevOps/system engineers who touch Kubernetes intermittently but need to move fast when they do.

------

## **Installation**

### **Option A — Install script (recommended)**

macOS / Linux:

```bash
curl -sSL https://raw.githubusercontent.com/stanzinofree/kcsi/main/install.sh | bash
```

### **Option B — Download a prebuilt binary (GitHub Releases)**

Download the latest binary for your OS/ARCH from:
https://github.com/stanzinofree/kcsi/releases

Example (macOS / Linux):

```bash
chmod +x kcsi
sudo mv kcsi /usr/local/bin/kcsi
kcsi version
```

Windows:
- Download the latest kcsi-windows-*.exe from Releases
- Put it in a folder inside PATH (or add that folder to PATH)

### **Option C — Build from source**

```bash
git clone https://github.com/stanzinofree/kcsi.git
cd kcsi
go build -o kcsi
./kcsi version
```

### **Enable shell completion (recommended)**

bash:

```bash
source <(kcsi completion bash)
```

zsh:

```bash
source <(kcsi completion zsh)
```

fish:

```fish
kcsi completion fish | source
```

PowerShell:

```powershell
kcsi completion powershell | Out-String | Invoke-Expression
```

------

## **Quick Start (30 seconds)**

```bash
# **Check installation**
kcsi version
# **Try it (TAB drives the flow)**
kcsi logs
# **→ TAB to select namespace**
# **→ TAB to select pod**
# **→ TAB to select container (if needed)**
```

That’s it. You’re ready.

------

## **The 10 commands you’ll actually use**

```bash
# **1) Stream logs (TAB completes namespace/pod/container)**
kcsi logs
# **2) Exec / interactive access**
kcsi attach
# **3) Describe resources**
kcsi describe
# **4) Cluster activity (events)**
kcsi events
# **5) Port-forward**
kcsi port-forward
# **6) Get resources**
kcsi get
# **7) Delete resources (with confirmation)**
kcsi delete
# **8) Rollout status**
kcsi rollout status
# **9) Rollout restart**
kcsi rollout restart
# **10) Debug pod (ephemeral container)**
kcsi debug
```

Every command supports TAB completion: start typing, press TAB, select from the list. No flags to remember.

------

## **Features**

What makes KCSI different:
- Cascading TAB selection: namespace → resource → pod → container (one smooth flow)
- Guardrails for safety: destructive ops require explicit confirmation (unless you force it)
- Day-2 ops workflows built-in: events, rollout, debug, check errors, pvc helpers, dns tools
- Works with kubectl muscle memory: familiar verbs and mental model (no paradigm shift)
- Cross-platform single binary: Go-based, no runtime, no containers

------

## **Safety & Security notes**

- Confirmation prompts on destructive actions (delete, drain, rollout restart)
- Read-only by default for most commands (logs, describe, get, events)
- Uses your kubeconfig context and credentials (same access model as kubectl)
- No telemetry: KCSI does not implement data collection or phone-home logic
- Secrets helpers may print sensitive values in plain text — be mindful of screen sharing and shell history

Secret handling notes:
docs/SECURITY_SECRETS.md

------

## **Documentation**

- Full docs: https://stanzinofree.github.io/kcsi/
- Cheatsheet: https://stanzinofree.github.io/kcsi/cheatsheet/
- Roadmap: https://stanzinofree.github.io/kcsi/roadmap/

------

## **Support KCSI ☕️**

If KCSI saves you time (or prevents an “oops” in production), you can support the project:
- Buy Me a Coffee (one-time): https://buymeacoffee.com/smilzao
- GitHub Sponsors (recurring): https://github.com/sponsors/stanzinofree

Workshops and customization packs are available for teams (aliases, guardrails, presets).

------





## **Advanced**

<details>
<summary><b>Advanced installation options</b></summary>


macOS / Linux (manual download examples):

```bash
# **macOS (Intel)**
curl -L https://github.com/stanzinofree/kcsi/releases/latest/download/kcsi-darwin-amd64 -o kcsi
# **macOS (Apple Silicon)**
curl -L https://github.com/stanzinofree/kcsi/releases/latest/download/kcsi-darwin-arm64 -o kcsi
# **Linux (amd64)**
curl -L https://github.com/stanzinofree/kcsi/releases/latest/download/kcsi-linux-amd64 -o kcsi
chmod +x kcsi
sudo mv kcsi /usr/local/bin/kcsi
kcsi version
```

</details>
<details>
<summary><b>Setup autocompletion (persistent)</b></summary>

Bash (Linux):

```bash
kcsi completion bash > /etc/bash_completion.d/kcsi
```

Bash (macOS with bash-completion installed):

```bash
kcsi completion bash > /usr/local/etc/bash_completion.d/kcsi
# **path may vary depending on your brew prefix**

```

Zsh:

```bash
echo “autoload -U compinit; compinit” >> ~/.zshrc
kcsi completion zsh > “${fpath[1]}/_kcsi”
```

Fish:

```fish
kcsi completion fish > ~/.config/fish/completions/kcsi.fish
```

PowerShell:

```powershell
# **Add to your PowerShell profile (path in $PROFILE)**
kcsi completion powershell | Out-String | Invoke-Expression
```

</details>
<details>
<summary><b>Usage examples</b></summary>

Get pods with namespace autocomplete:

```bash
kcsi get pods -n 
kcsi get pods -n kube-system
```

Stream logs with cascading autocomplete:

```bash
kcsi logs -n 
kcsi logs -n kube-system 
kcsi logs -f -n kube-system my-pod
kcsi logs –tail 100 -n kube-system my-pod
kcsi logs -n kube-system my-pod -c 
```

Delete resources safely:

```bash
kcsi delete pod -n 
kcsi delete pod -n default 
kcsi delete pod -n default my-pod
# **confirmation prompt**
kcsi delete pod -n default my-pod –force
```

Monitor cluster events:

```bash
kcsi events
kcsi events -n production
kcsi events -w
```

Check for pod errors:

```bash
kcsi check errors
```

Debug pods with ephemeral containers:

```bash
kcsi debug -n production my-pod
```

Port forwarding:

```bash
kcsi port-forward -n default my-pod 8080:80
```

View and decode secrets:

```bash
kcsi get secrets decoded my-secret -n production
kcsi get secrets show my-secret -n production -k api-key
```

Rollout management:

```bash
kcsi rollout restart deployment my-app -n production
kcsi rollout status deployment my-app -n production
kcsi rollout history deployment my-app -n production
kcsi rollout undo deployment my-app -n production
kcsi rollout undo deployment my-app -n production –to-revision=3
```

Apply configurations:

```bash
kcsi apply -f deployment.yaml -n production
kcsi apply -f ./k8s-manifests –recursive -n production
kcsi apply -k ./overlays/production
kcsi apply -f deployment.yaml -n production –dry-run
```

Edit resources with automatic backup:

```bash
kcsi edit deployment my-app -n production
```

</details>

------

## **Contributing**

Contributions are welcome! Small, focused PRs are best.

Suggested conventions:
- feat: new features
- fix: bug fixes
- docs: documentation changes
- chore: maintenance tasks
- test: tests

Basic workflow:

```bash
git checkout main
git pull origin main
git checkout -b fix/your-feature-name
git add .
git commit -m “fix: description of your changes”
git push origin fix/your-feature-name
```

Create a PR via GitHub CLI (optional):

```bash
gh pr create –title “Your PR Title” –body “Detailed description” –base main
```
------

## **License**

MIT — see LICENSE

------

Made with ❤️ by sysadmins, for sysadmins.
