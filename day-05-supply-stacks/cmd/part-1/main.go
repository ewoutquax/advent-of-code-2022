package main

import (
	"fmt"

	"aoc.com/2022/day-05/internal/puzzle"
	"aoc.com/2022/day-05/pkg/rootdir"
	utils "github.com/ewoutquax/aoc-go-utils"
)

func main() {
	fmt.Printf("Result of part0:\n\n%s\n", SolvePuzzlePart())
}

func SolvePuzzlePart() string {
	blocks := utils.ReadFileAsBlocks(rootdir.Get())
	return puzzle.ResolveInstructionsV1(blocks, 9)
}
