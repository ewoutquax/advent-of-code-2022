package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	out := parseInput(input())
	assert.Equal(t, 40, len(out))
}

func TestFindPath(t *testing.T) {
	locations := parseInput(input())
	pathLength := findPath(locations)
	assert.Equal(t, 31, pathLength)
}

func TestFindPathDown(t *testing.T) {
	locations := parseInput(input())
	pathLength := findPathDown(locations)
	assert.Equal(t, 29, pathLength)
}

func input() []string {
	return []string{"Sabqponm",
		"abcryxxl",
		"accszExk",
		"acctuvwj",
		"abdefghi",
	}
}
