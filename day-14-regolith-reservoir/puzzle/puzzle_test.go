package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	blockedLocations, highestY := parseInput(input())
	assert.Equal(t, 20, len(blockedLocations))
	assert.Equal(t, 9, highestY)
}

func input() []string {
	return []string{
		"498,4 -> 498,6 -> 496,6",
		"503,4 -> 502,4 -> 502,9 -> 494,9",
	}
}
