package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingSingleNumber(t *testing.T) {
	resultNode := ParseInput("[9]")
	assert.Equal(t, false, resultNode.isInteger)

	resultSubNode := resultNode.childNode
	assert.Equal(t, 9, resultSubNode.value)
}

func TestParsingTwoNumbers(t *testing.T) {
	resultNode := ParseInput("[9,[10]]")
	assert.Equal(t, false, resultNode.isInteger)

	firstSubNode := resultNode.childNode
	assert.Equal(t, 9, firstSubNode.value)

	secondSubNode := firstSubNode.nextNode
	assert.Equal(t, false, secondSubNode.isInteger)

	firstSubSubNode := secondSubNode.childNode
	assert.Equal(t, true, firstSubSubNode.isInteger)
	assert.Equal(t, 10, firstSubSubNode.value)
}
