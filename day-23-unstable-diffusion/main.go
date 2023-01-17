package main

import (
	"fmt"

	"github.com/ewoutquax/advent-of-code-2022/day-23/puzzle"
	"github.com/ewoutquax/advent-of-code-2022/utils"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1())
	fmt.Println("Result of part-2: ", solvePart2())
}

func solvePart1() int {
	return puzzle.CountEmptyGroundTilesAfterRounds(10, utils.ReadFileAsLines())
}

func solvePart2() int {
	return puzzle.FirstRoundWithoutMovement(utils.ReadFileAsLines())
}
