package main

import (
	"fmt"

	"github.com/ewoutquax/advent-of-code-2022/day-21/puzzle"
	"github.com/ewoutquax/advent-of-code-2022/utils"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1())
	fmt.Println("Result of part-2: ", solvePart2())
}

func solvePart1() int {
	return puzzle.GetValueRootMonkey(utils.ReadFileAsLines())
}

func solvePart2() int {
	return puzzle.GetNeededValueHuman(utils.ReadFileAsLines())
}
