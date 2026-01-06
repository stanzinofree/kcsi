#!/bin/bash

# Build script for kcsi - creates binaries for multiple platforms

set -e

# Read version and app name from version.yaml
VERSION_FILE="pkg/version/version.yaml"

if [[ ! -f "$VERSION_FILE" ]]; then
    echo "Error: $VERSION_FILE not found!" >&2
    exit 1
fi

# Parse VERSION and APP_NAME from YAML using grep and awk
VERSION=$(grep '^version:' "$VERSION_FILE" | awk '{print $2}')
APP_NAME=$(grep '^name:' "$VERSION_FILE" | awk '{print $2}')

if [[ -z "$VERSION" ]] || [[ -z "$APP_NAME" ]]; then
    echo "Error: Could not parse version or name from $VERSION_FILE" >&2
    exit 1
fi

BUILD_DIR="bin"

# Get build information
BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Try to get version from git tag first, fall back to VERSION from version.yaml
GIT_TAG=$(git describe --tags --exact-match 2>/dev/null || echo "")
if [[ -n "$GIT_TAG" ]]; then
    # Remove 'v' prefix if present (v0.6.3 -> 0.6.3)
    BUILD_VERSION="${GIT_TAG#v}"
    echo "Using version from git tag: ${GIT_TAG} -> ${BUILD_VERSION}"
else
    # Fall back to version.yaml
    BUILD_VERSION="${VERSION}"
    echo "Using version from ${VERSION_FILE}: ${BUILD_VERSION}"
fi

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}Building ${APP_NAME} v${BUILD_VERSION}${NC}"
echo -e "${YELLOW}Build Date: ${BUILD_DATE}${NC}"
echo -e "${YELLOW}Git Commit: ${GIT_COMMIT}${NC}"
echo ""

# Clean previous builds
rm -rf ${BUILD_DIR}
mkdir -p ${BUILD_DIR}

# Prepare ldflags to inject build information
LDFLAGS="-s -w -X 'github.com/stanzinofree/kcsi/pkg/version.version=${BUILD_VERSION}' -X 'github.com/stanzinofree/kcsi/pkg/version.buildDate=${BUILD_DATE}' -X 'github.com/stanzinofree/kcsi/pkg/version.gitCommit=${GIT_COMMIT}'"

# Build for different platforms
platforms=(
    "darwin/amd64"
    "darwin/arm64"
    "linux/amd64"
    "linux/arm64"
    "linux/arm"
    "windows/amd64"
    "windows/arm64"
)

for platform in "${platforms[@]}"; do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    output_name="${BUILD_DIR}/${APP_NAME}-${GOOS}-${GOARCH}"

    if [[ $GOOS == "windows" ]]; then
        output_name+='.exe'
    fi

    echo -e "${GREEN}Building for ${GOOS}/${GOARCH}...${NC}"

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name -ldflags="${LDFLAGS}" .

    if [[ $? -ne 0 ]]; then
        echo "An error occurred building for ${GOOS}/${GOARCH}! Aborting..." >&2
        exit 1
    fi
done

echo -e "${GREEN}Build complete!${NC}"
echo ""
echo "Binaries created in ${BUILD_DIR}:"
ls -lh ${BUILD_DIR}
