package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rameez471/game/pkg"
)

func main() {
	p := tea.NewProgram(pkg.InitialModel())

	if p.Start() != nil {
		fmt.Println("Could not trigger program")
		os.Exit(1)
	}
}
