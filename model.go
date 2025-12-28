package main

import tea "github.com/charmbracelet/bubbletea"

// Model represents the TUI state
type Model struct {
	config     Config
	stack      [][]Item // menu stack for nesting
	current    []Item
	selected   string
	runCommand string
	isTmuxCmd  bool
	quitting   bool
	transient  bool
}

// NewModel creates a new model from config
func NewModel(config Config) Model {
	return Model{
		config:  config,
		current: config.Items,
	}
}

// Init implements tea.Model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()

		// Back/quit
		if key == "esc" || key == "ctrl+c" {
			if len(m.stack) > 0 {
				// Pop back to parent menu
				m.current = m.stack[len(m.stack)-1]
				m.stack = m.stack[:len(m.stack)-1]
				return m, nil
			}
			m.quitting = true
			return m, tea.Quit
		}

		if key == "q" && len(m.stack) == 0 {
			m.quitting = true
			return m, tea.Quit
		}

		// Find matching item
		for _, item := range m.current {
			if item.Separator {
				continue
			}
			if item.Key == key {
				// Submenu
				if len(item.Menu) > 0 {
					m.stack = append(m.stack, m.current)
					m.current = item.Menu
					return m, nil
				}

				// Tmux command (priority over shell command)
				if item.Tmux != "" {
					m.runCommand = item.Tmux
					m.isTmuxCmd = true
					return m, tea.Quit
				}

				// Shell command
				if item.Command != "" {
					m.runCommand = item.Command
					m.isTmuxCmd = false
					return m, tea.Quit
				}

				// Value output
				if item.Value != "" {
					m.selected = item.Value
				} else {
					m.selected = item.Name
				}

				if item.Transient {
					m.transient = true
					return m, nil
				}

				return m, tea.Quit
			}
		}
	}
	return m, nil
}

// Result returns the selected value
func (m Model) Result() string {
	return m.selected
}

// Command returns the command to run and whether it's a tmux command
func (m Model) Command() (string, bool) {
	return m.runCommand, m.isTmuxCmd
}

// Quitting returns whether the user quit without selection
func (m Model) Quitting() bool {
	return m.quitting
}
