package tui

import tea "github.com/charmbracelet/bubbletea"

const (
	ADD    = "add"
	LIST   = "list"
	PLAY   = "play"
	STOP   = "stop"
	RESUME = "resume"
	BACK   = "back"
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
					m.mode = viewAddMusic
					m.cursor = 0

				case LIST:
					m.mode = viewMusicList
					return m, autoBackCmd()
				}
			}

		case viewAddMusic:
			switch msg.String() {

			case "b":
				m.mode = viewMenu
				m.cursor = 0

			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			case "down", "j":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}

			case "enter":
				m.musicQueue.Enqueue(m.choices[m.cursor])

			case "q", "ctrl+c":
				return m, tea.Quit
			}
		}
	}

	return m, nil
}
