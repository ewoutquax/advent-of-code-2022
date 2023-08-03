package main_test

import (
	"testing"

	. "aoc.com/2022/day-02/cmd/part-1"
	"github.com/stretchr/testify/assert"
)

func TestSolvePuzzlePart(t *testing.T) {
	assert.Equal(t, "11150", SolvePuzzlePart())
}
