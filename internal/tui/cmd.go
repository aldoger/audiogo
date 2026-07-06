package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type autoBackMsg struct{}

func autoBackCmd() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return autoBackMsg{}
	})
}
