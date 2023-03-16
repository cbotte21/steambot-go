package position

// A Position containing an X and Y point (int).
type Position struct {
	X, Y int
}

func NewPosition(x, y int) Position {
	return Position{x, y}
}

// SetPosition sets a positions location using cartesian coordinates.
func (position *Position) SetPosition(x, y int) {
	position.X = x
	position.Y = y
}

// GetX returns the X value of a position.
func (position *Position) GetX() int {
	return position.X
}

// GetY returns the Y value of a position/
func (position *Position) GetY() int {
	return position.Y
}
