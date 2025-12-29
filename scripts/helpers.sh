#!/usr/bin/env bash
#
# helpers.sh - Shared utility functions for bubble-keys plugin
#

# Get a tmux option value with a default fallback
get_option() {
    local option="$1"
    local default_value="$2"
    local option_value

    option_value=$(tmux show-option -gqv "$option" 2>/dev/null)

    if [ -z "$option_value" ]; then
        echo "$default_value"
    else
        echo "$option_value"
    fi
}

# Get plugin directory (where the plugin is installed)
get_plugin_dir() {
    local script_path="${BASH_SOURCE[1]:-${BASH_SOURCE[0]}}"
    cd "$(dirname "$script_path")/.." && pwd
}
