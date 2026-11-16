package tui

import (
	"strconv"
	"time"

	tea "charm.land/bubbletea/v2"
	core "github.com/Rotwo/instead/internal/core"
)

type Model struct {
	tasks []core.Task
}

func (m Model) Init() tea.Cmd { return nil /* nil means no I/O */ }

func InitialModel() Model {
	return Model{
		tasks: []core.Task{
			{
				Title:    "Task 1",
				Duration: time.Minute * 5,
			},
			{
				Title:    "Task 2",
				Duration: time.Minute * 10,
			},
		},
	}
}

// from bubbletea examples, tweaked for a simple behaviour
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

// from bubbletea examples, tweaked for a simple representation
func (m Model) View() tea.View {
	// The header
	s := "instead\n\n"

	// Content
	for i := 0; i < len(m.tasks); i++ {
		var task core.Task = m.tasks[i]

		s += strconv.Itoa(i) + " - " + task.Title + " - " + task.Duration.String() + "\n"
	}

	// The footer
	s += "\n\nPress q to quit.\n"

	// Send the UI for rendering
	return tea.NewView(s)
}
