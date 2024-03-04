#!/bin/bash

#  Configuration

INSTALL_DIR="/usr/local/bin"
CONFIG_DIR="$HOME/.config/tchat"
CLIENT_DIR="./client" # Directory containing the client source code

# utils

print() {
    BLUE='\033[0;34m'
    GREEN='\033[0;32m'
    RED='\033[0;31m'
    NC='\033[0m' # No Color

    if [ "$1" = "error" ]; then
        echo "${BLUE}tChat >> ${RED}$2${NC}"
    else
        echo "${BLUE}tChat >> ${GREEN}$2${NC}"
    fi
}

#  Error Handling

set -e # Exit immediately on errors

function handle_error() {
    print "error" "Installation failed. Please check the error messages above."
    exit 1
}
trap handle_error ERR

# Permission Handling

# Check for root privileges if needed for your installation
print "info" "Checking for root privileges..."

if [[ $EUID -ne 0 ]]; then
    print "error" "Please run this script with sudo or as root"
    exit 1
fi

# Check OS (install only on mac and linux)
print "info" "Checking OS..."

if [[ "$(uname -s)" != "Darwin" ]] && [[ "$(uname -s)" != "Linux" ]]; then
    print "error" "This script is only supported on macOS and Linux systems."
    exit 1
fi

# Check Dependencies

# check for make
print "info" "Checking for make..."

if ! command -v make &>/dev/null; then
    print "error" "make could not be found. Please install it."
    exit 1
fi

# check for git
print "info" "Checking for git..."

if ! command -v git &>/dev/null; then
    print "error" "git could not be found. Please install it."
    exit 1
fi

# check for golang
print "info" "Checking for go..."

if ! command -v go &>/dev/null; then
    print "error" "go could not be found. Please install it."
    exit 1
fi

# Main Installation Logic
print "info" "Building client..."

cd "$CLIENT_DIR" && make build # Build the client

print "info" "Build successful!"
print "info" "Installing tchat..."

print "info" "Creating configuration directory: $CONFIG_DIR"

mkdir -p "$CONFIG_DIR"
chmod 755 "$CONFIG_DIR"

cp ./bin/tchat "$INSTALL_DIR"

print "info" "Installation successful!"
