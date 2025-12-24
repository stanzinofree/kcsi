#!/bin/bash

# Build script for kcsi - creates binaries for multiple platforms

set -e

VERSION="0.3.0"
APP_NAME="kcsi"
BUILD_DIR="bin"

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}Building ${APP_NAME} v${VERSION}${NC}"

# Clean previous builds
rm -rf ${BUILD_DIR}
mkdir -p ${BUILD_DIR}

# Build for different platforms
platforms=(
    "darwin/amd64"
    "darwin/arm64"
    "linux/amd64"
    "linux/arm64"
    "linux/arm"
)

for platform in "${platforms[@]}"; do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    
    output_name="${BUILD_DIR}/${APP_NAME}-${GOOS}-${GOARCH}"
    
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi
    
    echo -e "${GREEN}Building for ${GOOS}/${GOARCH}...${NC}"
    
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name -ldflags="-s -w" .
    
    if [ $? -ne 0 ]; then
        echo "An error occurred building for ${GOOS}/${GOARCH}! Aborting..."
        exit 1
    fi
done

echo -e "${GREEN}Build complete!${NC}"
echo ""
echo "Binaries created in ${BUILD_DIR}:"
ls -lh ${BUILD_DIR}
