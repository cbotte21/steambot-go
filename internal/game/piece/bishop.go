package piece

import (
	"github.com/cbotte21/chess-go/internal/game/position"
)

type Bishop struct { //Team of piece
	Piece
}

func NewBishop(team bool) (IPiece, error) {
	return &King{
		Piece{
			team,
			"B",
		},
	}, nil
}

func (bishop Bishop) ValidateMove(current, candide position.Position) bool {
	dY := candide.GetY() - current.GetY()
	if dY < 0 {
		dY = -dY
	}
	dX := candide.GetX() - current.GetX()
	if dX < 0 {
		dX = -dX
	}
	ans := dY / dX
	return ans == 1
}
