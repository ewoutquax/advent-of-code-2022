package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Examples(t *testing.T) {
	count := NumberOfLocationsOfKnot(inputPart1(), 1)
	assert.Equal(t, 13, count)
}

func TestPart2Examples(t *testing.T) {
	count := NumberOfLocationsOfKnot(inputPart2(), 9)
	assert.Equal(t, 36, count)
}

func inputPart1() []string {
	return []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}
}

func inputPart2() []string {
	return []string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	}
}
