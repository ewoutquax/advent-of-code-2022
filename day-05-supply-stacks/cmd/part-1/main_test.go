package main_test

import (
	"testing"

	. "aoc.com/2022/day-05/cmd/part-1"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, "HBTMTBSDC", SolvePuzzlePart())
}
