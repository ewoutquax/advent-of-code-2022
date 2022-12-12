package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Solve(t *testing.T) {
	assert.Equal(t, 584, solvePart1())
}

func TestPart2Solve(t *testing.T) {
	assert.Equal(t, 933, solvePart2())
}
