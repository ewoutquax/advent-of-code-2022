package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1("", 4))
	fmt.Println("Result of part-2: ", solvePart1("", 14))
}

func solvePart1(input string, size int) int {
	if len(input) == 0 {
		input = read_file()
	}

	return findFirstMarker(input, size)
}

func findFirstMarker(input string, size int) int {
	for idx := size; idx < len(input); idx += 1 {
		message := input[idx-size : idx]

		if unique(message) == size {
			return idx
		}
	}

	panic("No marker found")
}

func unique(message string) int {
	var slice []string = strings.Split(message, "")

	var keys = make(map[string]bool)
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
		}
	}

	return len(keys)
}

func read_file() (content string) {
	raw, err := os.ReadFile("input.txt")
	check(err)
	content = strings.TrimSuffix(string(raw), "\n")
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
