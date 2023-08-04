package main

import (
	"fmt"
	"strconv"

	"aoc.com/2022/day-02/internal/puzzle"
	"aoc.com/2022/day-02/pkg/rootdir"
	utils "github.com/ewoutquax/aoc-go-utils"
)

func main() {
	fmt.Printf("Result of part0:\n\n%s\n", SolvePuzzlePart())
}

func SolvePuzzlePart() string {
	total := puzzle.CalculateTotalScoreForRounds(utils.ReadFileAsLines(rootdir.Get()))

	return strconv.Itoa(total)
}
