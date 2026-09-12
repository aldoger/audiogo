package tui

import (
	"time"

	"github.com/aldoger/audiogo/internal/service"
	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg struct{}

func tickCmd() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return TickMsg{}
	})
}

type SongFinishedMsg struct{}

func waitForSong(ap *service.AudioPlayer) tea.Cmd {
	return func() tea.Msg {
		<-ap.Done()
		return SongFinishedMsg{}
	}
}

type MusicStartedMsg struct {
	Duration time.Duration
	Err      error
}

func playMusic(player *service.AudioPlayer, music string) tea.Cmd {
	return func() tea.Msg {
		duration, err := player.Play(music)

		return MusicStartedMsg{
			Duration: duration,
			Err:      err,
		}
	}
}
