package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Solve(t *testing.T) {
	assert.Equal(t, 49288254556480, solvePart1())
}

func TestPart2Solve(t *testing.T) {
	assert.Equal(t, 3558714869436, solvePart2())
}
