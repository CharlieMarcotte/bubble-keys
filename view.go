package main

import (
	"fmt"
	"strings"
)

// truncate shortens a string to maxLen with ellipsis
func truncate(s string, maxLen int) string {
	if maxLen <= 0 || len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}

// View implements tea.Model
func (m Model) View() string {
	if m.quitting || (m.selected != "" && !m.transient) || m.runCommand != "" {
		return ""
	}

	var lines []string
	maxWidth := m.config.Layout.MaxWidth

	// Title (only at root level)
	if m.config.Title != "" && len(m.stack) == 0 {
		title := m.config.Title
		if maxWidth > 0 {
			title = truncate(title, maxWidth)
		}
		lines = append(lines, titleStyle.Render(title))
		lines = append(lines, "")
	}

	// Menu items
	for _, item := range m.current {
		if item.Separator {
			lines = append(lines, separatorStyle.Render("────────"))
			continue
		}

		label := item.Name
		if maxWidth > 0 {
			label = truncate(label, maxWidth)
		}
		if len(item.Menu) > 0 {
			label = submenuStyle.Render(label)
		}

		line := fmt.Sprintf("%s %s", keyStyle.Render(item.Key), labelStyle.Render(label))
		lines = append(lines, line)
	}

	// Back hint for nested menus
	if len(m.stack) > 0 {
		lines = append(lines, "")
		lines = append(lines, separatorStyle.Render("esc: back"))
	}

	return boxStyle.Render(strings.Join(lines, "\n"))
}
