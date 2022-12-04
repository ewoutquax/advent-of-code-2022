package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	assert.Equal(t, 8, calculateScore("A", "Y"))
	assert.Equal(t, 1, calculateScore("B", "X"))
	assert.Equal(t, 6, calculateScore("C", "Z"))

	input := []string{
		"A Y",
		"B X",
		"C Z",
	}

	assert.Equal(t, 15, solvePart1(input))
}

func TestPart1Solve(t *testing.T) {
	assert.Equal(t, 11150, solvePart1([]string{}))
}

func TestPart2Examples(t *testing.T) {
	assert.Equal(t, 4, calculateScorePart2("A", "Y"))
	assert.Equal(t, 1, calculateScorePart2("B", "X"))
	assert.Equal(t, 7, calculateScorePart2("C", "Z"))

	input := []string{
		"A Y",
		"B X",
		"C Z",
	}

	assert.Equal(t, 12, solvePart2(input))
}

func TestPart2Solve(t *testing.T) {
	assert.Equal(t, 208191, solvePart2([]string{}))
}
