package main

import (
	"fmt"

	"aoc.com/2022/day-06/puzzle"
	"aoc.com/2022/day-06/utils"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1())
	fmt.Println("Result of part-2: ", solvePart2())
}

func solvePart1() int {
	return puzzle.FindFirstMarker(utils.ReadFileAsString(), 4)
}

func solvePart2() int {
	return puzzle.FindFirstMarker(utils.ReadFileAsString(), 14)
}
