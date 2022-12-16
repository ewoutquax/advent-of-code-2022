package puzzle

import (
	"aoc.com/2022/day-13/utils"
)

type Node struct {
	isInteger bool
	value     int
	childNode *Node
	nextNode  *Node
}

func parseInput(input string) (rootNode Node) {
	rootNode.isInteger = false

	childNode, _ := parseInputRecursive(input, 1)
	rootNode.childNode = &childNode

	return
}

func parseInputRecursive(input string, startIndex int) (Node, int) {
	currentChar := string(input[startIndex])
	var currentNode Node

	// if currentChar == "[" {
	// 	node.isInteger = false
	// 	childNode, index := parseInputRecursive(nextNode, input, startIndex+1)
	// 	node.childNode = &childNode
	// }

	if currentChar != "]" {
		number, index := parseNumber(input, startIndex)
		currentNode.isInteger = true
		currentNode.value = number

		nextNode, index := parseInputRecursive(input, index)
		currentNode.nextNode = &nextNode
	}

	// if currentChar == "]" {
	// 	return node, startIndex + 1
	// }
}

func parseNumber(input string, startIndex int) (int, int) {
	var goOn bool = true
	var index int

	for index = startIndex; goOn; index += 1 {
		currentChar := string(input[index])
		goOn = currentChar == "," || currentChar == "]"
	}

	return utils.ConvStrToI(input[startIndex:index]), index
}
