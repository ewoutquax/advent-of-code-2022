package puzzle

import "strconv"

type Position struct {
	x int // Column of the location in the maze
	y int // Row of the location in the maze
}

// Convert the position to a string key with the `x` and y, joined by a comma (,)
func (p Position) toKey() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}
