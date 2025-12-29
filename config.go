package main

import (
	"bufio"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Style defines the visual appearance of the menu
type Style struct {
	Key       string `yaml:"key"`       // Key color (ANSI 256 or hex)
	Label     string `yaml:"label"`     // Label color
	Title     string `yaml:"title"`     // Title color
	Submenu   string `yaml:"submenu"`   // Submenu indicator color
	Separator string `yaml:"separator"` // Separator color
	Border    string `yaml:"border"`    // Border style: rounded, double, thick, hidden
	Padding   string `yaml:"padding"`   // Padding inside box
}

// Layout defines size constraints
type Layout struct {
	MaxWidth int `yaml:"max_width"` // Max label width before truncation
	MaxLines int `yaml:"max_lines"` // Max visible lines before scrolling
}

// Item represents a menu item
type Item struct {
	Name      string `yaml:"name"`      // Display label
	Key       string `yaml:"key"`       // Key trigger
	Command   string `yaml:"command"`   // Shell command
	Tmux      string `yaml:"tmux"`      // Tmux command
	Value     string `yaml:"value"`     // Output value
	Menu      []Item `yaml:"menu"`      // Submenu items
	Separator bool   `yaml:"separator"` // Is separator
	Transient bool   `yaml:"transient"` // Stay open after select
}

// Config represents the full configuration
type Config struct {
	Title  string `yaml:"title"`
	Style  Style  `yaml:"style"`
	Layout Layout `yaml:"layout"`
	Items  []Item `yaml:"items"`
}

// LoadConfig loads configuration from a YAML file
func LoadConfig(path string) (Config, error) {
	var config Config

	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	return config, err
}

// ParseItems parses menu items from a reader in format: key:label or key:label:value
func ParseItems(r io.Reader) []Item {
	var items []Item
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 3)
		if len(parts) >= 2 {
			item := Item{
				Key:  strings.TrimSpace(parts[0]),
				Name: strings.TrimSpace(parts[1]),
			}
			if len(parts) == 3 {
				item.Value = strings.TrimSpace(parts[2])
			}
			items = append(items, item)
		}
	}

	return items
}

// ParseStdin parses menu items from stdin
func ParseStdin() []Item {
	return ParseItems(os.Stdin)
}
