#!/bin/bash

# ------------------------
#  Configuration
# ------------------------

# Installation directory for chat binary (modify as needed)
INSTALL_DIR="/usr/local/bin"

# Configuration directory (follows conventions for user config)
CONFIG_DIR="$HOME/.config/tchat"

# Client directory
CLIENT_DIR="./client"

# ------------------------
#  Error Handling
# ------------------------

set -e # Exit immediately on errors

function handle_error() {
    echo "Installation failed. Please check the error messages above."
    exit 1
}
trap handle_error ERR

# ------------------------
# Permission Handling
# ------------------------

# Check for root privileges if needed for your installation
if [[ $EUID -ne 0 ]]; then
   echo "Please run this script with sudo or as root"
   exit 1
fi

# ------------------------
# Dependency Handling
# ------------------------

# Example: Check for a hypothetical 'networkutil' dependency
#if ! command -v networkutil &> /dev/null; then
#    echo "The 'networkutil' package is required. Please install it first."
#    exit 1
#fi

# ------------------------
# Main Installation Logic
# ------------------------

# Build the chat binary
cd "$CLIENT_DIR" && make build

# Create configuration directory if it doesn't exist
mkdir -p "$CONFIG_DIR"

# change the permission of the config directory
chmod 755 "$CONFIG_DIR"

# Copy chat binary
cp ./bin/tchat "$INSTALL_DIR"

## Example: Copy a default config file if desired
#if [[ ! -f "$CONFIG_DIR/config.yaml" ]]; then
##    cp ./default-config.yaml "$CONFIG_DIR/config.yaml"
#
#    echo "Creating a default config file..."
#    touch "$CONFIG_DIR/config.yaml"
#    chmod 777 "$CONFIG_DIR/config.yaml"
#fi

echo "Installation successful!"
echo "Configuration files can be found at: $CONFIG_DIR"
