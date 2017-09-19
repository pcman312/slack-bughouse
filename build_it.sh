#!/bin/bash

# Remove the old binary if it exists
if [ -f slack-bughouse ]; then
    echo "removing old linux binary file"
    rm -f slack-bughouse
fi

# Build a new linux binary for the docker container
echo "compiling new linux binary"
GOOS=linux GOARCH=amd64 go build -o slack-bughouse

# Create the new container
echo "building new docker container"
docker build -t jturpin/slack-bughouse .