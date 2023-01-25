package main

import (
	"fmt"

	"github.com/ewoutquax/advent-of-code-2022/day-25/puzzle"
	"github.com/ewoutquax/advent-of-code-2022/utils"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1())
}

func solvePart1() string {
	return puzzle.SumSnafusAsSnafu(utils.ReadFileAsLines())
}
