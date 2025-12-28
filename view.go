package main

import (
	"fmt"
	"strings"
)

// View implements tea.Model
func (m Model) View() string {
	if m.quitting || (m.selected != "" && !m.transient) || m.runCommand != "" {
		return ""
	}

	var lines []string

	// Title (only at root level)
	if m.config.Title != "" && len(m.stack) == 0 {
		lines = append(lines, titleStyle.Render(m.config.Title))
		lines = append(lines, "")
	}

	// Menu items
	for _, item := range m.current {
		if item.Separator {
			lines = append(lines, separatorStyle.Render("  ──────────"))
			continue
		}

		label := item.Name
		if len(item.Menu) > 0 {
			label = submenuStyle.Render(item.Name)
		}

		line := fmt.Sprintf("  %s  %s", keyStyle.Render(item.Key), labelStyle.Render(label))
		lines = append(lines, line)
	}

	// Back hint for nested menus
	if len(m.stack) > 0 {
		lines = append(lines, "")
		lines = append(lines, separatorStyle.Render("  esc: back"))
	}

	return boxStyle.Render(strings.Join(lines, "\n")) + "\n"
}
