#!/usr/bin/env bash
#
# launch.sh - Calculate size and launch popup
#

CURRENT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$CURRENT_DIR/helpers.sh"

CONFIG_PATH=$(get_option "@gum-keys-config" "$HOME/.config/gum-keys/config.yaml")
POSITION=$(get_option "@gum-keys-position" "R,S")

# Calculate dimensions
read -r WIDTH HEIGHT < <("$CURRENT_DIR/calc-size.sh" "$CONFIG_PATH")

# Parse position
POPUP_X=$(echo "$POSITION" | cut -d',' -f1)
POPUP_Y=$(echo "$POSITION" | cut -d',' -f2)
[ -z "$POPUP_Y" ] && POPUP_Y="$POPUP_X"

# Launch popup with calculated size
exec tmux display-popup \
    -E \
    -w "$WIDTH" \
    -h "$HEIGHT" \
    -x "$POPUP_X" \
    -y "$POPUP_Y" \
    "$CURRENT_DIR/popup.sh"
