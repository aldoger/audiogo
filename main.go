package main

import (
	"log"
	"os"

	"github.com/aldoger/audiogo/internal/service"
	"github.com/aldoger/audiogo/internal/tui"
	"github.com/aldoger/audiogo/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Check & initialize music directory
	musicDir, err := utils.MusicDirExist()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		os.Exit(1)
	}

	// Get music files
	musicFiles, err := utils.ListMusic(musicDir)
	if musicFiles == nil && err != nil {
		log.Fatalf("Error: %s", err.Error())
		os.Exit(1)
	}

	playlist, err := service.LoadPlaylists()
	if err != nil {
		log.Fatalf("Error %s", err.Error())
		os.Exit(1)
	}

	// Start Program
	p := tea.NewProgram(tui.InitialModel(&musicFiles, playlist))
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error: %s", err.Error())
		os.Exit(1)
	}
}
