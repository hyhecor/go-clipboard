GO := go
BIN_DIR := bin
CMDS := gocopy gopaste

.PHONY: all build install uninstall clean
all: build

build: $(CMDS:%=$(BIN_DIR)/%)

$(BIN_DIR)/%:
	$(GO) build -o $@ ./cmd/$*

install:
	$(GO) install ./...

uninstall:
	-@for cmd in $(CMDS); do \
		rm -f "$$(go env GOPATH)/bin/$$cmd"; \
	done

clean:
	rm -rf $(BIN_DIR)
