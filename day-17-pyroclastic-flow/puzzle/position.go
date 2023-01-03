package puzzle

type Position struct {
	x int
	y int
}

func (pos Position) toValue() int {
	return pos.y*7 + pos.x
}
