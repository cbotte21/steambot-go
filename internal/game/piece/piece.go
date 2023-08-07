package piece

import (
	"errors"
	"github.com/cbotte21/chess-go/internal/game"
	"github.com/cbotte21/chess-go/internal/game/position"
)

type Type int

type IPiece interface {
	ValidateMove(final position.Position, state game.Game) error
}

type Piece struct {
	initial position.Position
}

func InvalidMoveError() error {
	return errors.New("invalid move")
}
