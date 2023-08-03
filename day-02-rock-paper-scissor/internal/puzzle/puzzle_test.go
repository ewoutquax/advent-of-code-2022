package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	assert.Equal(t, 8, calculateScoreForRound("A", "Y"))
	assert.Equal(t, 1, calculateScoreForRound("B", "X"))
	assert.Equal(t, 6, calculateScoreForRound("C", "Z"))

	input := []string{
		"A Y",
		"B X",
		"C Z",
	}

	assert.Equal(t, 15, CalculateTotalScoreForRounds(input))
}

func TestPart2Examples(t *testing.T) {
	assert.Equal(t, 4, calculateScoreByRequiredOutcome("A", "Y"))
	assert.Equal(t, 1, calculateScoreByRequiredOutcome("B", "X"))
	assert.Equal(t, 7, calculateScoreByRequiredOutcome("C", "Z"))

	input := []string{
		"A Y",
		"B X",
		"C Z",
	}

	assert.Equal(t, 12, CalculateTotalScoreByRequiredOutcomes(input))
}
