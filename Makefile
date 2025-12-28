.PHONY: build install fmt clean test

BINARY := gum-keys
INSTALL_DIR := $(HOME)/.local/bin

build:
	go build -o $(BINARY) .

install: build
	@mkdir -p $(INSTALL_DIR)
	cp $(BINARY) $(INSTALL_DIR)/

fmt:
	go fmt ./...

test:
	go test -v ./...

clean:
	rm -f $(BINARY)
