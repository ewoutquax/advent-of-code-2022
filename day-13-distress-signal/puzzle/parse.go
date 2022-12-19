package puzzle

import (
	"fmt"
	"strconv"

	"aoc.com/2022/day-13/utils"
)

type Node struct {
	isInteger bool
	isList    bool
	value     int
	childNode *Node
	nextNode  *Node
}

func printNode(node *Node) (out string) {
	if node.isInteger {
		out += strconv.Itoa(node.value)
	}

	if node.isList {
		out += "[" + printNode(node.childNode) + "]"
	}

	if node.nextNode != nil {
		out += "," + printNode(node.nextNode)
	}

	return
}

func ParseInput(input string) Node {
	rootNode, _ := parseInputRecursive(input, 0)

	return rootNode
}

func parseInputRecursive(input string, startIndex int) (Node, int) {
	var currentChar string
	var index int
	var currentNode = Node{
		isInteger: false,
		isList:    false,
		childNode: nil,
		nextNode:  nil,
	}

	for index = startIndex; index < len(input)-1 && string(input[index]) != "]"; index += 0 {
		currentChar = string(input[index])
		fmt.Println("1: currentChar '", currentChar, "' at index:", index)

		switch currentChar {
		case "[":
			fmt.Println("Jump into subloop")
			childNode, newIndex := parseInputRecursive(input, index+1)
			index = newIndex
			fmt.Println("Returning from subloop; continuing on index:", index)
			currentNode.isList = true
			currentNode.childNode = &childNode

		case ",":
			fmt.Println("; found ',': continuing via recursion from index:", index)
			nextNode, newIndex := parseInputRecursive(input, index+1)
			currentNode.nextNode = &nextNode
			index = newIndex
			return currentNode, newIndex

		case "]":
			continue

		default:
			fmt.Println("Going to parse a number")
			number, newIndex := parseNumber(input, index)
			index = newIndex
			currentNode.isInteger = true
			currentNode.value = number
			fmt.Println("Parsed number ", number)
		}
	}

	fmt.Println("End of loop; return to parent")
	return currentNode, index + 1
}

func parseNumber(input string, startIndex int) (int, int) {
	var index int

	var goOn bool = true
	for index = startIndex; goOn; index += 1 {
		nextChar := string(input[index+1])
		goOn = nextChar != "," && nextChar != "]"
	}
	foundNumber := utils.ConvStrToI(input[startIndex:index])

	fmt.Println("Found number from pos", startIndex, "till", index-1)

	return foundNumber, index
}
