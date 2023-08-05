package main

import (
	"fmt"

	"aoc.com/2022/day-05/internal/puzzle"
	"aoc.com/2022/day-05/pkg/rootdir"
	"github.com/ewoutquax/aoc-go-utils"
)

func main() {
	fmt.Printf("Result of part-2:\n\n%s\n", SolvePuzzlePart())
}

func SolvePuzzlePart() string {
	blocks := utils.ReadFileAsBlocks(rootdir.Get())
	return puzzle.ResolveInstructionsV2(blocks, 9)
}
