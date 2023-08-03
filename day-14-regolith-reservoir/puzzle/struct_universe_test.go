package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFillWithSandUntilFloor(t *testing.T) {
	blockedPositions, maxY := parseInput(input())
	universe := Universe{
		blockedPositions: blockedPositions,
		floorY:           maxY + 2,
	}

	universe.fillWithSandUntilFloor()

	assert.Equal(t, 24, universe.nrSands)
}

func TestFillWithSandUntilTop(t *testing.T) {
	blockedPositions, maxY := parseInput(input())
	universe := Universe{
		blockedPositions: blockedPositions,
		floorY:           maxY + 2,
	}
	universe.fillWithSandUntilTop()

	assert.Equal(t, 93, universe.nrSands)
}
