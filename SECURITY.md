# Security Policy

## Supported Versions

We release patches for security vulnerabilities. Which versions are eligible for receiving such patches depends on the CVSS v3.0 Rating:

| Version | Supported          |
| ------- | ------------------ |
| 0.5.x   | :white_check_mark: |
| < 0.5   | :x:                |

## Reporting a Vulnerability

If you discover a security vulnerability within kcsi, please send an email to the maintainer at your contact email. All security vulnerabilities will be promptly addressed.

Please include the following information in your report:

- Type of issue (e.g. buffer overflow, SQL injection, cross-site scripting, etc.)
- Full paths of source file(s) related to the manifestation of the issue
- The location of the affected source code (tag/branch/commit or direct URL)
- Any special configuration required to reproduce the issue
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if possible)
- Impact of the issue, including how an attacker might exploit the issue

## Security Measures

kcsi implements the following security measures:

- **No Shell Injection**: All kubectl commands use `exec.Command` with separate arguments
- **Input Validation**: Namespace and resource names are validated via kubectl API
- **No Hardcoded Credentials**: No secrets or credentials in the codebase
- **Dependency Scanning**: Automated with Dependabot and SonarCloud
- **Code Scanning**: Automated with CodeQL on every commit
- **Safe Defaults**: Confirmation prompts for destructive operations

## Security Best Practices

When using kcsi:

1. Always use the latest version
2. Keep your kubectl and Kubernetes cluster up to date
3. Use RBAC to limit permissions for the kubectl context
4. Review autocompletion suggestions before executing
5. Use `--force` flag with caution in automated scripts

## Code Security Analysis

This project uses:

- **CodeQL**: Automated code scanning for security vulnerabilities
- **SonarCloud**: Continuous code quality and security analysis
- **Dependabot**: Automated dependency updates
- **Go Security Checker**: `go vet` and static analysis

## Acknowledgments

We appreciate the security research community for helping keep kcsi and its users safe.
