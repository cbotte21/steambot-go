package piece

import "errors"

/*
Precondition: identifier should contain two characters. <TEAM_IDENTIFIER><PIECE_IDENTIFIER>
*/
func GetPiece(identifier string, team bool) (IPiece, error) {
	switch identifier {
	case "k":
		king, _ := NewKing(team)
		return king, nil
	case "q":
		queen, _ := NewQueen(team)
		return queen, nil
	case "r":
		rook, _ := NewRook(team)
		return rook, nil
	case "b":
		bishop, _ := NewBishop(team)
		return bishop, nil
	case "n":
		knight, _ := NewKnight(team)
		return knight, nil
	case "p":
		pawn, _ := NewPawn(team)
		return pawn, nil
	}
	return nil, nil
}

func GetPieceError() error {
	return errors.New("could not allocate piece")
}
