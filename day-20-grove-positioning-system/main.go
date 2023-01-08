package main

import (
	"fmt"

	"aoc.com/2022/day-20/puzzle"
	"aoc.com/2022/day-20/utils"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1())
	fmt.Println("Result of part-2: ", solvePart2())
}

func solvePart1() int {
	return puzzle.GrooveHash(utils.ReadFileAsNumbers())
}

func solvePart2() int {
	return puzzle.GrooveHashHard(utils.ReadFileAsNumbers())
}
