package main

import (
	"fmt"
	"strconv"

	"aoc.com/2022/day-02/internal/puzzle"
	"aoc.com/2022/day-02/pkg/utils"
)

func main() {
	fmt.Println("Result of part-2: ", SolvePuzzlePart())
}

func SolvePuzzlePart() string {
	total := puzzle.CalculateTotalScoreByRequiredOutcomes(utils.ReadFileAsLines())

	return strconv.Itoa(total)
}
