package tui

import (
	"time"

	"github.com/aldoger/audiogo/internal/service"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	width  int
	height int

	mode          viewMode
	options       []string
	playOptions   []string
	currentMusic  string
	musicDuration string
	totalDuration time.Duration
	currentTime   time.Duration
	choices       *[]service.MusicFile
	menuCursor    int
	mainCursor    int
	message       ModelMessage
	player        *service.AudioPlayer
	musicQueue    service.MusicQueue
	playlists     *service.PlaylistData
}

func InitialModel(musicFiles *[]service.MusicFile, playlists *service.PlaylistData) model {
	return model{
		mode:        viewMenu,
		options:     []string{"add", "play", "list"},
		playOptions: []string{"pause", "resume", "next"},
		choices:     musicFiles,
		player:      service.InitAudioPlayer(),
		musicQueue:  service.InitMusicQueue(),
		playlists:   playlists,
	}
}

func (m model) Init() tea.Cmd {
	return tickCmd() // start the tick loop once
}
