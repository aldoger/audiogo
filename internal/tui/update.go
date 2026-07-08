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

	case autoBackMsg:
		m.mode = viewMenu
		return m, nil

	case SongFinishedMsg:
		next := m.musicQueue.Dequeue()
		if next == "" {
			m.message = WarningMessage{Text: "No more music in queue"}
			return m, nil
		}
		m.message = InfoMessage{Text: "Playing Next Music..."}
		if err := m.player.Play(next); err != nil {
			m.message = WarningMessage{Text: err.Error()}
			return m, nil
		}
		return m, waitForSong(m.player)

	case tea.KeyMsg:
		switch m.mode {

		case viewMenu:
			switch msg.String() {

			case "q", "ctrl+c":
				return m, tea.Quit

			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			case "down", "j":
				if m.cursor < len(m.options)-1 {
					m.cursor++
				}

			case "enter":
				switch m.options[m.cursor] {
				case ADD:
					// check for another music if added in directory
					dirPath, _ := utils.DirExist()
					checkMusic, _ := utils.ListMusic(dirPath)
					m.choices = &checkMusic

					m.mode = viewAddMusic
					m.cursor = 0

				case LIST:
					m.mode = viewMusicList
					return m, autoBackCmd()

				case PLAY:
					music := m.musicQueue.Dequeue()
					if music == "" {
						m.message = WarningMessage{Text: "No music in queue yet!!"}
						return m, nil
					}
					m.message = InfoMessage{Text: "Playing..."}
					if err := m.player.Play(music); err != nil {
						m.message = WarningMessage{Text: err.Error()}
						return m, nil
					}
					return m, waitForSong(m.player)

				case PAUSE:
					m.message = InfoMessage{Text: "Music pause!"}
					m.player.Pause()

				case RESUME:
					m.message = InfoMessage{Text: "Resume..."}
					m.player.Resume()

				case NEXT:
					next := m.musicQueue.Dequeue()
					if next == "" {
						m.message = WarningMessage{Text: "No more music in queue"}
						return m, nil
					}
					m.message = InfoMessage{Text: "Playing Next Music..."}
					if err := m.player.Play(next); err != nil {
						m.message = WarningMessage{Text: err.Error()}
						return m, nil
					}
					return m, waitForSong(m.player)
				}
			}

		case viewAddMusic:
			switch msg.String() {

			case "b":
				m.mode = viewMenu
				m.cursor = 0
				m.message = nil

			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			case "down", "j":
				if m.cursor < len(*m.choices)-1 {
					m.cursor++
				}

			case "enter":
				m.musicQueue.Enqueue((*m.choices)[m.cursor])
				msg := fmt.Sprintf("%s added to queue!", (*m.choices)[m.cursor])
				m.message = SelectedMessage{Text: msg}

			case "q", "ctrl+c":
				return m, tea.Quit
			}
		}
	}

	return m, nil
}
