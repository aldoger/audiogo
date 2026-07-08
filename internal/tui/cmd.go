package tui

import (
	"time"

	"github.com/aldoger/audiogo/internal/service"
	tea "github.com/charmbracelet/bubbletea"
)

type autoBackMsg struct{}

func autoBackCmd() tea.Cmd {
	return tea.Tick(time.Second*3, func(time.Time) tea.Msg {
		return autoBackMsg{}
	})
}

type SongFinishedMsg struct{}

func waitForSong(ap *service.AudioPlayer) tea.Cmd {
	return func() tea.Msg {
		<-ap.Done()
		return SongFinishedMsg{}
	}
}
