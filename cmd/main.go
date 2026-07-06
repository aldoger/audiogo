package main

import (
	"log"
	"os"

	"github.com/aldoger/audiogo/internal/tui"
	"github.com/aldoger/audiogo/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	musicDir, err := utils.DirExist()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
		os.Exit(1)
	}

	musicFiles, err := utils.ListMusic(musicDir)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	p := tea.NewProgram(tui.InitialModel(musicFiles))
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error: %s", err.Error())
		os.Exit(1)
	}

}
