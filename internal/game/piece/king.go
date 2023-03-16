package piece

import (
	"github.com/cbotte21/chess-go/internal/game/position"
	"math"
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
	const epsilon = .001
	return math.Abs(float64(current-candide)-1) < epsilon || current-candide == 0
}

func (king King) ValidateMove(current, candide position.Position) bool { //Can multiply difference by -1 or piece calculations
	return verify(current.X, candide.X) || verify(current.Y, candide.Y)
}
