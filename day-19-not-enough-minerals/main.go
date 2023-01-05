package main

import (
	"fmt"

	"aoc.com/2022/day-19/puzzle"
	"aoc.com/2022/day-19/utils"
)

func main() {
	fmt.Println("Result of part-1.5: ", solvePart1())
}

func solvePart1() int {
	return puzzle.SumBlueprintQualities(utils.ReadFileAsLines())
}
