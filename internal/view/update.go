package view

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.tickRate = m.tickRate - 25
		case "j":
			m.tickRate = m.tickRate + 25
		case "1":
			m.tickRate = 450
		case "2":
			m.tickRate = 400
		case "3":
			m.tickRate = 350
		case "4":
			m.tickRate = 300
		case "5":
			m.tickRate = 250
		}
	}

	return m, nil
}

func (m GameModel) tick() tea.Cmd {
	d := time.Duration(m.tickRate) * time.Millisecond
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m GameModel) nextBoard() GameModel {
	b := m.currentBoard
	m.boardHistory = append(m.boardHistory, b)
	m.currentBoard = b.NextBoard()

	return m
}

func (m GameModel) lastBoard() GameModel {
	n := len(m.boardHistory)
	if n > 0 {
		b := m.boardHistory[n-1]
		m.boardHistory = m.boardHistory[:n-1]
		m.currentBoard = b
		return m
	}

	return m
}

func (m GameModel) reset() GameModel {
	b := m.boardHistory[0]
	m.boardHistory = m.boardHistory[:0]
	m.currentBoard = b
	return m
}
