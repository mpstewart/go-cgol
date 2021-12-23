package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mpstewart/go-cgol/internal/cgol"
	"golang.org/x/term"
)

type gameModel struct {
	Board        *cgol.Board
	boardHistory []*cgol.Board
	paused       bool
	frequency    int64
}

type tickMsg time.Time

func (m gameModel) nextBoard() gameModel {
	b := m.Board
	m.boardHistory = append(m.boardHistory, b)
	m.Board = b.NextBoard()

	return m
}

func (m gameModel) lastBoard() gameModel {
	n := len(m.boardHistory)
	if n > 0 {
		b := m.boardHistory[n-1]
		m.boardHistory = m.boardHistory[:n-1]
		m.Board = b
		return m
	}

	return m
}

func (m gameModel) reset() gameModel {
	b := m.boardHistory[0]
	m.boardHistory = m.boardHistory[:0]
	m.Board = b
	return m
}

func (m gameModel) tick() tea.Cmd {
	d := time.Duration(m.frequency) * time.Millisecond
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m gameModel) Init() tea.Cmd {
	return m.tick()
}

func (m gameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tickMsg:
		if m.paused {
			return m, nil
		} else {
			return m.nextBoard(), m.tick()
		}

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		// history scrubbing
		case "p", " ":
			if m.paused {
				m.paused = false
				m = m.nextBoard()
			} else {
				m.paused = true
			}

			return m, m.tick()
		case "r":
			return m.reset(), nil
		case "h":
			return m.lastBoard(), nil
		case "l":
			return m.nextBoard(), nil

		// tick manipulation
		case "k":
			m.frequency = m.frequency - 25
		case "j":
			m.frequency = m.frequency + 25
		case "1":
			m.frequency = 450
		case "2":
			m.frequency = 400
		case "3":
			m.frequency = 350
		case "4":
			m.frequency = 300
		case "5":
			m.frequency = 250
		}
	}

	return m, nil
}

func (m gameModel) View() string {
	var sb strings.Builder
	fmt.Fprint(&sb, m.Board.DrawBoard())
	fmt.Fprintf(&sb, "h: %d, w: %d, Hz: %.02f, paused: %t", m.Board.Height, m.Board.Width, 1.00/(float64(m.frequency)/1000.00), m.paused)

	return sb.String()
}

func main() {
	fmt.Print("\033[H\033[2J")

	w, h, _ := term.GetSize(1)
	m := initialModel(w/2-2, h-4) // i have no idea why width needs to be divided by 2

	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		panic(err)
	}
}

func initialModel(x, y int) gameModel {
	b := cgol.NewBoard(x, y)
	b.PutGliderAt(0, 0)
	b.PutSquareAt(4, 4)
	b.PutLightweightSpaceShip(70, 8)

	m := gameModel{
		Board:     b,
		paused:    false,
		frequency: 250,
	}

	return m
}
