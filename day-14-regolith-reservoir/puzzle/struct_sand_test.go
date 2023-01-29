package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFall(t *testing.T) {
	blockedPositions, maxY := parseInput(input())

	universe := Universe{
		blockedPositions: blockedPositions,
		nrSands:          0,
		floorY:           maxY + 2,
	}

	sand := spawnSand()
	sand.fallUntilRest(&universe)

	assert.Equal(t, 500, sand.position.x)
	assert.Equal(t, 8, sand.position.y)
}

func TestCanFallIntoTheVoid(t *testing.T) {
	universe := Universe{
		blockedPositions: make(map[string]bool),
		nrSands:          0,
		floorY:           201,
	}

	sand := spawnSand()
	sand.fallUntilRest(&universe)

	assert.Equal(t, 500, sand.position.x)
	assert.Equal(t, 200, sand.position.y)
	assert.True(t, sand.isOnFloor(&universe))
}
