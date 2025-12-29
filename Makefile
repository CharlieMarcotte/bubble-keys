.PHONY: build install uninstall fmt clean test

BINARY := gum-keys
BIN_DIR := bin
PLUGIN_DIR := $(HOME)/.tmux/plugins/gum-keys

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY) .

install: build
	@mkdir -p $(HOME)/.tmux/plugins
	@ln -sfn $(PWD) $(PLUGIN_DIR)
	@tmux run-shell $(PLUGIN_DIR)/gum-keys.tmux 2>/dev/null || true
	@echo "Installed. Press prefix + Space to use."

uninstall:
	@rm -f $(PLUGIN_DIR)
	@echo "Uninstalled."

fmt:
	go fmt ./...

test:
	go test -v ./...

clean:
	rm -rf $(BIN_DIR)
