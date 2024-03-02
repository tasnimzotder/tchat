#!/bin/bash

#  Configuration

INSTALL_DIR="/usr/local/bin"
CONFIG_DIR="$HOME/.config/tchat"
CLIENT_DIR="./client" # Directory containing the client source code

#  Error Handling

set -e # Exit immediately on errors

function handle_error() {
    echo "Installation failed. Please check the error messages above."
    exit 1
}
trap handle_error ERR

# Permission Handling

# Check for root privileges if needed for your installation
if [[ $EUID -ne 0 ]]; then
   echo "Please run this script with sudo or as root"
   exit 1
fi

# Main Installation Logic

cd "$CLIENT_DIR" && make build # Build the client

mkdir -p "$CONFIG_DIR"
chmod 755 "$CONFIG_DIR"

cp ./bin/tchat "$INSTALL_DIR"

echo "Installation successful!"
echo "Configuration files can be found at: $CONFIG_DIR"
