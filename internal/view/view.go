package view

import (
	"fmt"
	"strings"
)

func (m GameModel) View() string {
	var sb strings.Builder
	fmt.Fprint(&sb, m.boardView())
	fmt.Fprint(&sb, m.statusBarView())

	return sb.String()
}

// drawing characters
const (
	tl = "┌"
	tr = "┐"
	h  = "─"
	v  = "│"
	bl = "└"
	br = "┘"
)

func (m GameModel) boardView() string {
	b := m.currentBoard
	// global string builder for the whole board
	var sb strings.Builder

	// top border, bottom border
	var tb strings.Builder
	var bb strings.Builder

	// create the top and bottom borders
	fmt.Fprint(&tb, tl)
	fmt.Fprint(&bb, bl)
	for x := 0; x < b.Width; x += 1 {
		fmt.Fprintf(&tb, "%s%s", h, h)
		fmt.Fprintf(&bb, "%s%s", h, h)
	}
	fmt.Fprintf(&tb, "%s%s", h, tr)
	fmt.Fprintf(&bb, "%s%s", h, br)

	// draw the top border to the global board builder
	fmt.Fprintf(&sb, "%s\n", tb.String())

	// iterate the board
	for y := 0; y < b.Height; y += 1 {
		// draw a vertical border at the start of each new row
		fmt.Fprintf(&sb, "%s ", v)
		for x := 0; x < b.Width; x += 1 {
			if cell := b.GetCellAt(x, y); cell != nil {
				fmt.Fprintf(&sb, "%s ", cell.State)
			}
		}

		// draw a vertical border at the end of each new row
		fmt.Fprintf(&sb, "%s\n", v)
	}

	// draw the bottom border to the global board builder
	fmt.Fprintf(&sb, "%s\n", bb.String())

	return sb.String()
}

func (m GameModel) statusBarView() string {
	var sb strings.Builder
	fmt.Fprintf(
		&sb,
		"h: %d, w: %d, Hz: %.02f, paused: %t, mode: %s",
		m.currentBoard.Height,
		m.currentBoard.Width,
		1.00/(float64(m.tickRate)/1000.00),
		m.paused,
		m.mode,
	)
	return sb.String()
}
