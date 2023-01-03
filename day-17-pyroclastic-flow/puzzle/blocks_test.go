package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBlocks(t *testing.T) {
	assert.Equal(t, 5, len(parseBlocks()))
}

func TestParseBlock2(t *testing.T) {
	block := parseBlock(block2())

	expectedRocks := []RelativePosition{
		{position: Position{x: 1, y: 0}},
		{position: Position{x: 0, y: -1}},
		{position: Position{x: 1, y: -1}},
		{position: Position{x: 2, y: -1}},
		{position: Position{x: 1, y: -2}},
	}

	assert.Equal(t, 3, block.height)
	assert.Equal(t, expectedRocks, block.rocks)
}
