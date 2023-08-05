package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	assert.Equal(t, 7, FindFirstMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4))
	assert.Equal(t, 5, FindFirstMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 4))
	assert.Equal(t, 10, FindFirstMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4))
}

func TestPart2Examples(t *testing.T) {
	assert.Equal(t, 19, FindFirstMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))
	assert.Equal(t, 23, FindFirstMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 14))
}
