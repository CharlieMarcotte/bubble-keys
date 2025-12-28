package main

import "github.com/charmbracelet/lipgloss"

var (
	keyStyle       lipgloss.Style
	labelStyle     lipgloss.Style
	separatorStyle lipgloss.Style
	titleStyle     lipgloss.Style
	submenuStyle   lipgloss.Style
	boxStyle       lipgloss.Style
)

// InitStyles sets up default styles
func InitStyles() {
	keyStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	labelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	separatorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).Bold(true)
	submenuStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("222"))
	boxStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1)
}

// ApplyStyles applies config-defined style overrides
func ApplyStyles(s Style) {
	if s.Key != "" {
		keyStyle = keyStyle.Foreground(lipgloss.Color(s.Key))
	}
	if s.Label != "" {
		labelStyle = labelStyle.Foreground(lipgloss.Color(s.Label))
	}
	if s.Title != "" {
		titleStyle = titleStyle.Foreground(lipgloss.Color(s.Title))
	}
	if s.Submenu != "" {
		submenuStyle = submenuStyle.Foreground(lipgloss.Color(s.Submenu))
	}
	if s.Separator != "" {
		separatorStyle = separatorStyle.Foreground(lipgloss.Color(s.Separator))
	}
	if s.Border != "" {
		switch s.Border {
		case "rounded":
			boxStyle = boxStyle.Border(lipgloss.RoundedBorder())
		case "double":
			boxStyle = boxStyle.Border(lipgloss.DoubleBorder())
		case "thick":
			boxStyle = boxStyle.Border(lipgloss.ThickBorder())
		case "hidden", "none":
			boxStyle = boxStyle.Border(lipgloss.HiddenBorder())
		case "normal":
			boxStyle = boxStyle.Border(lipgloss.NormalBorder())
		}
	}
}
