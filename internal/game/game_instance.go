package game

import (
	"errors"
	"github.com/cbotte21/chess-go/internal/game/position"
)

/*

	White is turn 0

	//Include play... check piece.validate_move(cordinate 1, cordinate 2) && board.move
*/

// Game contains information relevant to the chess match.
type Game struct {
	board *string
	turn  *bool //0 == white
}

func (game *Game) GetBoard() string {
	return *game.board
}

// IsTurn returns a boolean if it is a player's turn
func (game *Game) IsTurn() bool {
	return *game.turn
}

// NewGame an instance of a Game.
func (game *Game) NewGame(state *string, playersTurn *bool) {
	game.board = state
	game.turn = playersTurn
}

// Move from the initial Position to the final Position.
// Returns: true on success
func (game *Game) Move(player string, initial, final position.Position) error {
	board := NewBoard(*game.board)
	if game.IsTurn() {
		return board.Move(initial, final)
	}
	return errors.New("it is not your turn")
}
