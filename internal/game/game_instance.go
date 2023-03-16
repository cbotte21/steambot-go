package game

import (
	"errors"
	"github.com/cbotte21/chess-go/internal/game/position"
	"github.com/cbotte21/chess-go/internal/game/templates"
)

/*

	White is turn 0

	//Include play... check piece.validate_move(cordinate 1, cordinate 2) && board.move
*/

// Game contains information relevant to the chess match.
type Game struct {
	player1, player2 string
	ranked           bool
	board            Board
	turn             bool //0 == white
}

const (
	WHITE bool = false
	BLACK      = true
)

// HasPlayer returns true if a match contains a player, otherwise false.
func (game *Game) HasPlayer(player string) bool {
	return player == game.player1 || player == game.player2
}

func (game *Game) GetBoard() string {
	return game.board.ToString()
}

// IsTurn returns a boolean if it is a player's turn
func (game *Game) IsTurn(player string) bool {
	return player == game.player1 && !game.turn || player == game.player2 && game.turn
}

// Initialize an instance of a Game.
func (game *Game) Initialize(template string, ranked bool, player1, player2 string) {
	game.player1 = player1
	game.player2 = player2
	game.board = Board{}
	game.ranked = ranked
	game.turn = WHITE

	game.board = NewBoard(templates.GetTemplate(template))
}

// Move from the initial Position to the final Position.
// Returns: true on success
func (game *Game) Move(player string, initial, final position.Position) error {
	if game.IsTurn(player) {
		return game.board.Move(initial, final)
	}
	return errors.New("it is the other player's turn")
}
