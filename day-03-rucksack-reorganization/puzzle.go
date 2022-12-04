package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const SCORES string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	fmt.Println("Result of part-1: ", solvePart1([]string{}))
	fmt.Println("Result of part-2: ", solvePart2([]string{}))
}

func solvePart1(lines []string) (score int) {
	if len(lines) == 0 {
		lines = read_file_as_lines()
	}

	for _, line := range lines {
		score += calculateScore(line)
	}

	return
}

func solvePart2(lines []string) (score int) {
	if len(lines) == 0 {
		lines = read_file_as_lines()
	}

	var badge string
	var groupId int
	group := make([]string, 3)

	groupId = 0
	for _, line := range lines {
		group[groupId] = line
		groupId += 1

		if groupId == 3 {
			group1 := strings.Split(group[0], "")
			group2 := strings.Split(group[1], "")
			group3 := strings.Split(group[2], "")

			badge = string(findIntersections(group1, findIntersections(group2, group3))[0])
			score += charScore(badge)

			groupId = 0
		}
	}

	return
}

func calculateScore(content string) int {
	list := strings.Split(content, "")

	left := list[0 : len(list)/2]
	right := list[len(list)/2 : len(list)]

	matching := string(findIntersections(left, right)[0])
	score := charScore(matching)

	return score
}

func charScore(input string) int {
	for score, char := range strings.Split(SCORES, "") {
		if char == input {
			return score + 1
		}
	}

	return 0
}

func findIntersections(left []string, right []string) (matches []string) {
	for _, char_left := range left {
		for _, char_right := range right {
			if char_left == char_right {
				matches = append(matches, char_left)
			}
		}
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
