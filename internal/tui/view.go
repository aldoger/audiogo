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
	viewSplash viewMode = iota
	viewMenu
	viewAddMusic
	viewMusicList
)

func (m model) View() string {
	header := lipgloss.PlaceHorizontal(m.width, lipgloss.Left, titleHeader())

	var content string
	switch m.mode {
	case viewMenu:
		content = m.menuView()
	case viewAddMusic:
		content = m.addMusicView()
	case viewMusicList:
		content = m.listMusicView()
	}

	body := docStyle.Render(content)
	return header + "\n\n" + body
}

func (m model) menuView() string {
	s := titleStyle.Render("Choose Option")
	s += "\n\n"

	for i, option := range m.options {
		if i == m.cursor {
			s += selectedStyle.Render("> " + option)
		} else {
			s += normalStyle.Render("  " + option)
		}
		s += "\n"
	}

	if m.message != nil {
		s += "\n"
		s += m.message.Render()
		s += "\n"
	}

	s += "\n"
	s += helpStyle.Render("↑/↓ Move • Enter Select • q Quit")

	return s
}

func (m model) addMusicView() string {
	s := titleStyle.Render("Add Music")
	s += "\n\n"

	for i, file := range m.choices {
		name := file

		if i == m.cursor {
			s += selectedStyle.Render("> " + name)
		} else {
			s += normalStyle.Render("  " + name)
		}
		s += "\n"
	}

	if m.message != nil {
		s += "\n"
		s += m.message.Render()
		s += "\n"
	}

	s += "\n"
	s += helpStyle.Render("Enter Add • b Back")

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
	s += helpStyle.Render("Returning to menu...")

	return s
}
