package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingSingleNumber(t *testing.T) {
	node_0 := ParseInput("[9]")
	assert.False(t, node_0.isInteger)
	assert.True(t, node_0.isList)

	node_0_0 := node_0.childNode
	assert.Equal(t, 9, node_0_0.value)
	assert.True(t, node_0_0.isInteger)
	assert.False(t, node_0_0.isList)
}

func TestParsingTwoNumbers(t *testing.T) {
	node_0 := ParseInput("[9,10]")
	assert.False(t, node_0.isInteger)
	assert.True(t, node_0.isList)

	node_0_0 := node_0.childNode
	assert.Equal(t, 9, node_0_0.value)
	assert.True(t, node_0_0.isInteger)
	assert.False(t, node_0_0.isList)

	node_0_1 := node_0_0.nextNode
	assert.True(t, node_0_1.isInteger)
	assert.False(t, node_0_1.isList)
}

func TestParsingSubloop(t *testing.T) {
	resultNode := ParseInput("[9,[10,345]]")
	fmt.Println("resultNode:", printNode(&resultNode))
	assert.Equal(t, false, resultNode.isInteger)

	node_1_1 := resultNode.childNode
	assert.Equal(t, 9, node_1_1.value)

	node_1_2 := node_1_1.nextNode
	assert.Equal(t, false, node_1_2.isInteger)

	node_2_1 := node_1_2.childNode
	assert.Equal(t, true, node_2_1.isInteger)
	assert.Equal(t, 10, node_2_1.value)

	node_2_2 := node_2_1.nextNode
	assert.Equal(t, true, node_2_2.isInteger)
	assert.Equal(t, 345, node_2_2.value)
}
