package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create temp YAML file
	content := `title: "Test Menu"
style:
  key: "#ff0000"
  border: "rounded"
items:
  - name: Item One
    key: a
    command: "echo one"
  - name: Item Two
    key: b
    value: "two"
  - separator: true
  - name: "+Submenu"
    key: s
    menu:
      - name: Nested
        key: n
        tmux: "new-window"
`

	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test.yaml")
	if err := os.WriteFile(configPath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp config: %v", err)
	}

	config, err := LoadConfig(configPath)
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Verify title
	if config.Title != "Test Menu" {
		t.Errorf("expected title 'Test Menu', got %q", config.Title)
	}

	// Verify style
	if config.Style.Key != "#ff0000" {
		t.Errorf("expected style.key '#ff0000', got %q", config.Style.Key)
	}
	if config.Style.Border != "rounded" {
		t.Errorf("expected style.border 'rounded', got %q", config.Style.Border)
	}

	// Verify items count (including separator)
	if len(config.Items) != 4 {
		t.Errorf("expected 4 items, got %d", len(config.Items))
	}

	// Verify first item
	if config.Items[0].Name != "Item One" {
		t.Errorf("expected first item name 'Item One', got %q", config.Items[0].Name)
	}
	if config.Items[0].Key != "a" {
		t.Errorf("expected first item key 'a', got %q", config.Items[0].Key)
	}
	if config.Items[0].Command != "echo one" {
		t.Errorf("expected first item command 'echo one', got %q", config.Items[0].Command)
	}

	// Verify value field
	if config.Items[1].Value != "two" {
		t.Errorf("expected second item value 'two', got %q", config.Items[1].Value)
	}

	// Verify separator
	if !config.Items[2].Separator {
		t.Error("expected third item to be separator")
	}

	// Verify submenu
	if len(config.Items[3].Menu) != 1 {
		t.Errorf("expected submenu with 1 item, got %d", len(config.Items[3].Menu))
	}
	if config.Items[3].Menu[0].Tmux != "new-window" {
		t.Errorf("expected nested tmux 'new-window', got %q", config.Items[3].Menu[0].Tmux)
	}
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	_, err := LoadConfig("/nonexistent/path/config.yaml")
	if err == nil {
		t.Error("expected error for nonexistent file, got nil")
	}
}

func TestLoadConfig_InvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "invalid.yaml")
	if err := os.WriteFile(configPath, []byte("invalid: yaml: content: ["), 0644); err != nil {
		t.Fatalf("failed to write temp config: %v", err)
	}

	_, err := LoadConfig(configPath)
	if err == nil {
		t.Error("expected error for invalid YAML, got nil")
	}
}

func TestParseItems(t *testing.T) {
	input := `a:Apple
b:Banana
c:Cherry:cherry-value`

	items := ParseItems(strings.NewReader(input))

	if len(items) != 3 {
		t.Fatalf("expected 3 items, got %d", len(items))
	}

	// First item: key:label
	if items[0].Key != "a" || items[0].Name != "Apple" {
		t.Errorf("item 0: expected key='a' name='Apple', got key=%q name=%q", items[0].Key, items[0].Name)
	}
	if items[0].Value != "" {
		t.Errorf("item 0: expected empty value, got %q", items[0].Value)
	}

	// Third item: key:label:value
	if items[2].Key != "c" || items[2].Name != "Cherry" || items[2].Value != "cherry-value" {
		t.Errorf("item 2: expected key='c' name='Cherry' value='cherry-value', got key=%q name=%q value=%q",
			items[2].Key, items[2].Name, items[2].Value)
	}
}

func TestParseItems_EmptyLines(t *testing.T) {
	input := `a:Apple

b:Banana

`
	items := ParseItems(strings.NewReader(input))

	if len(items) != 2 {
		t.Errorf("expected 2 items (empty lines skipped), got %d", len(items))
	}
}

func TestParseItems_Whitespace(t *testing.T) {
	input := "  a  :  Apple  \n  b:Banana:  spaced value  "

	items := ParseItems(strings.NewReader(input))

	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}

	if items[0].Key != "a" || items[0].Name != "Apple" {
		t.Errorf("whitespace not trimmed: key=%q name=%q", items[0].Key, items[0].Name)
	}

	if items[1].Value != "spaced value" {
		t.Errorf("value whitespace not trimmed: %q", items[1].Value)
	}
}

func TestParseItems_InvalidLines(t *testing.T) {
	input := `valid:Item
invalid-no-colon
also:valid`

	items := ParseItems(strings.NewReader(input))

	if len(items) != 2 {
		t.Errorf("expected 2 items (invalid line skipped), got %d", len(items))
	}
}
