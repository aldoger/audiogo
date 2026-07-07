package tui

import "github.com/charmbracelet/lipgloss"

type ModelMessage interface {
	Render() string
}

type WarningMessage struct {
	Text string
}

func (m WarningMessage) Render() string {
	return warningStyle.Render(m.Text)
}

type HelpMessage struct {
	Text string
}

func (m HelpMessage) Render() string {
	return helpStyle.Render(m.Text)
}

type SelectedMessage struct {
	Text string
}

func (m SelectedMessage) Render() string {
	return selectedStyle.Render(m.Text)
}

type InfoMessage struct {
	Text string
}

func (m InfoMessage) Render() string {
	return infoStyle.Render(m.Text)
}

var (
	docStyle = lipgloss.NewStyle().
			Margin(1, 2)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4"))

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#0b8704")).
			Bold(true)

	normalStyle = lipgloss.NewStyle()

	warningStyle = lipgloss.NewStyle().
			Bold(true).Foreground(lipgloss.Color("#b50a04"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262"))
)
