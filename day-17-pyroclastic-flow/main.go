package main

import (
	"fmt"

	"aoc.com/2022/day-17/puzzle"
	"aoc.com/2022/day-17/utils"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1())
	fmt.Println("Result of part-2: ", solvePart2())
}

func solvePart1() int {
	return puzzle.HeightStackedBlocks(utils.ReadFileAsString(), 2022)
}

func solvePart2() int {
	offsetBlocks := 1740 + 140 - 10
	offsetHeight := puzzle.HeightStackedBlocks(utils.ReadFileAsString(), offsetBlocks)
	fmt.Println("offsetHeight after", offsetBlocks, "blocks:", offsetHeight)
	return (1000000000000-offsetBlocks)/1730*2659 + offsetHeight
}
