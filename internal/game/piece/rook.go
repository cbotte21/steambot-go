package piece

import (
	"github.com/cbotte21/chess-go/internal/game"
	"github.com/cbotte21/chess-go/internal/game/position"
)

type Rook struct { //Team of piece
	Piece
}

func NewRook(position position.Position) (IPiece, error) {
	return &Rook{
		Piece{
			position,
		},
	}, nil
}

func (rook Rook) ValidateMove(final position.Position, state game.Game) error {
	dY := final.GetY() - rook.initial.GetY()
	if dY < 0 {
		dY = -dY
	}
	dX := final.GetX() - rook.initial.GetX()
	if dX < 0 {
		dX = -dX
	}

	if dY > 0 && dX == 0 || dX > 0 && dY == 0 {
		return nil
	}
	return InvalidMoveError()
}
