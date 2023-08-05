package main_test

import (
	"testing"

	. "aoc.com/2022/day-05/cmd/part-2"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, "PQTJRSHWS", SolvePuzzlePart())
}
