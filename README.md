# bubble-keys

A keyboard-driven menu TUI for quick actions. Define menus via YAML config or pipe items through stdin.

Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss).

## Installation

```bash
# Clone and build
git clone https://github.com/cmarcotte/bubble-keys.git
cd bubble-keys
make build

# Install to ~/.local/bin
make install
```

## Usage

### YAML Config

```bash
bubble-keys --config menu.yaml
```

### Stdin (piped input)

```bash
# Format: key:label or key:label:value
echo -e "a:Apple\nb:Banana\nc:Cherry" | bubble-keys
```

### Override title

```bash
bubble-keys --config menu.yaml --title "My Menu"
```

## Configuration

### Example YAML

```yaml
title: "Quick Actions"

style:
  key: "#f5a97f"       # Key color (hex or ANSI 256)
  label: "#cad3f5"     # Label color
  title: "#c6a0f6"     # Title color
  submenu: "#eed49f"   # Submenu indicator color
  separator: "#6e738d" # Separator color
  border: "rounded"    # rounded, double, thick, normal, hidden

items:
  - name: "Open Editor"
    key: e
    command: "nvim ."

  - name: "+Git"
    key: g
    menu:
      - name: Status
        key: s
        command: "git status"
      - name: Lazygit
        key: l
        tmux: "new-window lazygit"

  - separator: true

  - name: "Select Me"
    key: x
    value: "custom-output"
```

### Item Fields

| Field       | Description                              |
|-------------|------------------------------------------|
| `name`      | Display label                            |
| `key`       | Key trigger (single character)           |
| `command`   | Shell command to execute                 |
| `tmux`      | Tmux command (runs as `tmux <command>`)  |
| `value`     | Output value (printed to stdout)         |
| `menu`      | Nested submenu items                     |
| `separator` | If `true`, renders as a separator line   |
| `transient` | If `true`, menu stays open after select  |

### Style Options

| Field       | Values                                        |
|-------------|-----------------------------------------------|
| `key`       | Hex color (`#f5a97f`) or ANSI 256 (`212`)     |
| `label`     | Hex color or ANSI 256                         |
| `title`     | Hex color or ANSI 256                         |
| `submenu`   | Hex color or ANSI 256                         |
| `separator` | Hex color or ANSI 256                         |
| `border`    | `rounded`, `double`, `thick`, `normal`, `hidden` |

## Keybindings

| Key       | Action                    |
|-----------|---------------------------|
| `[key]`   | Activate menu item        |
| `Esc`     | Back / quit               |
| `q`       | Quit (at root menu)       |
| `Ctrl+C`  | Quit immediately          |

## License

MIT
