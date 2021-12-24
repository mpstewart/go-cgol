package view

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	switch m.mode {
	case modeNormal:
		return m.updateNormalMode(msg)
	case modeEdit:
		return m.updateEditMode(msg)
	}

	return m, nil
}

func (m GameModel) updateNormalMode(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		switch m.paused {
		case true:
			return m, nil
		case false:
			return m.nextBoard(), m.tick()
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "i":
			m.mode = modeEdit
			return m, nil
		}
	}

	return m, nil
}

func (m GameModel) updateEditMode(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.mode = modeNormal
			return m, nil
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
