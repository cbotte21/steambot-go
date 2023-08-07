package piece

import (
	"github.com/cbotte21/chess-go/internal/game"
	"github.com/cbotte21/chess-go/internal/game/position"
)

type Bishop struct { //Team of piece
	Piece
}

func NewBishop(position position.Position) (IPiece, error) {
	return &Bishop{
		Piece{
			position,
		},
	}, nil
}

func (bishop Bishop) ValidateMove(final position.Position, state game.Game) error {
	dY := final.GetY() - bishop.initial.GetY()
	if dY < 0 {
		dY = -dY
	}
	dX := final.GetX() - bishop.initial.GetX()
	if dX < 0 {
		dX = -dX
	}
	ans := dY / dX
	if ans != 1 {
		return InvalidMoveError()
	}
	return nil
}
