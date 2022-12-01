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

func solvePart1(lines []string) (max_calories int) {
	if len(lines) == 0 {
		lines = read_file_as_lines()
	}

	var calories int = 0

	max_calories = 0
	for _, line := range lines {
		if line == "" {
			if max_calories < calories {
				max_calories = calories
			}
			calories = 0
		} else {
			calories += conv_str_to_i(line)
		}
	}
	return
}

func solvePart2(lines []string) int {
	if len(lines) == 0 {
		lines = read_file_as_lines()
	}

	var calories int = 0
	var list = []int{}

	for _, line := range lines {
		if line == "" {
			list = append(list, calories)
			calories = 0
		} else {
			calories += conv_str_to_i(line)
		}
	}
	list = append(list, calories)

	sort.Ints(list)

	return list[len(list)-1] + list[len(list)-2] + list[len(list)-3]
}

func read_file_as_numbers() (numbers []int) {
	var lines []string = read_file_as_lines()

	for _, string := range lines {
		numbers = append(numbers, conv_str_to_i(string))
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
