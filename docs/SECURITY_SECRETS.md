# Security Considerations for Secrets Commands

## Overview

The `kcsi get secrets` commands provide read-only access to Kubernetes secrets with automatic base64 decoding. While these commands are designed to be safe and useful for debugging, users should be aware of security implications.

## Commands

### `kcsi get secrets decoded`
Displays all keys and values of a secret with base64 decoding applied.

### `kcsi get secrets show -k <key>`
Displays the decoded value of a specific key in a secret.

## Security Features

### ✅ What We Do to Keep It Safe

1. **Read-Only Operations**: These commands never modify secrets, only read them
2. **No Shell Injection**: Arguments are passed safely to kubectl
3. **Error Handling**: Proper validation of input and error messages
4. **Security Warnings**: Clear warnings when displaying sensitive data
5. **RBAC Respected**: Uses your kubectl context and respects Kubernetes RBAC permissions

## Security Considerations

### ⚠️ Things to Be Aware Of

1. **Plain Text Display**
   - Secret values are displayed in plain text on your terminal
   - **Risk**: Anyone looking at your screen can see the values
   - **Mitigation**: Only use in private, secure environments

2. **Terminal History**
   - Commands (but not output) are saved in shell history
   - **Risk**: Command history might be accessible to others
   - **Mitigation**: Use private terminal sessions

3. **Screen Recording**
   - Terminal sessions might be recorded by screen capture software
   - **Risk**: Sensitive data could be captured in recordings
   - **Mitigation**: Be aware of any recording software running

4. **Shared Terminal Sessions**
   - Tools like tmux, screen, or terminal sharing might expose secrets
   - **Risk**: Other users could see the secret values
   - **Mitigation**: Don't use these commands in shared sessions

## Best Practices

### ✅ Do:
- Use these commands in private, secure terminal sessions
- Clear your screen after viewing secrets (`clear` command)
- Be aware of who can see your screen
- Use `show -k` for single keys rather than `decoded` when possible (less exposure)
- Verify your kubectl context before viewing secrets

### ❌ Don't:
- Use these commands in shared terminal sessions (tmux, screen)
- Use these commands while screen sharing
- Use these commands in recorded demonstrations
- Leave secrets displayed on screen unattended
- Use these commands in CI/CD pipelines (use kubectl directly with proper secret management)

## Comparison with kubectl

These commands are equivalent to:

```bash
# kcsi get secrets decoded -n <ns> <secret>
# is similar to:
kubectl get secret <secret> -n <ns> -o json | jq '.data | map_values(@base64d)'

# kcsi get secrets show -n <ns> <secret> -k <key>
# is similar to:
kubectl get secret <secret> -n <ns> -o jsonpath='{.data.<key>}' | base64 -d
```

The main difference is convenience and better error handling.

## When NOT to Use These Commands

- **Production Debugging in Shared Environments**: Use kubectl audit logs instead
- **Automated Scripts**: Use proper secret management tools (Vault, Sealed Secrets, etc.)
- **CI/CD Pipelines**: Use native Kubernetes secret injection
- **Documentation/Training**: Use dummy secrets, not real ones

## Alternatives for Different Use Cases

### For Production Debugging:
- Use Kubernetes audit logs
- Use secret scanning tools
- Use proper secret management platforms (HashiCorp Vault, AWS Secrets Manager, etc.)

### For Development:
- These commands are perfect for local development and debugging
- Use them to quickly verify secret values in dev/test environments

### For Automation:
- Use native Kubernetes secret injection in pod specs
- Use External Secrets Operator for syncing from secret stores
- Use Sealed Secrets for GitOps workflows

## Reporting Security Issues

If you find security vulnerabilities in these commands, please report them via GitHub Issues.
