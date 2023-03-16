package piece

import (
	"github.com/cbotte21/chess-go/internal/game/position"
)

type Queen struct { //Team of piece
	Piece
}

func NewQueen(team bool) (IPiece, error) {
	return &King{
		Piece{
			team,
			"Q",
		},
	}, nil
}

func (queen Queen) ValidateMove(current, candide position.Position) bool {
	return true
}
