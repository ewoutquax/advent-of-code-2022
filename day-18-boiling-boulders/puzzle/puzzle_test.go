package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	boulders := linkNeighbours(parseInput(input1()))
	boulder := boulders[601]

	assert.Equal(t, 2, len(boulders))
	assert.Equal(t, 1, boulder.coordinate.x)
	assert.Equal(t, 1, boulder.coordinate.y)
	assert.Equal(t, 1, boulder.coordinate.z)
	assert.Equal(t, 1, len(boulder.neighbours))
}

func TestCountExposedSides(t *testing.T) {
	assert.Equal(t, 10, CountExposedSides(input1()))
	assert.Equal(t, 64, CountExposedSides(input2()))
}

func TestCountSidesExposedToFreeAir(t *testing.T) {
	assert.Equal(t, 10, CountSidesExposedToFreeAir(input1()))
	assert.Equal(t, 58, CountSidesExposedToFreeAir(input2()))
}

func input1() []string {
	return []string{
		"1,1,1",
		"2,1,1",
	}
}

func input2() []string {
	return []string{
		"2,2,2",
		"1,2,2",
		"3,2,2",
		"2,1,2",
		"2,3,2",
		"2,2,1",
		"2,2,3",
		"2,2,4",
		"2,2,6",
		"1,2,5",
		"3,2,5",
		"2,1,5",
		"2,3,5",
	}
}
