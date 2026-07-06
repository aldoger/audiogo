package tui

import (
	"os"

	"github.com/aldoger/audiogo/internal/config"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	width  int
	height int

	mode       viewMode
	options    []string
	choices    []os.DirEntry
	cursor     int
	selected   map[int]struct{}
	musicQueue config.MusicQueue
}

func InitialModel(musicFiles []os.DirEntry) model {
	return model{
		mode:       viewMenu,
		options:    []string{"add", "play", "list", "stop", "resume", "back", "next"},
		choices:    musicFiles,
		selected:   make(map[int]struct{}),
		musicQueue: config.NewMusicQueue(),
	}
}

func (model) Init() tea.Cmd {
	return nil
}
