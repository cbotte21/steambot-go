package piece

import (
	"github.com/cbotte21/chess-go/internal/game"
	"github.com/cbotte21/chess-go/internal/game/position"
)

type Pawn struct { //Team of piece
	Piece
}

func NewPawn(position position.Position) (IPiece, error) {
	return &Pawn{
		Piece{
			position,
		},
	}, nil
}

func (pawn Pawn) ValidateMove(final position.Position, state game.Game) error {
	return nil
}
