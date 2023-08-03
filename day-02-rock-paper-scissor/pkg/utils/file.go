package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
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
	fullPath := fmt.Sprintf("%s/%s/%s", basepath, "../..", "input.txt")
	absPath, _ := filepath.Abs(fullPath)

	raw, err := os.ReadFile(absPath)
	check(err)

	return strings.TrimSuffix(string(raw), "\n")
}
