package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrop4Block(t *testing.T) {
	cave := stackBlocks(input(), 4)

	var expectedHeights RowHeights
	expectedHeights.row = map[int]int{
		1: 4,
		2: 4,
		3: 6,
		4: 4,
		5: 7,
		6: 1,
	}

	drawCave(cave)

	assert.Equal(t, 7, cave.heighestRock)
	assert.Equal(t, expectedHeights, cave.heights)
}

func TestHeightStackedBlocks(t *testing.T) {
	assert.Equal(t, 3068, HeightStackedBlocks(input(), 2022))
}

func input() string {
	return ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
}
