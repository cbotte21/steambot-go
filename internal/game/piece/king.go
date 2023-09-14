package piece

import (
	"github.com/cbotte21/chess-go/internal/game/position"
)

type King struct { //Team of piece
	Piece
}

func NewKing(position position.Position) (IPiece, error) {
	return &King{
		Piece{
			position,
		},
	}, nil
}

func verify(current, candide int) error {
	deltaTiles := current - candide
	if deltaTiles < 0 {
		deltaTiles = -deltaTiles
	}
	if deltaTiles != 1 && deltaTiles != 0 {
		return InvalidMoveError()
	}
	return nil
}

func (king King) ValidateMove(final position.Position) error { //Can multiply difference by -1 or piece calculations
	err := verify(king.initial.X, final.X)
	if err != nil {
		return err
	}
	return verify(king.initial.Y, final.Y)
}
