package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	assert.Equal(t, 7, findFirstMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4))
	assert.Equal(t, 5, findFirstMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 4))
	assert.Equal(t, 10, findFirstMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4))
}

func TestPart2Examples(t *testing.T) {
	assert.Equal(t, 19, findFirstMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))
	assert.Equal(t, 23, findFirstMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 14))
}

func TestPart1Solve(t *testing.T) {
	assert.Equal(t, 1140, solvePart1("", 4))
}

func TestPart2Solve(t *testing.T) {
	assert.Equal(t, 3495, solvePart1("", 14))
}

