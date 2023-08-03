package puzzle

import (
	"strings"
)

func CalculateTotalScoreForRounds(lines []string) (totalScore int) {
	for _, line := range lines {
		choices := strings.Split(line, " ")
		totalScore += calculateScoreForRound(choices[0], choices[1])
	}

	return
}

func CalculateTotalScoreByRequiredOutcomes(lines []string) (totalScore int) {
	for _, line := range lines {
		choices := strings.Split(line, " ")
		totalScore += calculateScoreByRequiredOutcome(choices[0], choices[1])
	}

	return
}

func calculateScoreForRound(hisRaw string, mineRaw string) int {
	his := convertType(hisRaw)
	mine := convertType(mineRaw)

	return mine.value() + mine.scoreRound(his)
}

func calculateScoreByRequiredOutcome(hisRaw string, requiredOutcome string) int {
	his := convertType(hisRaw)
	mine := his.determineForOutcome(requiredOutcome)

	return mine.value() + mine.scoreRound(his)
}

func convertType(input string) Item {
	return map[string]Item{
		"A": Rock,
		"B": Paper,
		"C": Scissor,
		"X": Rock,
		"Y": Paper,
		"Z": Scissor,
	}[input]
}
