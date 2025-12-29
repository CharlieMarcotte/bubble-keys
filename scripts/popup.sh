#!/usr/bin/env bash
#
# popup.sh - Run gum-keys in popup
#

set -e

CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PLUGIN_DIR="$(dirname "$CURRENT_DIR")"

source "$CURRENT_DIR/helpers.sh"

BINARY="$PLUGIN_DIR/bin/gum-keys"
CONFIG_PATH=$(get_option "@gum-keys-config" "$HOME/.config/gum-keys/config.yaml")

# Verify binary exists
if [ ! -x "$BINARY" ]; then
    echo "gum-keys: Binary not found at $BINARY"
    echo "Run 'prefix + I' to install plugins."
    read -n 1 -s -r -p "Press any key to close..."
    exit 1
fi

# Verify config exists
if [ ! -f "$CONFIG_PATH" ]; then
    echo "gum-keys: Config not found at $CONFIG_PATH"
    read -n 1 -s -r -p "Press any key to close..."
    exit 1
fi

# Run gum-keys
exec "$BINARY" --config "$CONFIG_PATH"
