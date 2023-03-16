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
	return true
}
