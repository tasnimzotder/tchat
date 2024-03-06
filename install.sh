#!/bin/bash

# Configuration (adjust as needed)
INSTALL_DIR="/usr/local/bin"
CONFIG_DIR="$HOME/.config/tchat"
RELEASE_URL="https://github.com/tasnimzotder/tchat/releases/latest/download"
ASSET_DARWIN_ARM64="tchat.darwin.arm64"
ASSET_LINUX_AMD64="tchat.linux.amd64"

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
    case $OS in
    Darwin) echo "$ASSET_DARWIN_ARM64" ;;
    Linux) echo "$ASSET_LINUX_AMD64" ;;
    *) error_exit "Unsupported OS. Only macOS and Linux are supported." ;;
    esac
}

download_and_install() {
    asset=$(determine_asset)
    print_message "Downloading and installing tchat..."
    curl -L "$RELEASE_URL/$asset" -o "$INSTALL_DIR/tchat" || error_exit "Download failed."
    chmod +x "$INSTALL_DIR/tchat" || error_exit "Could not make tchat executable."
}

# Main Script Logic
download_and_install
print_message "Installation successful!"
