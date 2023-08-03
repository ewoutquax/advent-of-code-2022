package puzzle

import (
	"strconv"
	"strings"

	"github.com/ewoutquax/advent-of-code-2022/utils"
)

type Position struct {
	x int // X-coordinate of the position
	y int // Y-coordinate of the position; 0 is at the top, 200 is in the void
}

// Create a key for the current position: "x,y"
func (p Position) toKey() string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

func keyToPosition(k string) Position {
	parts := strings.Split(k, ",")

	return Position{
		x: utils.ConvStrToI(parts[0]),
		y: utils.ConvStrToI(parts[1]),
	}
}

func keyWithStepsToEndPosition(from Position, to Position, nrSteps int) (current Position) {
	current.x = from.x
	current.y = from.y

	if from.x != to.x {
		current.x = from.x + (to.x-from.x)/utils.Abs(to.x-from.x)*nrSteps
	}
	if from.y != to.y {
		current.y = from.y + (to.y-from.y)/utils.Abs(to.y-from.y)*nrSteps
	}

	return current
}
