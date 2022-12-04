package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1([][]string{{}}))
	// fmt.Println("Result of part-2: ", solvePart2([]string{[]string{}}))
}

func solvePart1(blocks [][]string) int {
	if len(blocks) == 1 {
		blocks = read_file_as_blocks()
	}

	return totalTopN(blocks, 1)
}

func solvePart2(blocks [][]string) int {
	if len(blocks) == 1 {
		blocks = read_file_as_blocks()
	}

	return totalTopN(blocks, 3)
}

func totalTopN(blocks [][]string, top_N int) (total int) {
	if len(blocks) == 1 {
		blocks = read_file_as_blocks()
	}

	var calories []int
	for _, lines := range blocks {
		elf := 0
		for _, line := range lines {
			elf += conv_str_to_i(line)
		}
		calories = append(calories, elf)
	}

	sort.Ints(calories)

	for i := 1; i <= top_N; i += 1 {
		total += calories[len(calories) - i]
	}

	return
}

func read_file_as_numbers() (numbers []int) {
	var lines []string = read_file_as_lines()

	for _, string := range lines {
		numbers = append(numbers, conv_str_to_i(string))
	}

	return
}

func read_file_as_blocks() (blocks [][]string) {
	var block_inputs []string = strings.Split(read_file(), "\n\n")

	for _, block_input := range block_inputs {
		blocks = append(blocks, strings.Split(block_input, "\n"))
	}
	return
}

func read_file_as_lines() []string {
	return strings.Split(read_file(), "\n")
}

func read_file() (content string) {
	raw, err := os.ReadFile("input.txt")
	check(err)
	content = strings.TrimSuffix(string(raw), "\n")
	return
}

func conv_str_to_i(s string) (i int) {
	i, err := strconv.Atoi(s)
	check(err)
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
