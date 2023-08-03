package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	basin := parseInput(input())

	assert.Equal(t, 0, basin.nrStep)
	assert.Equal(t, "1,0", basin.players[0].position.toKey())
	assert.Equal(t, 19, len(basin.blizzards))
	assert.Equal(t, 22, len(basin.wallLocations))
	assert.False(t, basin.playerReachedGoal)
}

func TestMoveBlizzards(t *testing.T) {
	var basin Basin = parseInput(input())
	basin.moveBlizzards()
	positionKeys := basin.blizzardLocations()

	assert.Equal(t, 19, len(basin.blizzards))
	assert.Equal(t, 14, len(positionKeys))
	assert.True(t, positionKeys["2,1"])
	assert.True(t, positionKeys["1,2"])
}

func TestMoveToEnd(t *testing.T) {
	assert.Equal(t, 18, LeastStepsToExit(input()))
}

func TestMoveToEndBeginningAndEndAgain(t *testing.T) {
	steps := LeastStepsToExitStartAndExit(input())

	assert.Equal(t, 54, steps)
}

func input() []string {
	return []string{
		"#.######",
		"#>>.<^<#",
		"#.<..<<#",
		"#>v.><>#",
		"#<^v^^>#",
		"######.#",
	}
}
