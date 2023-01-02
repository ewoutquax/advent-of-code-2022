package utils

import (
	"os"
	"strings"
)

func ReadFileAsNumbers() (numbers []int) {
	var lines []string = ReadFileAsLines()

	for _, string := range lines {
		numbers = append(numbers, ConvStrToI(string))
	}

	return
}

func ReadFileAsBlocks() (blocks [][]string) {
	var block_inputs []string = strings.Split(readFile(), "\n\n")

	for _, block_input := range block_inputs {
		blocks = append(blocks, strings.Split(block_input, "\n"))
	}
	return
}

func ReadFileAsLines() []string {
	return strings.Split(readFile(), "\n")
}

func readFile() string {
	raw, err := os.ReadFile("input.txt")
	Check(err)

	return strings.TrimSuffix(string(raw), "\n")
}
