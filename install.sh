#!/bin/bash

# Configuration (adjust as needed)
INSTALL_DIR="/usr/local/bin"
RELEASE_URL="https://github.com/tasnimzotder/tchat/releases/latest/download"
ASSET_DARWIN_ARM64="tchat.darwin.arm64"
ASSET_LINUX_AMD64="tchat.linux.amd64"
ASSET_LINUX_ARM64="tchat.linux.arm64"

# Output with colors
BLUE='\033[0;34m'
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

# Functions
print_message() {
    echo "${BLUE}tChat >> ${GREEN}$1${NC}"
}

error_exit() {
    echo "${BLUE}tChat >> ${RED}$1${NC}" >&2
    exit 1
}

determine_asset() {
    OS="$(uname -s)"
    ARCH=$(uname -m)

    if [ "$OS" = "Darwin" ]; then
        echo "$ASSET_DARWIN_ARM64"
    elif [ "$OS" = "Linux" ]; then
        case $ARCH in
        x86_64) echo "$ASSET_LINUX_AMD64" ;;
        arm64) echo "$ASSET_LINUX_ARM64" ;;
        aarch64) echo "$ASSET_LINUX_ARM64" ;;
        *) error_exit "Unsupported architecture. Only amd64 and arm64 are supported." ;;
        esac
    else
        error_exit "Unsupported OS. Only macOS and Linux are supported."
    fi
}

download_and_install() {
    asset=$(determine_asset)

    print_message "Detected OS: $(uname -s), Arch: $(uname -m)"
    print_message "Downloading and installing tchat..."

    curl -L "$RELEASE_URL/$asset" -o "$INSTALL_DIR/tchat" || error_exit "Download failed."
    chmod +x "$INSTALL_DIR/tchat" || error_exit "Could not make tchat executable."
}

# Main Script Logic
download_and_install
print_message "Installation successful!"
