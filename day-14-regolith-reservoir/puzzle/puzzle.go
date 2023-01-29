package puzzle

import "strings"

func CountSandInUniverseUntilFloor(lines []string) int {
	blockedPositions, maxY := parseInput(lines)
	var universe Universe = Universe{
		blockedPositions: blockedPositions,
		floorY:           maxY + 2,
	}

	universe.fillWithSandUntilFloor()

	return universe.nrSands
}

func CountSandInUniverseUntilTop(lines []string) int {
	blockedPositions, maxY := parseInput(lines)
	var universe Universe = Universe{
		blockedPositions: blockedPositions,
		floorY:           maxY + 2,
	}

	universe.fillWithSandUntilTop()

	return universe.nrSands
}

func parseInput(lines []string) (map[string]bool, int) {
	var blockedPositions = make(map[string]bool)
	var positions []Position
	var maxY int = 0

	for _, line := range lines {
		positions = parseLine(line)

		for _, position := range positions {
			blockedPositions[position.toKey()] = true
			if maxY < position.y {
				maxY = position.y
			}
		}
	}

	return blockedPositions, maxY
}

func parseLine(line string) (positions []Position) {
	var from, intermediate, to Position

	points := strings.Split(line, " -> ")
	from = keyToPosition(points[0])
	positions = append(positions, from)

	// loop through all points in the line
	for idx := 1; idx < len(points); idx++ {
		to = keyToPosition(points[idx])

		// loop through all positions between two points
		for steps := 1; intermediate != to; steps += 1 {
			intermediate = keyWithStepsToEndPosition(from, to, steps)
			positions = append(positions, intermediate)
		}

		from = to
	}

	return
}
