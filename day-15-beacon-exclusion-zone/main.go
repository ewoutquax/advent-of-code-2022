package main

import (
	"fmt"

	"aoc.com/2022/day-15/puzzle"
	"aoc.com/2022/day-15/utils"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1(2000000))
	fmt.Println("Result of part-2: ", solvePart2())
}

func solvePart1(row int) int {
	return puzzle.AllExcludesOnRow(utils.ReadFileAsLines(), row)
}

func solvePart2() int {
	return puzzle.FindOpenLocationHash(utils.ReadFileAsLines())
}
