#!/bin/bash

GIT_REPO=https://github.com/tasnimzotder/tchat
CONTAINER_NAME=tchat-api
CONTAINER_TAG=latest
CONTAINER_REGISTRY=sjc.vultrcr.com/xtasnim

# WORK_DIR=/app
# working directory is current directory
WORK_DIR=$(pwd)

# error handling
set -e

# check if .env.local file exist and add env variables else exit
if [ -f "$WORK_DIR/.env.local" ]; then
    echo "Checking .env.local file..."

    while IFS= read -r line; do
        export "$line"
    done <"$WORK_DIR/.env.local"
else
    echo ".env.local file not found."
    exit 1
fi

# clone the repository if it doesn't exist
if [ ! -d "$WORK_DIR/tchat" ]; then
    echo "Cloning repository..."
    git clone "$GIT_REPO".git "$WORK_DIR/tchat"
fi

# pull the latest changes (force)
echo "Pulling latest changes..."
cd "$WORK_DIR/tchat" && git pull --force

# create a file acme.json for traefik if it doesn't exist
if [ ! -f "$WORK_DIR/tchat/acme.json" ]; then
    echo "Creating acme.json..."
    touch "$WORK_DIR/tchat/acme.json"
fi

# docker login to the container registry
echo "Logging in to the container registry..."
echo $(docker login https://$CONTAINER_REGISTRY -u $REGISTRY_USERNAME -p $REGISTRY_PASSWORD)

# docker compose up
echo "Deploying container..."
docker compose -f "$WORK_DIR/tchat/compose.yaml" up -d --build

# check docker compose status; if not, exit
if ! docker compose -f "$WORK_DIR/tchat/compose.yaml" ps | grep -q "Up"; then
    echo "Checking docker compose status..."
    exit 1
fi

# exit gracefully
echo "Container deployed successfully!"
exit 0
