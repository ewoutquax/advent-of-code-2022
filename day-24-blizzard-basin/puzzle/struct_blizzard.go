package puzzle

type Direction uint

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Blizzard struct {
	position  Position  // Position of the blizzard inside the basin
	direction Direction // Direction the blizzard is moving to
}

// Get a string-representation of the position of the blizzard
func (b Blizzard) toKey() string {
	return b.position.toKey()
}

func (b Blizzard) directionChar() (char string) {
	switch b.direction {
	case Up:
		char = "^"
	case Right:
		char = ">"
	case Down:
		char = "v"
	case Left:
		char = "<"
	}

	return
}
