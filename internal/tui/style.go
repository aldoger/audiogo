package tui

import "github.com/charmbracelet/lipgloss"

var (
	docStyle = lipgloss.NewStyle().
			Margin(1, 2)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7D56F4"))

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true)

	normalStyle = lipgloss.NewStyle()

	warningStyle = lipgloss.NewStyle().
			Bold(true).Foreground(lipgloss.Color("#b50a04"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262"))
)
