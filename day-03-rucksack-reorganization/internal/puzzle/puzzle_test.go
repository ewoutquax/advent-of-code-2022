package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	assert.Equal(t, 16, calculateScore("vJrwpWtwJgWrhcsFMMfFFhFp"))
	assert.Equal(t, 38, calculateScore("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"))
	assert.Equal(t, 42, calculateScore("PmmdzqPrVvPwwTWBwg"))
	assert.Equal(t, 22, calculateScore("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"))
	assert.Equal(t, 20, calculateScore("ttgJtRGJQctTZtZT"))
	assert.Equal(t, 19, calculateScore("CrZsJsPPZsGzwwsLwLmpwMDw"))
}

func TestPart2Examples(t *testing.T) {
	inputs := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	assert.Equal(t, 70, CalculateTotalBadgeScore(inputs))
}
