package main

import (
	"fmt"

	"aoc.com/2022/day-12/puzzle"
	"aoc.com/2022/day-12/utils"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1())
	fmt.Println("Result of part-2: ", solvePart2())
}

func solvePart1() int {
	return puzzle.FindLowestDistance(utils.ReadFileAsLines())
}

func solvePart2() int {
	return puzzle.FindLowestDistanceDown(utils.ReadFileAsLines())
}
