package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func dirExist(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func listMusic(path string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// Supported audio extensions
	audioExt := map[string]bool{
		".mp3":  true,
		".wav":  true,
		".flac": true,
	}

	var musicFiles []os.DirEntry

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(file.Name()))
		if audioExt[ext] {
			musicFiles = append(musicFiles, file)
		}
	}

	if len(musicFiles) < 1 {
		return nil, errors.New("empty directory, no audio files found")
	}

	return musicFiles, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func main() {

	Home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	MusicDir := Home + "/Music"

	result, err := dirExist(MusicDir)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	if !result {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	musicFiles, err := listMusic(MusicDir)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	p := tea.NewProgram(initialModel(musicFiles))
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error: %s", err.Error())
		os.Exit(1)
	}

}
