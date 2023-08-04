package main

import (
	"fmt"

	"aoc.com/2022/day-04/internal/puzzle"
	"aoc.com/2022/day-04/pkg/rootdir"
	utils "github.com/ewoutquax/aoc-go-utils"
)

func main() {
	fmt.Printf("Result of part-2:\n\n%d\n", solvePuzzlePart())
}

func solvePuzzlePart() int {
	return puzzle.CountPartlyOverlaps(utils.ReadFileAsLines(rootdir.Get()))
}
