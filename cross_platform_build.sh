#!/bin/bash

# Project name (change this to your project name)
BINARY_NAME=arbitrary-precision-calculator

# Platforms to build
PLATFORMS=("windows" "linux" "darwin")
ARCHITECTURES=("amd64" "arm64")

# Create a build directory
mkdir -p build

# Build for each platform and architecture
for OS in "${PLATFORMS[@]}"; do
    for ARCH in "${ARCHITECTURES[@]}"; do
        # Set output filename based on OS and architecture
        if [ "$OS" == "windows" ]; then
            OUTPUT="build/${BINARY_NAME}_${OS}_${ARCH}.exe"
        else
            OUTPUT="build/${BINARY_NAME}_${OS}_${ARCH}"
        fi

        # Build the binary
        echo "Building for ${OS}/${ARCH}"
        GOOS=$OS GOARCH=$ARCH go build -o "$OUTPUT"
    done
done

echo "Build complete. Binaries are in the 'build' directory."