package puzzle

import (
	"fmt"
	"strconv"
	"strings"

	"aoc.com/2022/day-09/utils"
)

const (
	Up    = "U"
	Right = "R"
	Down  = "D"
	Left  = "L"
)

type Knot struct {
	x         int
	y         int
	locations []string
}

func NumberOfLocationsOfKnot(instructions []string, knotId int) int {
	knots := execInstructions(instructions)

	return len(utils.Unique(knots[knotId].locations))
}

func execInstructions(instructions []string) map[int]Knot {
	var knots = make(map[int]Knot)

	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		direction := parts[0]
		repeats := utils.ConvStrToI(parts[1])

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

func followHead(tail Knot, head Knot) Knot {
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

func shouldFollow(tail Knot, head Knot) bool {
	return utils.Abs(head.x-tail.x) > 1 || utils.Abs(head.y-tail.y) > 1
}

func move(knot Knot, direction string) Knot {
	switch direction {
	case Up:
		knot.y += 1
	case Right:
		knot.x += 1
	case Down:
		knot.y -= 1
	case Left:
		knot.x -= 1
	default:
		fmt.Println("Found direction: ", direction)
		panic("Unknown direction")
	}

	return knot
}
