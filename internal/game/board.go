/*

	Basic 3x3 array. Functions to move Piece objects around.

	Package does not check for height/width constraints
*/

package game

import (
	"errors"
	piece2 "github.com/cbotte21/chess-go/internal/game/piece"
	position2 "github.com/cbotte21/chess-go/internal/game/position"
)

const width int = 8
const height int = 8

type Board [width][height]piece2.IPiece

// NewBoard will initialize a board instance from a string templates.
func NewBoard(template string) Board {
	board := Board{}
	for i := 0; i < width; i++ {
		for k := 0; k < height; k++ {
			currPiece, _ := piece2.GetPiece(string(template[i*width+height]), WHITE)
			position := position2.Position{}
			position.SetPosition(i, k)
			board.SetPosition(position, currPiece)
		}
	}
	return board
}

// InBounds returns whether a position is valid based on the board side
func (board *Board) InBounds(position position2.Position) bool {
	res := position.X >= 0 && position.Y >= 0
	return res && position.X < board.Width() && position.Y < board.Height()
}

// Move a piece from one Position to another.
// Pre-condition: Two valid, inbound positions.
// Post-condition: upon successfully moving, initial will be empty, final will occupy initial's initial IPiece.
func (board *Board) Move(initial, final position2.Position) error {
	p := board.GetPiece(initial)
	if !p.ValidateMove(initial, final) || !board.ValidateMove(p, final) || !board.InBounds(initial) || !board.InBounds(final) {
		return errors.New("invalid move")
	}
	board.SetPosition(initial, board.GetPiece(final))
	board.SetPosition(final, p)
	return nil
}

// GetPiece returns the IPiece in the specified position.
func (board *Board) GetPiece(position position2.Position) piece2.IPiece {
	return board[position.GetX()][position.GetY()]
}

// SetPosition sets a desired position to an IPiece.
func (board *Board) SetPosition(position position2.Position, piece piece2.IPiece) {
	board[position.GetX()][position.GetY()] = piece
}

// ValidateMove tells whether a desired position change is allowed on the board.
func (board *Board) ValidateMove(p piece2.IPiece, final position2.Position) bool {
	return p.GetTeam() != board.GetPiece(final).GetTeam()
}

// ToString returns a string representation of pieces, every piece has {team_identifier}{piece_identifier}
func (board *Board) ToString() string {
	var str string = ""
	for i := 0; i < width; i++ {
		for k := 0; k < height; k++ {
			if board[i][k].GetTeam() { //Add team
				str += "b"
			} else {
				str += "w"
			}
			str += board[i][k].GetIdentifier()
		}
	}
	return str
}

// Width returns the width of the 2d array
func (board *Board) Width() int {
	return width
}

// Height returns the height of the 2d array
func (board *Board) Height() int {
	return height
}
