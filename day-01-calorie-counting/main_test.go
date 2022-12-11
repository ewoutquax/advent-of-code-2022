package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Solve(t *testing.T) {
	assert.Equal(t, 71502, solvePart1())
}

func TestPart2Solve(t *testing.T) {
	assert.Equal(t, 208191, solvePart2())
}
