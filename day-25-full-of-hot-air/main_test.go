package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Solve(t *testing.T) {
	assert.Equal(t, "2=01-0-2-0=-0==-1=01", solvePart1())
}
