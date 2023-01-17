package puzzle

import (
	"testing"

	"github.com/ewoutquax/advent-of-code-2022/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	elves := parseInput(input())

	assert.Equal(t, 5, len(elves))
	assert.Equal(t, 2, elves["2,1"].position.x)
	assert.Equal(t, 1, elves["2,1"].position.y)
}

func TestDetermineMoveToPosition(t *testing.T) {
	elves := parseInput(input())

	setMoveToPosition(0, &elves)

	assert.Equal(t, 2, elves["2,1"].moveTo.x)
	assert.Equal(t, 0, elves["2,1"].moveTo.y)
}

func TestFinalizePositions(t *testing.T) {
	elves := parseInput(input())

	setMoveToPosition(0, &elves)
	movedElves := finalizePositions(elves)

	var locations []string
	for key := range movedElves {
		locations = append(locations, key)
	}

	assert.True(t, utils.InSlice(locations, "2,0"))
	assert.True(t, utils.InSlice(locations, "3,0"))
	assert.True(t, utils.InSlice(locations, "2,2"))
	assert.True(t, utils.InSlice(locations, "2,4"))
	assert.True(t, utils.InSlice(locations, "3,3"))
}

func TestMultipleRounds(t *testing.T) {
	elves := parseInput(input())

	movedElves := moveMultipleRounds(10, elves)

	var locations []string
	for key := range movedElves {
		locations = append(locations, key)
	}

	assert.True(t, utils.InSlice(locations, "2,0"))
	assert.True(t, utils.InSlice(locations, "4,1"))
	assert.True(t, utils.InSlice(locations, "0,2"))
	assert.True(t, utils.InSlice(locations, "4,3"))
	assert.True(t, utils.InSlice(locations, "2,5"))
}

func TestCountGroundTiles(t *testing.T) {
	elves := parseInput(input2())

	assert.Equal(t, 110, countEmptyGroundTiles(elves))
}

func TestFindLastRoundWithMovement(t *testing.T) {
	round := FirstRoundWithoutMovement(input3())

	assert.Equal(t, 20, round)
}

func input() []string {
	return []string{
		".....",
		"..##.",
		"..#..",
		".....",
		"..##.",
		".....",
	}
}

func input2() []string {
	return []string{
		"......#.....",
		"..........#.",
		".#.#..#.....",
		".....#......",
		"..#.....#..#",
		"#......##...",
		"....##......",
		".#........#.",
		"...#.#..#...",
		"............",
		"...#..#..#..",
	}
}

func input3() []string {
	return []string{
		"....#..",
		"..###.#",
		"#...#.#",
		".#...##",
		"#.###..",
		"##.#.##",
		".#..#..",
	}
}
