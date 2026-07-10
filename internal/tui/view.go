package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
)

func titleHeader() string {
	title := strings.TrimRight(figure.NewFigure("AUDIOGO", "larry3d", true).String(), "\n")
	return splashStyle.Render(title)
}

type viewMode int

const (
	viewMenu viewMode = iota
	viewAddMusic
	viewMusicList
)

func (m model) View() string {
	bodyHeight := max(20, m.height-6)

	menuWidth := 30
	rightWidth := m.width - menuWidth

	mainHeight := bodyHeight - 8
	helpHeight := 8

	menu := Box(
		m.menuView(),
		menuWidth,
		bodyHeight,
	)

	main := Box(
		m.currentView(rightWidth, mainHeight),
		rightWidth,
		mainHeight,
	)

	help := Box(
		m.helpView(),
		rightWidth,
		helpHeight,
	)

	right := lipgloss.JoinVertical(
		lipgloss.Left,
		main,
		help,
	)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		menu,
		right,
	)
}

func (m model) currentView(width, height int) string {
	switch m.mode {
	case viewMenu:
		return lipgloss.Place(
			width-2,
			height-2,
			lipgloss.Center,
			lipgloss.Center,
			titleHeader(),
		)

	case viewAddMusic:
		return lipgloss.Place(
			width-2,
			height-2,
			lipgloss.Center,
			lipgloss.Center,
			m.addMusicView(),
		)

	case viewMusicList:
		return m.listMusicView()

	default:
		return ""
	}
}

func (m model) menuView() string {
	s := titleStyle.Render("Choose Operation")
	s += "\n\n"

	for i, option := range m.options {
		if i == m.menuCursor {
			s += selectedStyle.Render("> " + option)
		} else {
			s += normalStyle.Render("  " + option)
		}
		s += "\n"
	}

	s += "\n"
	return s
}

func (m model) helpView() string {
	var s string
	s += helpStyle.Render("↑/↓ Move")
	s += "\n"
	s += helpStyle.Render("Enter Select")
	s += "\n"
	s += helpStyle.Render("b Back")
	s += "\n"
	s += helpStyle.Render("q Quit")

	return s
}

func (m model) addMusicView() string {
	s := titleStyle.Render("Add Music")
	s += "\n\n"

	if len(*m.choices) != 0 {
		for i, file := range *m.choices {
			name := file.Title

			if i == m.mainCursor {
				s += selectedStyle.Render("> " + name)
			} else {
				s += normalStyle.Render("  " + name)
			}
			s += "\n"
		}
	} else {
		s += infoStyle.Render("No music files found!")
		s += "\n"
	}

	if m.message != nil {
		s += "\n"
		s += m.message.Render()
		s += "\n"
	}

	s += "\n"
	return s
}

func (m model) listMusicView() string {
	s := titleStyle.Render("Music Queue")
	s += "\n\n"

	queue := m.musicQueue.ListMusicInQueue()

	if len(queue) == 0 {
		s += helpStyle.Render("(queue is empty)")
		return s
	}

	for i, song := range queue {
		s += normalStyle.Render(fmt.Sprintf("%2d. %s", i+1, song))
		s += "\n"
	}

	s += "\n"
	return s
}
