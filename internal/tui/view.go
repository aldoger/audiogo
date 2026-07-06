package tui

import "fmt"

type viewMode int

const (
	viewMenu viewMode = iota
	viewAddMusic
	viewMusicList
)

func (m model) View() string {
	switch m.mode {

	case viewMenu:
		return m.menuView()

	case viewAddMusic:
		return m.addMusicView()

	case viewMusicList:
		return m.listMusicView()
	}

	return ""
}

func (m model) menuView() string {
	s := "Choose option:\n\n"

	for i, option := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, option)
	}

	s += "\nEnter to select • q to quit\n"
	return s
}

func (m model) addMusicView() string {
	s := "Add music to queue:\n\n"

	for i, file := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, file.Name())
	}

	s += "\nPress Enter: add music\nPress b: back\n"
	return s
}

func (m model) listMusicView() string {
	s := "List music in queue:\n\n"

	musicList := m.musicQueue.ListMusicInQueue()

	if len(musicList) == 0 {
		s += "(queue is empty)\n"
		return s
	}

	for i, name := range musicList {
		s += fmt.Sprintf("%d. %s\n", i+1, name)
	}

	return s
}
