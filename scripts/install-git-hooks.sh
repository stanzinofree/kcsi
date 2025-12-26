#!/bin/bash
# Install git hooks for the project

HOOKS_DIR="$(git rev-parse --git-dir)/hooks"
SCRIPTS_DIR="$(dirname "$0")/git-hooks"

echo "Installing git hooks..."

# Install pre-commit hook
if [ -f "$SCRIPTS_DIR/pre-commit" ]; then
    cp "$SCRIPTS_DIR/pre-commit" "$HOOKS_DIR/pre-commit"
    chmod +x "$HOOKS_DIR/pre-commit"
    echo "✓ Installed pre-commit hook (gofmt)"
fi

echo "Git hooks installed successfully!"
echo ""
echo "The pre-commit hook will automatically format Go files with gofmt before each commit."
