package cgol

import (
	"fmt"
)

type CellState int

const (
	CellStateDead CellState = iota
	CellStateAlive
)

type Cell struct {
	State CellState
}

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

	// fmt.Printf("%+v", newBoard)

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

func (b *Board) PutGliderAt(x, y int) {
	if tl := b.GetCellAt(x, y); tl == nil {
		panic(
			fmt.Sprintf("could not place glider at %d, %d, not a point on the board", x, y),
		)
	}
	b.GetCellAt(x+1, y+0).State = CellStateAlive
	b.GetCellAt(x+2, y+1).State = CellStateAlive
	b.GetCellAt(x+0, y+2).State = CellStateAlive
	b.GetCellAt(x+1, y+2).State = CellStateAlive
	b.GetCellAt(x+2, y+2).State = CellStateAlive

}

func (b *Board) PutSquareAt(x, y int) {
	if tl := b.GetCellAt(x, y); tl == nil {
		panic(
			fmt.Sprintf("could not put square at %d, %d, not a point on the board", x, y),
		)
	}
	b.GetCellAt(x, y).State = CellStateAlive
	b.GetCellAt(x+1, y).State = CellStateAlive
	b.GetCellAt(x, y+1).State = CellStateAlive
	b.GetCellAt(x+1, y+1).State = CellStateAlive
}

func (b *Board) PutLightweightSpaceShip(x, y int) {
	if tl := b.GetCellAt(x, y); tl == nil {
		panic(
			fmt.Sprintf("could not put lwss at %d, %d, not a point on the board", x, y),
		)
	}
	b.GetCellAt(x+1, y).State = CellStateAlive
	b.GetCellAt(x+4, y).State = CellStateAlive
	b.GetCellAt(x, y+1).State = CellStateAlive
	b.GetCellAt(x, y+2).State = CellStateAlive
	b.GetCellAt(x, y+3).State = CellStateAlive
	b.GetCellAt(x+4, y+2).State = CellStateAlive
	b.GetCellAt(x+1, y+3).State = CellStateAlive
	b.GetCellAt(x+2, y+3).State = CellStateAlive
	b.GetCellAt(x+3, y+3).State = CellStateAlive

}

//   x     x
//     x   x   x
// x x x   x x
