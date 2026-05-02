BINARY := api
PKG := ./cmd/api
BIN_DIR := bin

.PHONY: run build test tidy clean

run:
	go run $(PKG)

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY) $(PKG)

test:
	go test ./...

tidy:
	go mod tidy

clean:
	rm -rf $(BIN_DIR)
