package tui

import (
	"time"

	"github.com/aldoger/audiogo/internal/service"
	tea "github.com/charmbracelet/bubbletea"
)

type Music struct {
	Title    string
	Duration time.Duration
}

type model struct {
	width  int
	height int

	mode         viewMode
	options      []string
	playOptions  []string
	currentMusic Music
	currentTime  time.Duration
	choices      *[]service.MusicFile
	menuCursor   int
	mainCursor   int
	message      ModelMessage
	player       *service.AudioPlayer
	musicQueue   service.MusicQueue
}

func InitialModel(musicFiles *[]service.MusicFile) model {
	return model{
		mode:        viewMenu,
		options:     []string{"add", "play", "list", "search"},
		playOptions: []string{"pause", "resume", "next"},
		choices:     musicFiles,
		player:      service.NewAudioPlayer(),
		musicQueue:  service.NewMusicQueue(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
