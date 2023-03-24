package piece

import (
	"github.com/cbotte21/chess-go/internal/game/position"
)

type Knight struct { //Team of piece
	Piece
}

func NewKnight(team bool) (IPiece, error) {
	return &King{
		Piece{
			team,
			"N",
		},
	}, nil
}

func (knight Knight) ValidateMove(current, candide position.Position) bool {
	dY := candide.GetY() - current.GetY()
	if dY < 0 {
		dY = -dY
	}
	dX := candide.GetX() - current.GetX()
	if dX < 0 {
		dX = -dX
	}
	return dY == 2 && dX == 1 || dX == 2 && dY == 1
}
