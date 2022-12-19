package puzzle

import (
	"fmt"

	"aoc.com/2022/day-13/utils"
)

type Node struct {
	isInteger bool
	value     int
	childNode *Node
	nextNode  *Node
}

func ParseInput(input string) Node {
	rootNode, _ := parseInputRecursive(input, 0)

	return rootNode
}

func parseInputRecursive(input string, startIndex int) (Node, int) {
	var currentChar string
	var currentNode Node
	var index int

	for index = startIndex; index < len(input)-1 && string(input[index]) != "]"; index += 0 {
		currentChar = string(input[index])
		fmt.Println("1: currentChar '", currentChar, "' at index:", index)

		switch currentChar {
		case "[":
			fmt.Println("Jump into subloop")
			currentNode.isInteger = false
			childNode, tempIndex := parseInputRecursive(input, index+1)
			fmt.Println("Returning from subloop")
			currentNode.childNode = &childNode
			index = tempIndex

		case ",":
			index += 1

		case "]":
			continue

		default:
			fmt.Println("Going to parse a number")
			number, newIndex := parseNumber(input, index)
			index = newIndex
			currentNode.isInteger = true
			currentNode.value = number

			fmt.Println("Parsed number ", number, "; continuing with index:", index)
		}
	}

	return currentNode, index + 1
}

func parseNumber(input string, startIndex int) (int, int) {
	var index int

	var goOn bool = true
	for index = startIndex; goOn; index += 1 {
		currentChar := string(input[index])
		fmt.Println("4: currentChar '", currentChar, "' at index:", index)
		goOn = currentChar != "," && currentChar != "]"
	}
	index -= 1

	fmt.Println("Found number from pos", startIndex, "till", index-1)

	foundNumber := utils.ConvStrToI(input[startIndex:index])
	return foundNumber, index
}
