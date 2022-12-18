package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluations(t *testing.T) {
	var left, right Node
	left = ParseInput("[1,1,3,1,1]")
	right = ParseInput("[1,1,5,1,1]")
	assert.True(t, isValid(left, right))

	left = ParseInput("[[1],[2,3,4]]")
	right = ParseInput("[[1],4]")
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
