package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	configPath := flag.String("config", "", "Path to YAML config file")
	title := flag.String("title", "", "Menu title")
	flag.Parse()

	// Initialize default styles
	InitStyles()

	var config Config

	if *configPath != "" {
		var err error
		config, err = LoadConfig(*configPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Check for stdin
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			config.Items = ParseStdin()
		} else {
			fmt.Fprintln(os.Stderr, "Usage: bubble-keys --config config.yaml")
			fmt.Fprintln(os.Stderr, "   or: echo 'key:label' | bubble-keys")
			os.Exit(1)
		}
	}

	// Override title from flag
	if *title != "" {
		config.Title = *title
	}

	// Apply config styles
	ApplyStyles(config.Style)

	if len(config.Items) == 0 {
		fmt.Fprintln(os.Stderr, "No items provided")
		os.Exit(1)
	}

	// Run TUI
	m := NewModel(config)
	p := tea.NewProgram(m)
	finalModel, err := p.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Handle result
	if fm, ok := finalModel.(Model); ok {
		command, isTmux := fm.Command()
		if command != "" {
			var cmd *exec.Cmd
			if isTmux {
				args := strings.Fields(command)
				cmd = exec.Command("tmux", args...)
			} else {
				cmd = exec.Command("sh", "-c", command)
			}
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			cmd.Run()
		} else if result := fm.Result(); result != "" {
			fmt.Println(result)
		}
	}
}
