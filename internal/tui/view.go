package tui

import (
	"fmt"
	"strings"

	"github.com/aldoger/audiogo/internal/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
)

func titleHeader() string {
	title := strings.TrimRight(
		figure.NewFigure("AUDIOGO", "larry3d", true).String(),
		"\n",
	)

	return splashStyle.Render(title)
}

type viewMode int

const (
	viewMenu viewMode = iota
	viewAddMusic
	viewMusicList
	viewPlayMusic
)

// --------------------------------------------------
// Main View
// --------------------------------------------------

func (m model) View() string {
	if m.width <= 0 || m.height <= 0 {
		return ""
	}

	appHeight := max(12, m.height/2)

	appWidth := m.width - 2

	headerHeight := 1
	footerHeight := 1

	bodyHeight := appHeight - headerHeight - footerHeight

	sidebarWidth := 24
	contentWidth := appWidth - sidebarWidth

	header := m.headerView()

	sidebar := Box(
		m.menuView(),
		sidebarWidth,
		bodyHeight,
	)

	content := Box(
		m.currentView(
			contentWidth,
			bodyHeight,
		),
		contentWidth,
		bodyHeight,
	)

	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		sidebar,
		content,
	)

	footer := m.footerView()

	app := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
		footer,
	)

	// Center the 50%-height application vertically.
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		app,
	)
}

// --------------------------------------------------
// Layout Helpers
// --------------------------------------------------

func (m model) sidebarWidth() int {
	return 24
}

func (m model) contentWidth() int {
	return m.width - m.sidebarWidth()
}

func (m model) contentHeight() int {
	return max(8, m.height/2-2)
}

// --------------------------------------------------
// Header
// --------------------------------------------------

func (m model) headerView() string {
	status := "● STOPPED"

	if m.player != nil {
		if m.player.IsPaused() {
			status = "● PAUSED"
		} else {
			status = "● PLAYING"
		}
	}

	left := titleStyle.Render("AUDIOGO")
	right := helpStyle.Render(status)

	availableWidth := max(
		1,
		m.width-lipgloss.Width(left),
	)

	rightAligned := lipgloss.PlaceHorizontal(
		availableWidth,
		lipgloss.Right,
		right,
	)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		rightAligned,
	)
}

// --------------------------------------------------
// Sidebar
// --------------------------------------------------

func (m model) menuView() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("Library"))
	b.WriteString("\n\n")

	for i, option := range m.options {
		if i == m.menuCursor {
			b.WriteString(
				selectedStyle.Render("> " + option),
			)
		} else {
			b.WriteString(
				normalStyle.Render("  " + option),
			)
		}

		b.WriteString("\n")
	}

	return b.String()
}

// --------------------------------------------------
// Footer
// --------------------------------------------------

func (m model) footerView() string {
	help := lipgloss.JoinHorizontal(
		lipgloss.Center,

		helpStyle.Render("↑/↓ Navigate"),
		"   ",
		helpStyle.Render("Enter Select"),
		"   ",
		helpStyle.Render("Space Play/Pause"),
		"   ",
		helpStyle.Render("b Back"),
		"   ",
		helpStyle.Render("q Quit"),
	)

	return help
}

// --------------------------------------------------
// Current View
// --------------------------------------------------

func (m model) currentView(width, height int) string {
	switch m.mode {

	case viewMenu:
		return m.homeView(width, height)

	case viewAddMusic:
		return m.addMusicView(width, height)

	case viewMusicList:
		return m.listMusicView(width, height)

	case viewPlayMusic:
		return m.playMusicView(width, height)

	default:
		return ""
	}
}

// --------------------------------------------------
// Home
// --------------------------------------------------

func (m model) homeView(width, height int) string {
	var b strings.Builder

	title := titleStyle.Render("Welcome to AUDIOGO")

	b.WriteString(title)
	b.WriteString("\n\n")

	description := helpStyle.Render(
		"Your music player, running entirely in the terminal.",
	)

	b.WriteString(description)
	b.WriteString("\n\n\n")

	// ----------------------------------------------
	// Statistics
	// ----------------------------------------------

	musicCount := 0

	if m.choices != nil {
		musicCount = len(*m.choices)
	}

	queue := m.musicQueue.ListMusicInQueue()
	queueCount := len(queue)

	statsWidth := max(20, (width-8)/2)

	musicBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Width(statsWidth).
		Render(
			titleStyle.Render("MUSIC") +
				"\n\n" +
				fmt.Sprintf("%d tracks", musicCount),
		)

	queueBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Width(statsWidth).
		Render(
			titleStyle.Render("QUEUE") +
				"\n\n" +
				fmt.Sprintf("%d tracks", queueCount),
		)

	stats := lipgloss.JoinHorizontal(
		lipgloss.Top,
		musicBox,
		"  ",
		queueBox,
	)

	b.WriteString(stats)
	b.WriteString("\n\n")

	// ----------------------------------------------
	// Current Music
	// ----------------------------------------------

	b.WriteString(titleStyle.Render("Current Track"))
	b.WriteString("\n\n")

	if m.currentMusic.Title != "" {
		b.WriteString(
			selectedStyle.Render(
				"♪ " + m.currentMusic.Title,
			),
		)
	} else {
		b.WriteString(
			helpStyle.Render("No music currently playing."),
		)
	}

	return lipgloss.Place(
		width-2,
		height-2,
		lipgloss.Left,
		lipgloss.Top,
		b.String(),
	)
}

// --------------------------------------------------
// Add Music
// --------------------------------------------------

func (m model) addMusicView(width, height int) string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("Add Music"))
	b.WriteString("\n\n")

	b.WriteString(
		helpStyle.Render(
			"Select music files to add to your library.",
		),
	)

	b.WriteString("\n\n")

	if m.choices == nil || len(*m.choices) == 0 {
		b.WriteString(
			infoStyle.Render("No music files found!"),
		)

		if m.message != nil {
			b.WriteString("\n\n")
			b.WriteString(m.message.Render())
		}

		return lipgloss.Place(
			width-2,
			height-2,
			lipgloss.Left,
			lipgloss.Top,
			b.String(),
		)
	}

	// ----------------------------------------------
	// File list
	// ----------------------------------------------

	listWidth := max(10, width-6)
	listHeight := max(5, height-10)

	var list strings.Builder

	for i, file := range *m.choices {

		name := file.Title

		// Prevent very long filenames from destroying
		// the layout.
		maxNameWidth := max(5, listWidth-6)

		if lipgloss.Width(name) > maxNameWidth {
			name = truncateString(name, maxNameWidth)
		}

		if i == m.mainCursor {
			list.WriteString(
				selectedStyle.Render("> " + name),
			)
		} else {
			list.WriteString(
				normalStyle.Render("  " + name),
			)
		}

		if i != len(*m.choices)-1 {
			list.WriteString("\n")
		}
	}

	fileList := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Width(listWidth).
		Height(listHeight).
		Render(list.String())

	b.WriteString(fileList)
	b.WriteString("\n\n")

	// ----------------------------------------------
	// Message
	// ----------------------------------------------

	if m.message != nil {
		b.WriteString(m.message.Render())
	}

	return lipgloss.Place(
		width-2,
		height-2,
		lipgloss.Left,
		lipgloss.Top,
		b.String(),
	)
}

// --------------------------------------------------
// Music Queue
// --------------------------------------------------

func (m model) listMusicView(width, height int) string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("Music Queue"))
	b.WriteString("\n\n")

	queue := m.musicQueue.ListMusicInQueue()

	if len(queue) == 0 {
		b.WriteString(
			helpStyle.Render("(queue is empty)"),
		)

		return lipgloss.Place(
			width-2,
			height-2,
			lipgloss.Left,
			lipgloss.Top,
			b.String(),
		)
	}

	b.WriteString(
		helpStyle.Render(
			fmt.Sprintf("%d tracks in queue", len(queue)),
		),
	)

	b.WriteString("\n\n")

	// ----------------------------------------------
	// Queue list
	// ----------------------------------------------

	var list strings.Builder

	for i, song := range queue {

		prefix := fmt.Sprintf("%02d  ", i+1)

		// Highlight current track.
		if i == 0 {
			list.WriteString(
				selectedStyle.Render(
					prefix + "♪ " + song,
				),
			)
		} else {
			list.WriteString(
				normalStyle.Render(
					prefix + song,
				),
			)
		}

		if i != len(queue)-1 {
			list.WriteString("\n")
		}
	}

	queueBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Width(max(10, width-6)).
		Height(max(5, height-8)).
		Render(list.String())

	b.WriteString(queueBox)

	return lipgloss.Place(
		width-2,
		height-2,
		lipgloss.Left,
		lipgloss.Top,
		b.String(),
	)
}

// --------------------------------------------------
// Now Playing
// --------------------------------------------------

func (m model) playMusicView(width, height int) string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("NOW PLAYING"))
	b.WriteString("\n\n")

	music := m.currentMusic

	title := music.Title
	if title == "" {
		title = "No music selected"
	}

	b.WriteString(
		selectedStyle.Render("♪ " + title),
	)

	b.WriteString("\n\n")

	progress := 0.42

	barWidth := max(20, width-20)
	filled := int(progress * float64(barWidth))

	bar := strings.Repeat("━", filled)

	if filled < barWidth {
		bar += "╺"
	}

	bar += strings.Repeat(
		"━",
		max(0, barWidth-filled-1),
	)

	current := utils.FormatDuration(m.currentTime)
	duration := utils.FormatDuration(music.Duration)

	b.WriteString(
		helpStyle.Render(
			fmt.Sprintf("%s / %s", current, duration),
		),
	)

	b.WriteString("\n")

	b.WriteString(
		lipgloss.PlaceHorizontal(
			width-4,
			lipgloss.Center,
			bar,
		),
	)

	b.WriteString("\n\n")

	playback := "⏸ Pause"

	if m.player != nil && m.player.IsPaused() {
		playback = "▶ Resume"
	}

	controls := lipgloss.JoinHorizontal(
		lipgloss.Center,
		helpStyle.Render("⏮ Previous"),
		"   ",
		selectedStyle.Render(playback),
		"   ",
		helpStyle.Render("⏭ Next"),
	)

	b.WriteString(
		lipgloss.PlaceHorizontal(
			width-4,
			lipgloss.Center,
			controls,
		),
	)

	return b.String()
}

// --------------------------------------------------
// Small Terminal View
// --------------------------------------------------

func (m model) smallView() string {
	var b strings.Builder

	b.WriteString(
		titleStyle.Render("AUDIOGO"),
	)

	b.WriteString("\n\n")

	switch m.mode {

	case viewMenu:
		b.WriteString(m.menuView())

	case viewAddMusic:
		b.WriteString(m.addMusicView(m.width, m.height))

	case viewMusicList:
		b.WriteString(m.listMusicView(m.width, m.height))

	case viewPlayMusic:
		b.WriteString(m.playMusicView(m.width, m.height))
	}

	b.WriteString("\n\n")
	b.WriteString(
		helpStyle.Render("q Quit"),
	)

	return b.String()
}

// --------------------------------------------------
// Utilities
// --------------------------------------------------

func truncateString(s string, maxWidth int) string {
	if maxWidth <= 0 {
		return ""
	}

	if lipgloss.Width(s) <= maxWidth {
		return s
	}

	if maxWidth <= 3 {
		return s[:maxWidth]
	}

	return s[:maxWidth-3] + "..."
}
