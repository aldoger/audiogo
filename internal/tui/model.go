package tui

import (
	"github.com/aldoger/audiogo/internal/config"
	"github.com/aldoger/audiogo/internal/service"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	width  int
	height int

	mode       viewMode
	options    []string
	choices    []string
	cursor     int
	message    ModelMessage
	player     service.AudioPlayer
	selected   map[int]struct{}
	musicQueue config.MusicQueue
}

func InitialModel(musicFiles []string) model {
	return model{
		mode:       viewMenu,
		options:    []string{"add", "play", "list", "pause", "resume", "next"},
		choices:    musicFiles,
		player:     service.NewAudioPlayer(),
		selected:   make(map[int]struct{}),
		musicQueue: config.NewMusicQueue(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
