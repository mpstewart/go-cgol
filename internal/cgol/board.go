package cgol

type Board struct {
	Height int
	Width  int
	Cells  [][]*Cell
}

func NewBoard(width, height int) *Board {
	newBoard := new(Board)
	cells := make([][]*Cell, 0)

	for i := 0; i < height; i += 1 {
		row := make([]*Cell, width)
		for j := 0; j < width; j += 1 {
			row[j] = &Cell{State: CellStateDead}
		}
		cells = append(cells, row)
	}

	newBoard.Height = height
	newBoard.Width = width
	newBoard.Cells = cells

	return newBoard
}

func (b *Board) NextBoard() *Board {
	nextBoard := NewBoard(b.Width, b.Height)

	calculateCell := func(x int, y int, c *Cell) {
		cs := b.NextCellState(x, y)
		nextBoard.Cells[y][x] = &Cell{
			State: cs,
		}
	}

	b.DoForEachCell(calculateCell)

	return nextBoard
}

func (b *Board) NextCellState(x, y int) CellState {
	c := b.GetCellAt(x, y)
	n := b.CountLiveNeighbors(x, y)

	if c.State == CellStateAlive {
		if n == 2 || n == 3 {
			return CellStateAlive
		}
		return CellStateDead
	} else if c.State == CellStateDead {
		if n == 3 {
			return CellStateAlive
		}
	}

	return CellStateDead
}

func (b *Board) CountLiveNeighbors(x, y int) int {
	neighbors := b.GetNeighborCells(x, y)
	n := 0
	for _, c := range neighbors {
		if c != nil && c.State == CellStateAlive {
			n += 1
		}
	}

	return n
}

func (b *Board) GetNeighborCells(x, y int) []*Cell {
	return []*Cell{
		b.GetCellAt(x-1, y-1),
		b.GetCellAt(x, y-1),
		b.GetCellAt(x+1, y-1),
		b.GetCellAt(x-1, y),
		b.GetCellAt(x+1, y),
		b.GetCellAt(x-1, y+1),
		b.GetCellAt(x, y+1),
		b.GetCellAt(x+1, y+1),
	}
}

func (b *Board) GetCellAt(x, y int) *Cell {
	if x < 0 || y < 0 || x >= b.Width || y >= b.Height {
		return nil
	}

	return b.Cells[y][x]
}

func (b *Board) DoForEachCell(f func(int, int, *Cell)) {
	for y := 0; y < b.Height; y += 1 {
		for x := 0; x < b.Width; x += 1 {
			c := b.GetCellAt(x, y)
			f(x, y, c)
		}
	}
}
