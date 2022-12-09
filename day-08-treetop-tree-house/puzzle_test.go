package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	trees := parseInput(input())

	topleft := trees[0]
	bottomright := trees[404]

	assert.True(t, isVisible(topleft))
	assert.True(t, isVisible(bottomright))
	assert.False(t, isVisible(trees[303]))

	assert.Equal(t, 25, len(trees))
	assert.Equal(t, 3, topleft.height)
	assert.Equal(t, true, topleft.direction[NORTH].is_leaf)
	assert.Equal(t, true, topleft.direction[WEST].is_leaf)
	assert.Equal(t, false, topleft.direction[SOUTH].is_leaf)
	assert.Equal(t, false, topleft.direction[EAST].is_leaf)
	assert.Equal(t, false, bottomright.direction[NORTH].is_leaf)
	assert.Equal(t, false, bottomright.direction[WEST].is_leaf)
	assert.Equal(t, true, bottomright.direction[SOUTH].is_leaf)
	assert.Equal(t, true, bottomright.direction[EAST].is_leaf)

	nrVisibleTrees := countVisibleTrees(trees)
	assert.Equal(t, 21, nrVisibleTrees)
}

func TestPart2Examples(t *testing.T) {
	trees := parseInput(input())

	score := maxScenicScore(trees)
	assert.Equal(t, 8, score)
}

func TestSolvePart1(t *testing.T) {
	count := solvePart1([]string{})
	assert.Equal(t, 1708, count)
}

func TestSolvePart2(t *testing.T) {
	maxScenicScore := solvePart2([]string{})
	assert.Equal(t, 504000, maxScenicScore)
}

func input() []string {
	return []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}
}
