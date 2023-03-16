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

// ValidSlope returns true if delta_slope == 1
func ValidSlope() {

}

func (bishop Bishop) ValidateMove(current, candide position.Position) bool {
	return true
}
