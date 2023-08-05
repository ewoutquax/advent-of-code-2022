package main

import (
	"fmt"

	"aoc.com/2022/day-06/internal/puzzle"
	"aoc.com/2022/day-06/pkg/rootdir"
	utils "github.com/ewoutquax/aoc-go-utils"
)

func main() {
	fmt.Printf("Result of part-2:\n\n%d\n", SolvePuzzlePart())
}

func SolvePuzzlePart() int {
	return puzzle.FindFirstMarker(utils.ReadFileAsLine(rootdir.Get()), 14)
}
