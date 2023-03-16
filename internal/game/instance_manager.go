package game

import (
	"errors"
)

/*

	Must find better way of storing players. Including sockets.

*/

// DefaultTemplate is a string containing the initial board layout for a game.
// The string is formatted as such...
// <TEAM_IDENTIFIER><PIECE_IDENTIFIER>x64
// These two variables should both be one character, and repeated for a total of 64 times.
// First 2 characters will represent the bottom right of the board.
const DefaultTemplate string = "default" //Extract to class with enums

// Instances contains an array of Game
type Instances struct {
	games []Game
}

// CreateGame constructs a Game. Should only be called internally by GameQueue.
func (manager *Instances) CreateGame(player1, player2 string, ranked bool) error {
	game := Game{}
	game.Initialize(DefaultTemplate, ranked, player1, player2)
	manager.games = append(manager.games, game)
	return nil
}

// GetGame returns a Game containing a player
func (manager *Instances) GetGame(player string) (Game, error) {
	for _, game := range manager.games {
		if game.HasPlayer(player) {
			return game, nil
		}
	}
	return Game{}, errors.New("could not find game")
}

// DisbandGame abruptly ends a Game instance. Deleting it from all records.
func (manager *Instances) DisbandGame(game Game) { //TODO: Does this work?!?!? LMAO
	for i, g := range manager.games {
		if g == game {
			for i < len(manager.games)-1 {
				manager.games[i] = manager.games[i+1]
				i++
			}
			manager.games = manager.games[:len(manager.games)-1]
		}
	}
}

// InGame returns whether a schema.Player is currently in a match.
func (manager *Instances) InGame(player string) bool {
	_, err := manager.GetGame(player)
	return err == nil
}

// QuitGame reports the game as a forfeit by player.
func (manager *Instances) QuitGame(player string) {
	game, err := manager.GetGame(player)
	//TODO: Insert game to database
	if err == nil {
		manager.DisbandGame(game)
	}
}
