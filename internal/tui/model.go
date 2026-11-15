package tui

import (
	tea "charm.land/bubbletea/v2"
)

// needs to be moved to core
type Tasks [10]string

type Model struct {
	tasks Tasks
}

func (m Model) Init() tea.Cmd { return nil /* nil means no I/O */ }

func InitialModel() Model {
	return Model{
		tasks: Tasks{"Task 1", "Task 2"},
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
	s := "What should we buy at the market?\n\n"

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return tea.NewView(s)
}
