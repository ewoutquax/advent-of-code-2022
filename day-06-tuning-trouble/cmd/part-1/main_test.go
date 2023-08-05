package main_test

import (
	"testing"

	. "aoc.com/2022/day-06/cmd/part-1"
	"github.com/stretchr/testify/assert"
)

func TestSolvePuzzlePart(t *testing.T) {
	assert.Equal(t, 1140, SolvePuzzlePart())
}
