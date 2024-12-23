MAIN_FILE=main.go
BIN_DIR=./bin

.PHONY: build

build:
	go build -ldflags '-linkmode "external" -extldflags "-static"' -o $(BIN_DIR)/main ${MAIN_FILE}

env:
	docker compose up -d

test:
	go test -v -coverprofile=coverage.out ./...