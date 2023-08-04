package main

import (
	"fmt"
	"strconv"

	"aoc.com/2022/day-03/internal/puzzle"
	"aoc.com/2022/day-03/pkg/utils"
)

func main() {
	fmt.Print("Result of part-1:\n", solvePuzzlePart(), "\n")
}

func solvePuzzlePart() (score string) {
	total := puzzle.CalculateTotalScore(utils.ReadFileAsLines())

	return strconv.Itoa(total)
}
