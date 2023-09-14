package piece

import (
	"github.com/cbotte21/chess-go/internal/game/position"
)

type Queen struct { //Team of piece
	Piece
}

func NewQueen(position position.Position) (IPiece, error) {
	return &King{
		Piece{
			position,
		},
	}, nil
}

func XOR(x, y error) error {
	if x != nil {
		return x
	}
	return y
}

func (queen Queen) ValidateMove(final position.Position) error {
	return XOR(Rook(queen).ValidateMove(queen.initial), Bishop(queen).ValidateMove(queen.initial))
}
