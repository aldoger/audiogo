.PHONY: check-go build

check-go:
	ifeq ($(OS),Windows_NT)
		@where go >nul 2>&1 || (echo Go is not installed. && exit 1)
	else
		@command -v go >/dev/null 2>&1 || (echo "Go is not installed." && exit 1)
	endif

download-deps: check-go
	go mod download

build: check-go
	go build -o bin/audiogo
