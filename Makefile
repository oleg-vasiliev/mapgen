## Download dependencies and verify
.PHONY: verify
verify:
	go mod download
	go mod verify

## Format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## Build binary
.PHONY: build
build:
	go mod verify
	go build -o=./bin/ ./cmd/mapgen

.PHONY: run
run:
	go run ./cmd/mapgen --window 1024x768 --seed 100x70@123

.PHONY: run-fullscreen
run-fullscreen:
	go run ./cmd/mapgen --fullscreen
