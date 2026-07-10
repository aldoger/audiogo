package tui

import (
	"github.com/aldoger/audiogo/internal/service"
	tea "github.com/charmbracelet/bubbletea"
)

type SongFinishedMsg struct{}

func waitForSong(ap *service.AudioPlayer) tea.Cmd {
	return func() tea.Msg {
		<-ap.Done()
		return SongFinishedMsg{}
	}
}
