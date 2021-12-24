package view

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mpstewart/go-cgol/internal/cgol"
)

type mode string

const (
	modeNormal mode = "normal"
	modeEdit   mode = "edit"
)

type GameModel struct {
	currentBoard *cgol.Board
	boardHistory []*cgol.Board
	paused       bool
	tickRate     int64
	mode         mode
	cursorPos    []int
}

func InitialModel(x, y int) GameModel {
	cb := cgol.NewBoard(x, y)
	bh := make([]*cgol.Board, 0)
	bh = append(bh, cb)
	p := false
	tr := int64(250)
	mode := modeEdit
	cp := []int{0, 0}

	m := GameModel{
		currentBoard: cb,
		boardHistory: bh,
		paused:       p,
		tickRate:     tr,
		mode:         mode,
		cursorPos:    cp,
	}

	return m
}

func (m GameModel) Init() tea.Cmd {
	return m.tick()
}
