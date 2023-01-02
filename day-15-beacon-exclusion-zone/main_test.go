package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Solve(t *testing.T) {
	assert.Equal(t, 5176944, solvePart1(2000000))
}

func TestPart2Solve(t *testing.T) {
	assert.Equal(t, 13350458933732, solvePart2())
}
