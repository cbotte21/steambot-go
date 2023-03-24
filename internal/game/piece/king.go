package piece

import (
	"github.com/cbotte21/chess-go/internal/game/position"
)

type King struct { //Team of piece
	Piece
}

func NewKing(team bool) (IPiece, error) {
	return &King{
		Piece{
			team,
			"K",
		},
	}, nil
}

func verify(current, candide int) bool {
	deltaTiles := current - candide
	if deltaTiles < 0 {
		deltaTiles = -deltaTiles
	}
	return deltaTiles == 1 || deltaTiles == 0
}

func (king King) ValidateMove(current, candide position.Position) bool { //Can multiply difference by -1 or piece calculations
	return verify(current.X, candide.X) || verify(current.Y, candide.Y)
}
