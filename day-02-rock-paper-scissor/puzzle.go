package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1([]string{}))
	fmt.Println("Result of part-2: ", solvePart2([]string{}))
}

func solvePart1(lines []string) (totalScore int) {
	if len(lines) == 0 {
		lines = read_file_as_lines()
	}

	for _, line := range lines {
		choices := strings.Split(line, " ")
		totalScore += calculateScore(choices[0], choices[1])
	}

	return
}

func solvePart2(lines []string) (totalScore int) {
	if len(lines) == 0 {
		lines = read_file_as_lines()
	}

	for _, line := range lines {
		choices := strings.Split(line, " ")
		totalScore += calculateScorePart2(choices[0], choices[1])
	}

	return
}

func calculateScore(his string, mine string) int {
	conv_his := convertType(his)
	conv_mine := convertType(mine)

	return scoreForType(conv_mine) + scoreForOutcome(conv_his, conv_mine)
}

func calculateScorePart2(his string, requiredOutcome string) int {
	conv_his := convertType(his)
	var mine string = determineMine(conv_his, requiredOutcome)

	return scoreForType(mine) + scoreForOutcome(conv_his, mine)
}

func determineMine(his string, requiredOutcome string) (mine string) {
	toWin := map[string]string{
		"rock":    "paper",
		"paper":   "scissor",
		"scissor": "rock",
	}

	toLose := map[string]string{
		"rock":    "scissor",
		"paper":   "rock",
		"scissor": "paper",
	}

	if requiredOutcome == "X" {
		// Lose turn
		mine = toLose[his]
	} else if requiredOutcome == "Z" {
		// Win turn
		mine = toWin[his]
	} else {
		// draw
		mine = his
	}

	return
}

func convertType(input string) string {
	return map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissor",
		"X": "rock",
		"Y": "paper",
		"Z": "scissor",
	}[input]
}

func scoreForType(choice string) int {
	return map[string]int{
		"rock":    1,
		"paper":   2,
		"scissor": 3,
	}[choice]
}

func scoreForOutcome(his string, mine string) (score int) {
	if his == mine {
		score = 3
	} else if his == "rock" && mine == "paper" ||
		his == "paper" && mine == "scissor" ||
		his == "scissor" && mine == "rock" {
		score = 6
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
