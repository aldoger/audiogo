package tui

import (
	"github.com/aldoger/audiogo/internal/service"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	width  int
	height int

	mode        viewMode
	options     []string
	playOptions []string
	choices     *[]service.MusicFile
	menuCursor  int
	mainCursor  int
	message     ModelMessage
	player      *service.AudioPlayer
	selected    map[int]struct{}
	musicQueue  service.MusicQueue
}

func InitialModel(musicFiles *[]service.MusicFile) model {
	return model{
		mode:        viewMenu,
		options:     []string{"add", "play", "list", "search"},
		playOptions: []string{"pause", "resume", "next"},
		choices:     musicFiles,
		player:      service.NewAudioPlayer(),
		selected:    make(map[int]struct{}),
		musicQueue:  service.NewMusicQueue(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
