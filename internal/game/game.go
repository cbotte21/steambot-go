/*
	BitBoard
*/

package game

import (
	"errors"
	"github.com/cbotte21/chess-go/internal/game/piece"
	"github.com/cbotte21/chess-go/internal/game/position"
	"github.com/cbotte21/chess-go/schema"
)

type Game []int64

const (
	PAWN      = 0
	KNIGHT    = 1
	ROOK      = 2
	BISHOP    = 3
	QUEEN     = 4
	KING      = 5
	PIECEITER = 5

	PLAYERONE = 6
	ENPASSANT = 7
	ALLITER   = 7
)

// NewGame will initialize a game instance.
func NewGame(template string) Game {
	board := Game{}
	// Piece bitboards
	for i, identifier := range template {
		pieceType := board.getTypeFromIdentifier(identifier)
		var offset int64 = 2 << i

		if pieceType == -1 {
			continue
		}

		board[pieceType] = board[pieceType] | offset
	}

	for i := 0; i < 64; i++ {
		var offset int64 = 2 << i

		if i < 16 { // Populate team bitboard
			board[PLAYERONE] = board[PLAYERONE] | offset
		}
		if (i >= 8 && i < 16) || (i < 56 && i >= 48) { // Populate enpassant bitboard
			board[ENPASSANT] = board[ENPASSANT] | offset
		}
	}

	return board
}

func LoadGame(cachedGamePointer *schema.CachedGame) Game {
	cachedGame := *cachedGamePointer
	game := Game{}
	game[PAWN] = cachedGame.Pawns
	game[KNIGHT] = cachedGame.Knights
	game[ROOK] = cachedGame.Rooks
	game[BISHOP] = cachedGame.Bishops
	game[QUEEN] = cachedGame.Queens
	game[KING] = cachedGame.Kings
	game[PLAYERONE] = cachedGame.P1BitBoard
	game[ENPASSANT] = cachedGame.Enpassants
	return game
}

func (game *Game) ExportUpdate(cachedGame schema.CachedGame) schema.CachedGame {
	cachedGame.Pawns = (*game)[PAWN]
	cachedGame.Knights = (*game)[KNIGHT]
	cachedGame.Bishops = (*game)[BISHOP]
	cachedGame.Rooks = (*game)[ROOK]
	cachedGame.Queens = (*game)[QUEEN]
	cachedGame.Kings = (*game)[KING]

	cachedGame.P1BitBoard = (*game)[PLAYERONE]
	cachedGame.Enpassants = (*game)[ENPASSANT]
	cachedGame.Turn = !cachedGame.Turn
	return cachedGame
}

func (game *Game) Move(initial, final position.Position, isWhite bool) error {
	if game.isAlly(final, isWhite) {
		return errors.New("you cannot capture you're team")
	}

	pieceType := game.getTypeFromPosition(initial)
	if pieceType == -1 {
		return errors.New("no piece belongs to square")
	}

	p, err := piece.GetPiece(initial, pieceType)
	if err != nil {
		return err
	}

	err = p.ValidateMove(final, *game)
	if err != nil {
		return err
	}

	game.clearFlagUniverse(final)
	game.moveFlag(pieceType, initial, final)
	game.moveFlag(PLAYERONE, initial, final)
	return nil
}

func (game *Game) moveFlag(piece piece.Type, initial, position position.Position) {
	offset := position.Offset() & (*game)[piece]
	// Set final
	(*game)[piece] = (*game)[piece] ^ offset
	// Delete initial
	var maxVal int64 = (2 << 65) - 1
	flag := maxVal ^ offset
	(*game)[piece] = (*game)[piece] ^ flag
}

func (game *Game) clearFlagUniverse(position position.Position) {
	var maxVal int64 = (2 << 65) - 1
	flag := maxVal ^ (2 << (position.X*8 + position.Y))

	for _, state := range *game {
		state = state ^ flag
	}
}

func (game *Game) match(position position.Position, state int) bool {
	var match bool = false
	if (position.Offset() & (*game)[state]) != 0 {
		match = true
	}
	return match
}

func (game *Game) getTypeFromPosition(position position.Position) piece.Type {
	for i := 0; i <= KING; i++ {
		if game.match(position, i) {
			return piece.Type(i)
		}
	}
	return -1
}

func (game *Game) getTypeFromIdentifier(identifier int32) piece.Type {
	switch identifier {
	case 'P':
		return PAWN
	case 'N':
		return KNIGHT
	case 'B':
		return BISHOP
	case 'R':
		return ROOK
	case 'Q':
		return QUEEN
	case 'K':
		return KING
	}
	return -1
}

func (game *Game) isAlly(position position.Position, isWhite bool) bool {
	// Get players bitboard
	var bitboard int64 = (*game)[PLAYERONE]
	if !isWhite {
		bitboard = game.playerTwoBitBoard()
	}

	for i := 0; i <= KING; i++ {
		if game.match(position, i) {
			if (position.Offset() & bitboard) == 0 {
				return true
			}
			break
		}
	}
	return false
}

func (game *Game) playerTwoBitBoard() int64 {
	var bitboard int64 = 0
	for i := 0; i < 64; i++ {
		offset := 2 << i
		if (offset & (*game)[PLAYERONE]) == 0 {
			for k := 0; k < KING; k++ {
				if (offset & (*game)[k]) != 0 {
					bitboard = bitboard | offset
				}
			}
		}
	}
	return bitboard
}
