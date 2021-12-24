package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mpstewart/go-cgol/internal/view"
	"golang.org/x/term"
)

func main() {
	// clear the screen
	fmt.Print("\033[H\033[2J")

	w, h, _ := term.GetSize(1)
	m := view.InitialModel(w/2-2, h-4)

	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		panic(err)
	}
}
