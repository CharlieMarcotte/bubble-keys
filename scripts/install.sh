#!/usr/bin/env bash
#
# install.sh - Build bubble-keys binary
#
# TPM automatically runs scripts matching *install*.sh after plugin install/update.
#

set -e

CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PLUGIN_DIR="$(dirname "$CURRENT_DIR")"
BIN_DIR="$PLUGIN_DIR/bin"
BINARY="$BIN_DIR/bubble-keys"

echo "bubble-keys: Building binary..."

# Check for Go
if ! command -v go &>/dev/null; then
    echo "bubble-keys: ERROR - Go is not installed."
    echo "bubble-keys: Install Go via: brew install go"
    exit 1
fi

# Check Go version (require 1.21+)
GO_VERSION=$(go version | grep -oE 'go[0-9]+\.[0-9]+' | sed 's/go//')
GO_MAJOR=$(echo "$GO_VERSION" | cut -d. -f1)
GO_MINOR=$(echo "$GO_VERSION" | cut -d. -f2)

if [ "$GO_MAJOR" -lt 1 ] || { [ "$GO_MAJOR" -eq 1 ] && [ "$GO_MINOR" -lt 21 ]; }; then
    echo "bubble-keys: ERROR - Go 1.21+ required. Found: go$GO_VERSION"
    exit 1
fi

# Create bin directory
mkdir -p "$BIN_DIR"

# Build the binary
cd "$PLUGIN_DIR"
go build -o "$BINARY" .

if [ -x "$BINARY" ]; then
    echo "bubble-keys: Binary built at $BINARY"
else
    echo "bubble-keys: ERROR - Build failed"
    exit 1
fi

echo "bubble-keys: Installation complete."
