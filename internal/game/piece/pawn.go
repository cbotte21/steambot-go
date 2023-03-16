package piece

import (
	"github.com/cbotte21/chess-go/internal/game/position"
)

type Pawn struct { //Team of piece
	Piece
}

func NewPawn(team bool) (IPiece, error) {
	return &Pawn{
		Piece{
			team,
			"P",
		},
	}, nil
}

func (pawn Pawn) ValidateMove(current, candide position.Position) bool {
	return true
}
