package puzzle

import (
	"strings"

	"github.com/ewoutquax/advent-of-code-2022/utils"
)

type Position struct {
	x int // Column of the location in the maze
	y int // Row of the location in the maze
}

func (p Position) toKey() int {
	return p.x*200 + p.y
}

type Location struct {
	position Position
	isWall   bool
	left     *Location
	right    *Location
	up       *Location
	down     *Location
}

func (l Location) toKey() int {
	return l.position.toKey()
}

type Direction uint

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Player struct {
	location Location
	facing   Direction
}

func (p Player) password() int {
	var facingValue int

	switch p.facing {
	case Right:
		facingValue = 0
	case Down:
		facingValue = 1
	case Left:
		facingValue = 2
	case Up:
		facingValue = 3
	}

	return (p.location.position.y+1)*1000 + (p.location.position.x+1)*4 + facingValue
}

// Make the player turn right 90 degrees in place: don't change location
func (p Player) turnRight() Player {
	switch p.facing {
	case Up:
		p.facing = Right
	case Right:
		p.facing = Down
	case Down:
		p.facing = Left
	case Left:
		p.facing = Up
	}

	return p
}

// Move the player forward 1 position the direction he's facing.
// Don't move if the target location is a wall.
func (p Player) moveForward() Player {
	var nextLoc *Location
	switch p.facing {
	case Up:
		nextLoc = p.location.up
	case Right:
		nextLoc = p.location.right
	case Down:
		nextLoc = p.location.down
	case Left:
		nextLoc = p.location.left
	}

	if !nextLoc.isWall {
		p.location = *nextLoc
	}

	return p
}

func PlayerPassword(blocks [][]string) int {
	topleft, steps := parseInput(blocks)
	player := followSteps(steps, initPlayer(topleft))

	return player.password()
}

func followSteps(steps string, player Player) Player {
	var nrSteps int
	var restSteps string

	if steps != "" {
		if string(steps[0]) == "R" {
			// Make the player turn right
			return followSteps(steps[1:], player.turnRight())
		} else if string(steps[0]) == "L" {
			// Make the player turn left, by turning him right 3 times
			return followSteps(steps[1:], player.turnRight().turnRight().turnRight())
		} else {
			// Determine how many steps to take
			idxLeft := strings.Index(steps, "L")
			idxRight := strings.Index(steps, "R")

			if idxLeft == -1 && idxRight == -1 {
				nrSteps = utils.ConvStrToI(steps)
				restSteps = ""
			} else if idxRight == -1 || idxLeft > 0 && idxRight > 0 && idxLeft < idxRight {
				nrSteps = utils.ConvStrToI(steps[:idxLeft])
				restSteps = steps[idxLeft:]
			} else {
				nrSteps = utils.ConvStrToI(steps[:idxRight])
				restSteps = steps[idxRight:]
			}

			for idx := nrSteps; idx > 0; idx -= 1 {
				player = player.moveForward()
			}

			return followSteps(restSteps, player)
		}
	}

	return player
}

func initPlayer(l Location) Player {
	return Player{
		location: l,
		facing:   Right,
	}
}

func parseInput(lines [][]string) (Location, string) {
	var topLocationPerColumns = make(map[int]*Location)
	var leftLocationPerRow *Location
	var ptrTopLeft *Location
	var locations = make(map[int]*Location)

	maze_lines := lines[0]

	for y, maze_line := range maze_lines {
		leftLocationPerRow = nil

		for x, charLocation := range strings.Split(maze_line, "") {
			if charLocation == " " {
				continue
			} else {
				currentLocation := Location{
					position: Position{x: x, y: y},
					isWall:   charLocation == "#",
				}

				// Add the new location to the fill list of all locations, for easier locating of locations
				locations[currentLocation.toKey()] = &currentLocation

				// Mark the location at the most left side of the first row, as it is the start location
				if ptrTopLeft == nil {
					ptrTopLeft = &currentLocation
				}

				// Set left and right neighbours.
				// When starting to parse a new row, a location can point to itself.
				if leftLocationPerRow == nil {
					leftLocationPerRow = &currentLocation
					currentLocation.left = &currentLocation
				} else {
					key := Position{x: x - 1, y: y}.toKey()
					currentLocation.left = locations[key]
					currentLocation.left.right = &currentLocation
				}
				currentLocation.right = leftLocationPerRow
				currentLocation.right.left = &currentLocation

				// Set up and down neighbours
				// When starting to parse a new column, a location can point to itself.
				if topLocationPerColumns[x] == nil {
					topLocationPerColumns[x] = &currentLocation
					currentLocation.up = &currentLocation
				} else {
					key := Position{x: x, y: y - 1}.toKey()
					currentLocation.up = locations[key]
					currentLocation.up.down = &currentLocation
				}
				currentLocation.down = topLocationPerColumns[x]
				currentLocation.down.up = &currentLocation
			}
		}
	}

	return *ptrTopLeft, lines[1][0]
}
