package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	assert.True(t, completeOverlap("6-6,4-6"))
	assert.True(t, completeOverlap("2-8,3-7"))
	assert.False(t, completeOverlap("2-4,6-8"))
	assert.False(t, completeOverlap("2-3,4-5"))
	assert.False(t, completeOverlap("5-7,7-9"))
	assert.False(t, completeOverlap("2-6,4-8"))
}

func TestPart1Solve(t *testing.T) {
	assert.Equal(t, 584, solvePart1([]string{}))
}

func TestPart2Examples(t *testing.T) {
	assert.False(t, hasOverlap("2-4,6-8"))
	assert.False(t, hasOverlap("2-3,4-5"))
	assert.True(t, hasOverlap("5-7,7-9"))
	assert.True(t, hasOverlap("2-8,3-7"))
	assert.True(t, hasOverlap("6-6,4-6"))
	assert.True(t, hasOverlap("2-6,4-8"))
}

func TestPart2Solve(t *testing.T) {
	assert.Equal(t, 933, solvePart2([]string{}))
}
