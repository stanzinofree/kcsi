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
  <a href="https://buymeacoffee.com/smilzao"><img alt="Buy Me A Coffee" src="https://img.shields.io/badge/buy%20me%20a%20coffee-support-yellow"></a>
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



⸻

Features • Installation • Quick Start • Commands • Documentation • Support • Contributing

⸻

Why KCSI?

You know Kubernetes. You just don’t remember the exact command syntax every time.

KCSI reduces context switching when you operate clusters intermittently. Instead of searching docs, copying commands, and fixing flags by hand, you get a guided flow:
	•	Cascading TAB completion: namespace → resource → pod → container
	•	Guardrails: confirmation prompts for destructive actions
	•	Day-2 ops shortcuts: events, rollout, debug, error checks, storage helpers
	•	Same kubectl mental model: no paradigm shift, no TUI required
	•	Cross-platform single binary (Go)

⸻

Features

What makes KCSI different:
	•	Cascading TAB selection: namespace → resource → pod → container, one smooth flow
	•	Guardrails for safety: destructive ops require explicit confirmation (unless you force it)
	•	Day-2 ops workflows built-in: events, rollout, debug, check errors, pvc helpers, dns tools
	•	Works with kubectl muscle memory: familiar verbs and mental model
	•	Cross-platform single binary: macOS, Linux, Windows — no runtime, no containers

⸻

Installation

Option A — Install script (recommended)

macOS / Linux:

```bash
curl -sSL https://raw.githubusercontent.com/stanzinofree/kcsi/main/install.sh | bash
```

Option B — Download a prebuilt binary (GitHub Releases)

Download the latest binary for your OS/ARCH from the Releases page, then put it in your PATH.

Example (macOS / Linux):

```bash
chmod +x kcsi
sudo mv kcsi /usr/local/bin/kcsi
kcsi –help
```

Windows:
	•	Download the latest kcsi-windows-*.exe
	•	Put it in a folder inside PATH (or add that folder to PATH)

Option C — Build from source

```bash
git clone https://github.com/stanzinofree/kcsi.git
cd kcsi
go build -o kcsi
./kcsi –help
```

Enable shell completion

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

⸻

Quick Start (30 seconds)

```bash

Check installation

kcsi version

Try it

kcsi logs

TAB to select namespace

TAB to select pod

TAB to select container (if needed)

```

That’s it. You’re ready.

⸻

The 10 commands you’ll actually use

```bash

1) Stream logs (TAB completes namespace/pod/container)

kcsi logs

2) Exec / interactive access

kcsi attach

3) Describe resources

kcsi describe

4) Cluster activity (events)

kcsi events

5) Port-forward

kcsi port-forward

6) Get resources

kcsi get

7) Delete resources (with confirmation)

kcsi delete

8) Rollout status

kcsi rollout status

9) Rollout restart

kcsi rollout restart

10) Debug pod (ephemeral container)

kcsi debug
```

Every command supports TAB completion: start typing, press TAB, select from the list.

⸻

Safety & Security notes
	•	Confirmation prompts on destructive actions (delete, drain, rollout restart)
	•	Read-only by default for most commands (logs, describe, get, events)
	•	Uses your kubeconfig context and credentials (same access model as kubectl)
	•	No telemetry: KCSI does not collect usage data
	•	Secret helpers may print sensitive values in plain text — be mindful of screen sharing and shell history

For secret-related notes, see: docs/SECURITY_SECRETS.md

⸻

Documentation
	•	Full docs: https://stanzinofree.github.io/kcsi/
	•	Cheatsheet: https://stanzinofree.github.io/kcsi/cheatsheet/
	•	Roadmap: https://stanzinofree.github.io/kcsi/roadmap/

⸻

Support KCSI ☕️

If KCSI saves you time (or prevents an “oops” in production), you can support the project:
	•	Buy Me a Coffee: https://buymeacoffee.com/smilzao
	•	GitHub Sponsors: https://github.com/sponsors/stanzinofree

Workshops and customization packs are available for teams (aliases, guardrails, presets).

⸻

Contributing

PRs are welcome — small and focused changes are best.

Suggested conventions:
	•	feat: new features
	•	fix: bug fixes
	•	docs: documentation changes
	•	chore: maintenance tasks
	•	test: tests

Please ensure CI checks pass before requesting review.

⸻

License

MIT — see LICENSE
