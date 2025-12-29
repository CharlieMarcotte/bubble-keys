#!/usr/bin/env bash
#
# gum-keys.tmux - TPM plugin entry point
#

CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SCRIPTS_DIR="$CURRENT_DIR/scripts"

source "$SCRIPTS_DIR/helpers.sh"

BINARY="$CURRENT_DIR/bin/gum-keys"
GUM_KEYS_KEY=$(get_option "@gum-keys-key" "Space")
GUM_KEYS_CONFIG=$(get_option "@gum-keys-config" "$HOME/.config/gum-keys/config.yaml")

# Check if binary exists
if [ ! -x "$BINARY" ]; then
    tmux display-message "gum-keys: binary not found. Run 'prefix + I' to install plugins."
    exit 0
fi

# Check if config exists, copy example if not
if [ ! -f "$GUM_KEYS_CONFIG" ]; then
    config_dir="$(dirname "$GUM_KEYS_CONFIG")"
    mkdir -p "$config_dir" 2>/dev/null
    if [ -f "$CURRENT_DIR/config.example.yaml" ]; then
        cp "$CURRENT_DIR/config.example.yaml" "$GUM_KEYS_CONFIG"
        tmux display-message "gum-keys: created config at $GUM_KEYS_CONFIG"
    fi
fi

# Bind key to launcher (calculates size dynamically)
tmux bind-key "$GUM_KEYS_KEY" run-shell "$SCRIPTS_DIR/launch.sh"
