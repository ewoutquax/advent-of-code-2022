package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const UP string = "U"
const LEFT string = "L"
const DOWN string = "D"
const RIGHT string = "R"

type knot struct {
	x         int
	y         int
	locations []string
}

func main() {
	fmt.Println("Result of part-1: ", solvePart1([]string{}))
	fmt.Println("Result of part-2: ", solvePart2([]string{}))
}

func solvePart1(instructions []string) int {
	if len(instructions) == 0 {
		instructions = readFileAsLines()
	}

	locations := execInstructions(instructions)

	return len(unique(locations))
}

func solvePart2(instructions []string) int {
	if len(instructions) == 0 {
		instructions = readFileAsLines()
	}

	knots := execInstructionsPart2(instructions)

	return len(unique(knots[9].locations))
}

func execInstructions(instructions []string) []string {
	var head knot
	var tail knot

	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		direction := parts[0]
		repeats := convStrToI(parts[1])

		for repeater := 0; repeater < repeats; repeater += 1 {
			head = move(head, direction)
			if shouldFollow(tail, head) {
				tail = followHead(tail, head)
			}
		}
	}

	return unique(tail.locations)
}

func execInstructionsPart2(instructions []string) map[int]knot {
	var knots = make(map[int]knot)

	for idx := 0; idx <= 9; idx += 1 {
		knots[idx] = knot{}
	}

	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		direction := parts[0]
		repeats := convStrToI(parts[1])

		for repeater := 0; repeater < repeats; repeater += 1 {
			knots[0] = move(knots[0], direction)

			for idx := 1; idx <= 9; idx += 1 {
				if shouldFollow(knots[idx], knots[idx-1]) {
					knots[idx] = followHead(knots[idx], knots[idx-1])
				}
			}
		}
	}

	return knots
}

func followHead(tail knot, head knot) knot {
	var location string
	location = strings.Join([]string{strconv.Itoa(tail.x), strconv.Itoa(tail.y)}, ",")
	tail.locations = append(tail.locations, location)

	if head.x-tail.x >= 1 {
		tail.x += 1
	}
	if head.x-tail.x <= -1 {
		tail.x -= 1
	}
	if head.y-tail.y >= 1 {
		tail.y += 1
	}
	if head.y-tail.y <= -1 {
		tail.y -= 1
	}
	location = strings.Join([]string{strconv.Itoa(tail.x), strconv.Itoa(tail.y)}, ",")
	tail.locations = append(tail.locations, location)

	return tail
}

func shouldFollow(tail knot, head knot) bool {
	return abs(head.x-tail.x) > 1 || abs(head.y-tail.y) > 1
}

func move(knot knot, direction string) knot {
	switch direction {
	case UP:
		{
			knot.y += 1
		}
	case RIGHT:
		{
			knot.x += 1
		}
	case DOWN:
		{
			knot.y -= 1
		}
	case LEFT:
		{
			knot.x -= 1
		}
	default:
		fmt.Println("Found direction: ", direction)
		panic("Unknown direction")
	}

	return knot
}

func unique(a []string) (out []string) {
	var keys = make(map[string]bool)

	for _, val := range a {
		if !keys[val] {
			keys[val] = true
			out = append(out, val)
		}
	}

	return
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func readFileAsLines() []string {
	return strings.Split(readFile(), "\n")
}

func readFile() (content string) {
	raw, err := os.ReadFile("input.txt")
	check(err)
	content = strings.TrimSuffix(string(raw), "\n")
	return
}

func convStrToI(s string) (i int) {
	i, err := strconv.Atoi(s)
	check(err)
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
