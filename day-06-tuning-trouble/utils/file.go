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
	var block_inputs []string = strings.Split(ReadFileAsString(), "\n\n")

	for _, block_input := range block_inputs {
		blocks = append(blocks, strings.Split(block_input, "\n"))
	}
	return
}

func ReadFileAsLines() []string {
	return strings.Split(ReadFileAsString(), "\n")
}

func ReadFileAsString() string {
	return strings.TrimSuffix(string(readFile()), "\n")
}

func readFile() (raw []byte) {
	raw, err := os.ReadFile("input.txt")
	Check(err)

	return
}
