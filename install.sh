#!/bin/bash
set -e

# KCSI Installation Script
# Detects OS and architecture, downloads the appropriate binary

VERSION="${KCSI_VERSION:-latest}"
INSTALL_DIR="${KCSI_INSTALL_DIR:-/usr/local/bin}"
REPO="stanzinofree/kcsi"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1"
    exit 1
}

# Detect OS
detect_os() {
    case "$(uname -s)" in
        Linux*)     echo "linux";;
        Darwin*)    echo "darwin";;
        *)          error "Unsupported OS: $(uname -s)";;
    esac
}

# Detect architecture
detect_arch() {
    case "$(uname -m)" in
        x86_64|amd64)   echo "amd64";;
        arm64|aarch64)  echo "arm64";;
        *)              error "Unsupported architecture: $(uname -m)";;
    esac
}

# Get download URL
get_download_url() {
    local os=$1
    local arch=$2
    
    if [ "$VERSION" = "latest" ]; then
        echo "https://github.com/${REPO}/releases/latest/download/kcsi-${os}-${arch}"
    else
        echo "https://github.com/${REPO}/releases/download/${VERSION}/kcsi-${os}-${arch}"
    fi
}

# Check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Main installation
main() {
    info "KCSI Installation Script"
    echo ""
    
    # Check for curl or wget
    if ! command_exists curl && ! command_exists wget; then
        error "curl or wget is required but not installed"
    fi
    
    # Detect system
    OS=$(detect_os)
    ARCH=$(detect_arch)
    info "Detected: ${OS}/${ARCH}"
    
    # Get download URL
    DOWNLOAD_URL=$(get_download_url "$OS" "$ARCH")
    info "Downloading from: ${DOWNLOAD_URL}"
    
    # Create temporary file
    TMP_FILE=$(mktemp)
    
    # Download binary with fallback to amd64 if arch-specific binary doesn't exist
    DOWNLOAD_SUCCESS=false
    
    if command_exists curl; then
        if curl -fsSL "$DOWNLOAD_URL" -o "$TMP_FILE" 2>/dev/null; then
            DOWNLOAD_SUCCESS=true
        elif [ "$ARCH" != "amd64" ]; then
            warn "${ARCH} binary not found, falling back to amd64"
            DOWNLOAD_URL=$(get_download_url "$OS" "amd64")
            info "Downloading from: ${DOWNLOAD_URL}"
            curl -fsSL "$DOWNLOAD_URL" -o "$TMP_FILE" && DOWNLOAD_SUCCESS=true
        fi
    else
        if wget -q "$DOWNLOAD_URL" -O "$TMP_FILE" 2>/dev/null; then
            DOWNLOAD_SUCCESS=true
        elif [ "$ARCH" != "amd64" ]; then
            warn "${ARCH} binary not found, falling back to amd64"
            DOWNLOAD_URL=$(get_download_url "$OS" "amd64")
            info "Downloading from: ${DOWNLOAD_URL}"
            wget -q "$DOWNLOAD_URL" -O "$TMP_FILE" && DOWNLOAD_SUCCESS=true
        fi
    fi
    
    if [ "$DOWNLOAD_SUCCESS" = false ]; then
        error "Download failed - binary not available for your platform"
    fi
    
    # Make executable
    chmod +x "$TMP_FILE"
    
    # Move to install directory
    info "Installing to ${INSTALL_DIR}/kcsi"
    
    if [ -w "$INSTALL_DIR" ]; then
        mv "$TMP_FILE" "${INSTALL_DIR}/kcsi" || error "Installation failed"
    else
        warn "Requesting sudo for installation to ${INSTALL_DIR}"
        sudo mv "$TMP_FILE" "${INSTALL_DIR}/kcsi" || error "Installation failed"
    fi
    
    # Verify installation
    if command_exists kcsi; then
        info "Installation successful!"
        echo ""
        kcsi version
        echo ""
        info "Run 'kcsi --help' to get started"
    else
        error "Installation completed but kcsi not found in PATH. Please add ${INSTALL_DIR} to your PATH"
    fi
}

main "$@"
