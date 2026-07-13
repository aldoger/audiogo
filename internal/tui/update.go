package tui

import (
	"fmt"

	"github.com/aldoger/audiogo/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	ADD    = "add"
	LIST   = "list"
	SEARCH = "search"
	PLAY   = "play"
	PAUSE  = "pause"
	RESUME = "resume"
	NEXT   = "next"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case SongFinishedMsg:
		next := m.musicQueue.Dequeue()
		if next == "" {
			m.message = WarningMessage{Text: "No more music in queue"}
			return m, nil
		}
		duration, err := m.player.Play(next)
		if err != nil {
			m.message = WarningMessage{Text: err.Error()}
			return m, nil
		}
		m.currentMusic = Music{Duration: duration, Title: next}
		return m, waitForSong(m.player)

	case tea.KeyMsg:
		switch m.mode {

		case viewMenu:
			switch msg.String() {

			case "q", "ctrl+c":
				return m, tea.Quit

			case "up", "k":
				if m.menuCursor > 0 {
					m.menuCursor--
				}

			case "down", "j":
				if m.menuCursor < len(m.options)-1 {
					m.menuCursor++
				}

			case "enter":
				switch m.options[m.menuCursor] {
				case ADD:
					// check for another music if added in directory
					dirPath, _ := utils.DirExist()
					checkMusic, _ := utils.ListMusic(dirPath)
					m.choices = &checkMusic

					m.mode = viewAddMusic

				case LIST:
					m.mode = viewMusicList
					return m, nil

				case PLAY:
					music := m.musicQueue.Dequeue()
					if music == "" {
						m.message = WarningMessage{Text: "No music in queue yet!!"}
						return m, nil
					}
					m.mode = viewPlayMusic
					duration, err := m.player.Play(music)
					if err != nil {
						m.message = WarningMessage{Text: err.Error()}
						return m, nil
					}
					m.currentMusic = Music{Duration: duration, Title: music}
					return m, waitForSong(m.player)
				}
			}

		case viewMusicList:
			switch msg.String() {
			case "b":
				m.mode = viewMenu
				m.menuCursor = 0
				m.message = nil
			}

		case viewAddMusic:
			switch msg.String() {

			case "b":
				m.mode = viewMenu
				m.menuCursor = 0
				m.message = nil

			case "up", "k":
				if m.mainCursor > 0 {
					m.mainCursor--
				}

			case "down", "j":
				if m.mainCursor < len(*m.choices)-1 {
					m.mainCursor++
				}

			case "enter":
				m.musicQueue.Enqueue((*m.choices)[m.mainCursor].Path)
				msg := fmt.Sprintf("%s added to queue!", (*m.choices)[m.mainCursor].Title)
				m.message = SelectedMessage{Text: msg}

			case "q", "ctrl+c":
				return m, tea.Quit
			}

		case viewPlayMusic:
			switch msg.String() {

			case "b":
				m.mode = viewMenu
				m.menuCursor = 0
				m.message = nil

			case "p":
				m.message = InfoMessage{Text: "Music pause!"}
				m.player.Pause()

			case "m":
				m.message = InfoMessage{Text: "Resume..."}
				m.player.Resume()

			case "n":
				next := m.musicQueue.Dequeue()
				if next == "" {
					m.message = WarningMessage{Text: "No more music in queue"}
					return m, nil
				}
				duration, err := m.player.Play(next)
				if err != nil {
					m.message = WarningMessage{Text: err.Error()}
					return m, nil
				}
				m.currentMusic = Music{Duration: duration, Title: next}
				return m, waitForSong(m.player)
			}
		}
	}

	return m, nil
}
