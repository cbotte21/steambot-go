package piece

import (
	"github.com/cbotte21/chess-go/internal/game"
	"github.com/cbotte21/chess-go/internal/game/position"
)

type Knight struct { //Team of piece
	Piece
}

func NewKnight(position position.Position) (IPiece, error) {
	return &King{
		Piece{
			position,
		},
	}, nil
}

func (knight Knight) ValidateMove(final position.Position, state game.Game) error {
	dY := final.GetY() - knight.initial.GetY()
	if dY < 0 {
		dY = -dY
	}
	dX := final.GetX() - knight.initial.GetX()
	if dX < 0 {
		dX = -dX
	}
	if (dY != 2 || dX != 1) && (dX != 2 || dY != 1) {
		return InvalidMoveError()
	}
	return nil
}
