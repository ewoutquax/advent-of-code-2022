package puzzle

import (
	"fmt"
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
	resultNode := ParseInput("[9,10]")
	fmt.Println("resultNode:", printNode(&resultNode))
	assert.Equal(t, false, resultNode.isInteger)

	node_0_0 := resultNode.childNode
	assert.Equal(t, 9, node_0_0.value)

	node_0_1 := node_0_0.nextNode
	assert.Equal(t, true, node_0_1.isInteger)
	assert.Equal(t, 10, node_0_1.value)
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
