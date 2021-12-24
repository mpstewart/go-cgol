package view

import (
	"fmt"
	"strings"
)

func (m GameModel) View() string {
	var sb strings.Builder
	fmt.Fprint(&sb, m.currentBoard.DrawBoard())
	fmt.Fprintf(&sb, "h: %d, w: %d, Hz: %.02f, paused: %t", m.currentBoard.Height, m.currentBoard.Width, 1.00/(float64(m.tickRate)/1000.00), m.paused)

	return sb.String()
}
