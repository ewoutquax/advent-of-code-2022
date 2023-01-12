package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	topleft, steps := parseInput(input())
	expectedTopLeftPosition := Position{x: 8, y: 0}
	expectedTopLeftLeftPosition := Position{x: 11, y: 0}

	assert.Equal(t, expectedTopLeftPosition, topleft.position)
	assert.False(t, topleft.isWall)

	assert.Equal(t, expectedTopLeftLeftPosition, topleft.left.position)
	assert.True(t, topleft.left.isWall)

	assert.Equal(t, "10R5L5R10L4R5L5", steps)
}

func TestInitPlayer(t *testing.T) {
	topleft, _ := parseInput(input())
	player := initPlayer(topleft)
	expectedPlayerPosition := Position{x: 8, y: 0}

	assert.Equal(t, expectedPlayerPosition, player.location.position)
	assert.Equal(t, Right, player.facing)
}

func TestFollowSteps(t *testing.T) {
	topleft, steps := parseInput(input())
	player := followSteps(steps, initPlayer(topleft))
	expectedPlayerPosition := Position{x: 7, y: 5}

	assert.Equal(t, expectedPlayerPosition, player.location.position)
	assert.Equal(t, 6032, player.password())
}

func input() [][]string {
	return [][]string{{
		"        ...#",
		"        .#..",
		"        #...",
		"        ....",
		"...#.......#",
		"........#...",
		"..#....#....",
		"..........#.",
		"        ...#....",
		"        .....#..",
		"        .#......",
		"        ......#.",
	}, {
		"10R5L5R10L4R5L5",
	}}
}
