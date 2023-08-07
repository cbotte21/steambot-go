package piece

import (
	"errors"
	"github.com/cbotte21/chess-go/internal/game/position"
)

func GetPiece(position position.Position, piece Type) (IPiece, error) {
	switch piece {
	case 5:
		king, _ := NewKing(position)
		return king, nil
	case 4:
		queen, _ := NewQueen(position)
		return queen, nil
	case 3:
		rook, _ := NewRook(position)
		return rook, nil
	case 2:
		bishop, _ := NewBishop(position)
		return bishop, nil
	case 1:
		knight, _ := NewKnight(position)
		return knight, nil
	case 0:
		pawn, _ := NewPawn(position)
		return pawn, nil
	}
	return nil, errors.New("invalid piece")
}
