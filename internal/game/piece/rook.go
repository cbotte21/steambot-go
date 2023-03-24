package piece

import (
	"github.com/cbotte21/chess-go/internal/game/position"
)

type Rook struct { //Team of piece
	Piece
}

func NewRook(team bool) (IPiece, error) {
	return &Rook{
		Piece{
			team,
			"R",
		},
	}, nil
}

func (rook Rook) ValidateMove(current, candide position.Position) bool {
	dY := candide.GetY() - current.GetY()
	if dY < 0 {
		dY = -dY
	}
	dX := candide.GetX() - current.GetX()
	if dX < 0 {
		dX = -dX
	}
	return dY > 0 && dX == 0 || dX > 0 && dY == 0
}
