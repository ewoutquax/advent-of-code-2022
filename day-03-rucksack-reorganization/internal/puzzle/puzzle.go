package puzzle

import (
	"strings"
)

const SCORES string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func CalculateTotalScore(lines []string) (score int) {
	for _, line := range lines {
		score += calculateScore(line)
	}

	return
}

func CalculateTotalBadgeScore(lines []string) (score int) {
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
