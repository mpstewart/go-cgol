package cgol

import "fmt"

//   X
//     X
// X X X
func (b *Board) PutGliderAt(x, y int) {
	if tl := b.GetCellAt(x, y); tl == nil {
		panic(
			fmt.Sprintf("could not place glider at %d, %d, not a point on the board", x, y),
		)
	}
	b.GetCellAt(x+0, y+0).KillCell()
	b.GetCellAt(x+1, y+0).VivifyCell()
	b.GetCellAt(x+2, y+0).KillCell()
	b.GetCellAt(x+0, y+1).KillCell()
	b.GetCellAt(x+1, y+1).KillCell()
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
