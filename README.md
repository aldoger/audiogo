# AudioGo

AudioGo is a simple terminal-based audio player built with Go. It provides a clean Text User Interface (TUI) for browsing and playing audio files directly from your terminal. The application is designed to be lightweight, fast, and easy to use without leaving the command line.

## Features

- 🎵 Browse audio files in a terminal interface
- ▶️ Play audio files
- ⏸️ Pause and resume playback
- ⏭️ Navigate through your music library
- ⌨️ Keyboard-driven controls
- ⚡ Lightweight and responsive

> **⚠️ Warning**
>
> AudioGo currently **only supports MP3, WAV, OGG, FLAC files**. Attempting to play other audio formats may result in playback errors or unsupported format messages.

## Controls

| Key | Action |
|-----|--------|
| ↑ / ↓ | Navigate through the file list |
| Enter | Play selected audio |
| Space | Pause / Resume playback |
| q | Quit the application |

> **Note:** Control keys may vary depending on future releases.

## Installation

### Install

To build and install AudioGo:

```bash
make install
```

The executable will be installed to:

```text
~/.local/bin/audiogo
```

Make sure `~/.local/bin` is included in your `PATH`.

### Uninstall

To remove AudioGo:

```bash
make uninstall
```


### Prerequisites

- Go 1.24 or newer
- GNU Make

### Development Steps

1. Clone the repository.

```bash
git clone https://github.com/aldoger/audiogo.git
cd audiogo
```

2. Download dependencies.

```bash
go mod download
```

3. Run the application.

```bash
make run
```

## Project Structure

```
audiogo/
├── cmd/          # Application entry point
├── internal/     # Internal packages
├── assets/       # Audio or static assets (optional)
├── Makefile
├── go.mod
└── README.md
```

## Built With

- Go
- Bubble Tea
- Bubbles
- Lip Gloss
- Beep

## License

This project is licensed under the MIT License.
