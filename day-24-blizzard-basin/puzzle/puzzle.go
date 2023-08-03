package puzzle

import (
	"strconv"
	"strings"
)

const WALL string = "#"
const START_LOCATION string = "E"
const EMPTY_LOCATION string = "."

func LeastStepsToExitStartAndExit(lines []string) int {
	var basin Basin = parseInput(lines)

	basin.moveToEnd = true
	basin.moveToBeginning = false
	basin.playerReachedGoal = false
	basin.players = []Player{{position: Position{x: 1, y: 0}}}

	for !basin.playerReachedGoal {
		basin.nrStep++
		basin.moveBlizzards()
		basin.movePlayers()
	}

	basin.moveToEnd = false
	basin.moveToBeginning = true
	basin.playerReachedGoal = false
	basin.players = []Player{{position: Position{x: basin.maxX, y: basin.maxY + 1}}}

	for !basin.playerReachedGoal {
		basin.nrStep++
		basin.moveBlizzards()
		basin.movePlayers()
	}

	basin.moveToEnd = true
	basin.moveToBeginning = false
	basin.playerReachedGoal = false
	basin.players = []Player{{position: Position{x: 1, y: 0}}}

	for !basin.playerReachedGoal {
		basin.nrStep++
		basin.moveBlizzards()
		basin.movePlayers()
	}

	return basin.nrStep
}

func LeastStepsToExit(lines []string) int {
	var basin Basin = parseInput(lines)

	basin.moveToEnd = true
	basin.moveToBeginning = false
	basin.playerReachedGoal = false

	for !basin.playerReachedGoal {
		basin.nrStep++
		basin.moveBlizzards()
		basin.movePlayers()
	}

	return basin.nrStep
}

func parseInput(lines []string) (basin Basin) {
	basin.players = []Player{{position: Position{x: 1, y: 0}}}
	basin.wallLocations = make(map[string]bool)

	for y, line := range lines {
		if basin.minY == 0 {
			basin.minY = y
		}
		if basin.maxY < y-1 {
			basin.maxY = y - 1
		}

		for x, char := range strings.Split(line, "") {
			if char == START_LOCATION {
				continue
			}
			if char == WALL {
				key := strconv.Itoa(x) + "," + strconv.Itoa(y)
				basin.wallLocations[key] = true
				continue
			}

			if basin.minX == 0 {
				basin.minX = x
			}
			if basin.maxX < x {
				basin.maxX = x
			}

			if char == EMPTY_LOCATION {
				continue
			}

			blizzard := Blizzard{position: Position{x: x, y: y}}

			switch char {
			case "^":
				blizzard.direction = Up
			case ">":
				blizzard.direction = Right
			case "v":
				blizzard.direction = Down
			case "<":
				blizzard.direction = Left
			default:
				panic("Unknow char to parse")
			}

			basin.blizzards = append(basin.blizzards, &blizzard)
		}
	}

	return basin
}
