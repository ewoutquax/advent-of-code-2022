package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	assert.Equal(t, 16, calculateScore("vJrwpWtwJgWrhcsFMMfFFhFp"))
	assert.Equal(t, 38, calculateScore("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"))
	assert.Equal(t, 42, calculateScore("PmmdzqPrVvPwwTWBwg"))
	assert.Equal(t, 22, calculateScore("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"))
	assert.Equal(t, 20, calculateScore("ttgJtRGJQctTZtZT"))
	assert.Equal(t, 19, calculateScore("CrZsJsPPZsGzwwsLwLmpwMDw"))
}

func TestPart1Solve(t *testing.T) {
	assert.Equal(t, 8085, solvePart1([]string{}))
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

	assert.Equal(t, 70, solvePart2(inputs))
}

func TestPart2Solve(t *testing.T) {
	assert.Equal(t, 2515, solvePart2([]string{}))
}
