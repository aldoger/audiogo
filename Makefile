.PHONY: check-go download-deps build install uninstall

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

install: build
	mkdir -p ~/.local/bin
	cp bin/audiogo ~/.local/bin/audiogo

uninstall:
	rm -f ~/.local/bin/audiogo
