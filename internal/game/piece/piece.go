package piece

import (
	"github.com/cbotte21/chess-go/internal/game/position"
)

type IPiece interface {
	ValidateMove(current, candide position.Position) bool
	SetTeam(team bool)
	GetTeam() bool
	GetIdentifier() string
}

type Piece struct {
	team       bool
	identifier string
}

func (piece Piece) GetTeam() bool {
	return piece.team
}

func (piece *Piece) SetTeam(team bool) {
	piece.team = team
}

func (piece *Piece) GetIdentifier() string {
	return piece.identifier
}
