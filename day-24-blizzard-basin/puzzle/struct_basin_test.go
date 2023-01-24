package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// After moving 4 times in this 4x4 basin, all blizzards should be back
// at their original positions
func TestMoveBlizzards2(t *testing.T) {
	var basin Basin = parseInput(inputBlizzard())

	origPositionKeys := basin.blizzardLocations2()
	basin.moveBlizzards()
	basin.moveBlizzards()
	basin.moveBlizzards()
	basin.moveBlizzards()
	newPositionKeys := basin.blizzardLocations2()

	assert.Equal(t, origPositionKeys, newPositionKeys)
}

func inputBlizzard() []string {
	return []string{
		"#.####",
		"#..<.#",
		"#v...#",
		"#...^#",
		"#.>..#",
		"####.#",
	}
}
