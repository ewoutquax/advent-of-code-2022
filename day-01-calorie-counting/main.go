package main

import (
	"aoc.com/2022/day-01/puzzle"
	"aoc.com/2022/day-01/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1())
	fmt.Println("Result of part-2: ", solvePart2())
}

func solvePart1() int {
	var blocks [][]string = readFileAsBlocks()

	return puzzle.TotalTopN(blocks, 1)
}

func solvePart2() int {
	var blocks [][]string = readFileAsBlocks()
	return puzzle.TotalTopN(blocks, 3)
}

func readFileAsNumbers() (numbers []int) {
	var lines []string = readFileAsLines()

	for _, string := range lines {
		numbers = append(numbers, utils.ConvStrToI(string))
	}

	return
}

func readFileAsBlocks() (blocks [][]string) {
	var block_inputs []string = strings.Split(readFile(), "\n\n")

	for _, block_input := range block_inputs {
		blocks = append(blocks, strings.Split(block_input, "\n"))
	}
	return
}

func readFileAsLines() []string {
	return strings.Split(readFile(), "\n")
}

func readFile() string {
	raw, err := os.ReadFile("input.txt")
	utils.Check(err)

	return strings.TrimSuffix(string(raw), "\n")
}
