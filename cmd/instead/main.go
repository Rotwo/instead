package main

import (
	"log"

	tea "charm.land/bubbletea/v2"
	tui "github.com/Rotwo/instead/internal/tui"
)

func main() {
	if _, err := tea.NewProgram(tui.InitialModel()).Run(); err != nil {
		log.Fatal(err)
	}
}
