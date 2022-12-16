package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	out := parseInput("[9]")
	expectedNode := Node{
		isInteger: true,
		value:     9,
	}
	assert.Equal(t, expectedNode, out)
}

func TestEvaluations(t *testing.T) {
	left := "[1,1,3,1,1]"
	right := "[1,1,5,1,1]"

	assert.True(t, isValid(left, right))
}

func input() [][]string {
	return [][]string{{
		"[1,1,3,1,1]",
		"[1,1,5,1,1]",
	}, {
		"[[1],[2,3,4]]",
		"[[1],4]",
	}, {
		"[9]",
		"[[8,7,6]]",
	}, {
		"[[4,4],4,4]",
		"[[4,4],4,4,4]",
	}, {
		"[7,7,7,7]",
		"[7,7,7]",
	}, {
		"[]",
		"[3]",
	}, {
		"[[[]]]",
		"[[]]",
	}, {
		"[1,[2,[3,[4,[5,6,7]]]],8,9]",
		"[1,[2,[3,[4,[5,6,0]]]],8,9]",
	}}
}
