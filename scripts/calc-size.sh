#!/usr/bin/env bash
#
# calc-size.sh - Read popup dimensions from config
#

CONFIG_PATH="${1:-$HOME/.config/gum-keys/config.yaml}"

# Defaults
WIDTH=20
HEIGHT=12
BORDER_PAD=2  # top + bottom border

if [ -f "$CONFIG_PATH" ]; then
    # Check if border is none
    border=$(grep -E '^\s*border:' "$CONFIG_PATH" | head -1 | sed 's/.*border:\s*//' | tr -d ' "'"'"'')
    [ "$border" = "none" ] && BORDER_PAD=0

    # Read max_width from config
    w=$(grep -E '^\s*max_width:' "$CONFIG_PATH" | head -1 | sed 's/.*max_width:\s*//' | tr -d ' ')
    [ -n "$w" ] && WIDTH=$((w + 4))  # label + key + space + padding

    # Read max_lines from config
    h=$(grep -E '^\s*max_lines:' "$CONFIG_PATH" | head -1 | sed 's/.*max_lines:\s*//' | tr -d ' ')
    [ -n "$h" ] && HEIGHT=$((h + 3 + BORDER_PAD))  # items + title(2) + buffer(1) + border
fi

echo "$WIDTH $HEIGHT"
