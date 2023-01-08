package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	list := parseInput(input(), 1)

	assert.Equal(t, 7, len(list))
	assert.Equal(t, 1, list[0].value)
	assert.Equal(t, 2, list[0].next.value)
	assert.Equal(t, 4, list[0].prev.value)
	assert.Equal(t, 1, list[0].prev.next.value)
}

func TestMixNumbers(t *testing.T) {
	var current *Number

	list := mixNumbers(parseInput(input(), 1))

	for idx := 0; idx < len(list); idx++ {
		if list[idx].value == 0 {
			current = list[idx]
		}
	}

	assert.Equal(t, 0, current.value)
	assert.Equal(t, 3, current.next.value)
	assert.Equal(t, -2, current.next.next.value)
	assert.Equal(t, 1, current.next.next.next.value)
	assert.Equal(t, 2, current.next.next.next.next.value)
	assert.Equal(t, -3, current.next.next.next.next.next.value)
	assert.Equal(t, 4, current.next.next.next.next.next.next.value)
}

func TestGrooveHash(t *testing.T) {
	assert.Equal(t, 3, GrooveHash(input()))
}

func TestGrooveHashHard(t *testing.T) {
	assert.Equal(t, 1623178306, GrooveHashHard(input()))
}

func input() []int {
	return []int{
		1,
		2,
		-3,
		3,
		-2,
		0,
		4,
	}
}
