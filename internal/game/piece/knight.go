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
	return true
}
