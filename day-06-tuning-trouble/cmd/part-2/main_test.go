package main_test

import (
	"testing"

	. "aoc.com/2022/day-06/cmd/part-2"
	"github.com/stretchr/testify/assert"
)

func TestSolvePuzzlePart(t *testing.T) {
	assert.Equal(t, 3495, SolvePuzzlePart())
}
