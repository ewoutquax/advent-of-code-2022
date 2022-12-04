package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1([]string{}))
	fmt.Println("Result of part-2: ", solvePart2([]string{}))
}

func solvePart1(lines []string) (count int) {
	if len(lines) == 0 {
		lines = read_file_as_lines()
	}

	for _, line := range lines {
		if completeOverlap(line) {
			count += 1
		}
	}

	return
}

func solvePart2(lines []string) (count int) {
	if len(lines) == 0 {
		lines = read_file_as_lines()
	}

	for _, line := range lines {
		if hasOverlap(line) {
			count += 1
		}
	}

	return
}

func completeOverlap(line string) (result bool) {
	parts := strings.Split(line, ",")

	var left_sections []int = expandSections(string(parts[0]))
	var right_sections []int = expandSections(string(parts[1]))

	all_sections := unique(append(left_sections, right_sections...))
	sort.Ints(all_sections)

	if equalSlices(all_sections, left_sections) ||
		equalSlices(all_sections, right_sections) {
		return true
	} else {
		return false
	}
}

func hasOverlap(line string) bool {
	parts := strings.Split(line, ",")

	var left_sections []int = expandSections(string(parts[0]))
	var right_sections []int = expandSections(string(parts[1]))

	all_sections := append(left_sections, right_sections...)
	unique_sections := unique(all_sections)

	return len(all_sections) != len(unique_sections)
}

func expandSections(section string) (sections []int) {
	var from_to []string = strings.Split(section, "-")

	var from int = conv_str_to_i(from_to[0])
	var to int = conv_str_to_i(from_to[1])

	for idx := from; idx <= to; idx += 1 {
		sections = append(sections, idx)
	}

	return
}

func unique(intSlice []int) (list []int) {
	var keys = make(map[int]bool)

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return
}

func equalSlices(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
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
