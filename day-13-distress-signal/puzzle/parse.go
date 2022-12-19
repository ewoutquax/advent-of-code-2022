package puzzle

import (
	"fmt"
	"strconv"

	"aoc.com/2022/day-13/utils"
)

type Node struct {
	isInteger bool
	value     int
	childNode *Node
	nextNode  *Node
}

func printNode(node *Node) (out string) {
	if node.isInteger {
		out += strconv.Itoa(node.value)
	} else {
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
	var currentNode Node
	var index int

	for index = startIndex; index < len(input)-1 && string(input[index]) != "]"; index += 0 {
		currentChar = string(input[index])
		fmt.Println("1: currentChar '", currentChar, "' at index:", index)

		switch currentChar {
		case "[":
			fmt.Println("Jump into subloop")
			childNode, newIndex := parseInputRecursive(input, index+1)
			index = newIndex
			fmt.Println("Returning from subloop")
			currentNode.isInteger = false
			currentNode.childNode = &childNode
			return currentNode, index

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
			fmt.Print("Parsed number ", number)

			if string(input[index]) == "," {
				fmt.Println("; found ',': continuing via recursion from index:", index)
				nextNode, newIndex := parseInputRecursive(input, index)
				currentNode.nextNode = &nextNode
				index = newIndex
			} else {
				fmt.Print("\n")
			}
			return currentNode, newIndex
		}
	}

	fmt.Println("Emptying the nextNode")
	currentNode.nextNode = nil
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
