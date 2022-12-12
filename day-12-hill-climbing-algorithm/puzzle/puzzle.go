package puzzle

import (
	"strings"
)

type Direction uint

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Location struct {
	coordinate int
	height     int
	isUsed     bool
	isStart    bool
	isEnd      bool
	neighbours map[Direction]*Location
}

const HEIGHTS string = "abcdefghijklmnopqrstuvwxyz"

func FindLowestDistance(lines []string) int {
	locations := parseInput(lines)
	return findPath(locations)
}

func FindLowestDistanceDown(lines []string) int {
	locations := parseInput(lines)
	return findPathDown(locations)
}

func findPath(locations map[int]*Location) (pathLength int) {
	paths := make(map[int][]*Location)

	for _, loc := range locations {
		if loc.isStart {
			paths[0] = []*Location{loc}
		}
	}

	var endReached bool = false
	for pathLength = 0; !endReached; pathLength += 1 {
		for _, currentLoc := range paths[pathLength] {
			for _, nextLoc := range currentLoc.neighbours {
				if !nextLoc.isUsed && nextLoc.height <= currentLoc.height+1 {
					if nextLoc.isEnd {
						endReached = true
					}
					nextLoc.isUsed = true
					paths[pathLength+1] = append(paths[pathLength+1], nextLoc)
				}
			}
		}
	}

	return
}

func findPathDown(locations map[int]*Location) (pathLength int) {
	paths := make(map[int][]*Location)

	for _, loc := range locations {
		if loc.isEnd {
			paths[0] = []*Location{loc}
		}
	}

	var endReached bool = false
	for pathLength = 0; !endReached; pathLength += 1 {
		for _, currentLoc := range paths[pathLength] {
			for _, nextLoc := range currentLoc.neighbours {
				if !nextLoc.isUsed && nextLoc.height >= currentLoc.height-1 {
					if nextLoc.height == 1 {
						endReached = true
					}
					nextLoc.isUsed = true
					paths[pathLength+1] = append(paths[pathLength+1], nextLoc)
				}
			}
		}
	}

	return
}

func parseInput(lines []string) map[int]*Location {
	out := make(map[int]*Location)

	for y, line := range lines {
		for x, value := range strings.Split(line, "") {
			current := Location{
				coordinate: y*100 + x,
				isStart:    false,
				isEnd:      false,
				isUsed:     false,
				neighbours: make(map[Direction]*Location),
			}

			if y > 0 {
				neighbour_up := out[(y-1)*100+x]
				neighbour_up.neighbours[Down] = &current
				current.neighbours[Up] = neighbour_up
			}
			if x > 0 {
				neighbour_left := out[y*100+x-1]
				neighbour_left.neighbours[Right] = &current
				current.neighbours[Left] = neighbour_left
			}

			switch value {
			case "S":
				current.height = 1
				current.isStart = true
			case "E":
				current.height = 26
				current.isEnd = true
			default:
				current.height = strings.Index(HEIGHTS, value) + 1
			}

			out[current.coordinate] = &current
		}
	}

	return out
}
