package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Solve(t *testing.T) {
	assert.Equal(t, 6387, solvePart1())
}

func TestPart2Solve(t *testing.T) {
	assert.Equal(t, 2455057187825, solvePart2())
}
