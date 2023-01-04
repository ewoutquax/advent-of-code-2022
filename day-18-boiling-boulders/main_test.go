package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Solve(t *testing.T) {
	assert.Equal(t, 4418, solvePart1())
}

func TestPart2Solve(t *testing.T) {
	assert.Equal(t, 2486, solvePart2())
}
