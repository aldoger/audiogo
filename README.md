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
> AudioGo currently **only supports MP3 files**. Attempting to play other audio formats (such as WAV, FLAC, M4A, AAC, or OGG) may result in playback errors or unsupported format messages.

## Controls

| Key | Action |
|-----|--------|
| ↑ / ↓ | Navigate through the file list |
| Enter | Play selected audio |
| Space | Pause / Resume playback |
| q | Quit the application |

> **Note:** Control keys may vary depending on future releases.

## Installation

### Prerequisites

- Go 1.24 or newer
- GNU Make

### Steps

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

## License

This project is licensed under the MIT License.