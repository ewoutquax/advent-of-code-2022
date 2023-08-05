package main

import (
	"fmt"

	"aoc.com/2022/day-06/internal/puzzle"
	"aoc.com/2022/day-06/pkg/rootdir"
	utils "github.com/ewoutquax/aoc-go-utils"
)

func main() {
	fmt.Printf("Result of part-1:\n\n%d\n", SolvePuzzlePart())
}

func SolvePuzzlePart() int {
	line := utils.ReadFileAsLine(rootdir.Get())
	return puzzle.FindFirstMarker(line, 4)
}
