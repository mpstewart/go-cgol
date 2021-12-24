package cgol

type CellState int

const (
	CellStateDead CellState = iota
	CellStateAlive
)

type Cell struct {
	State CellState
}

func (c *Cell) VivifyCell() {
	c.State = CellStateAlive
}

func (c *Cell) KillCell() {
	c.State = CellStateDead
}

func (c *Cell) ToggleCell() {
	switch c.State {
	case CellStateAlive:
		c.KillCell()
	case CellStateDead:
		c.VivifyCell()
	}
}
