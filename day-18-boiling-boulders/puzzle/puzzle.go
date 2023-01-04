package puzzle

import (
	"strings"

	"aoc.com/2022/day-18/utils"
)

type Coordinate struct {
	x int
	y int
	z int
}

const BOULDER_SIDES = 6
const UNIVERSE_LOW_BOUND = 0
const UNIVERSE_HIGH_BOUND = 24

func (c Coordinate) toIndex() int {
	return c.x*UNIVERSE_HIGH_BOUND*UNIVERSE_HIGH_BOUND + c.y*UNIVERSE_HIGH_BOUND + c.z
}

func (c Coordinate) withinUniverse() bool {
	return c.x >= UNIVERSE_LOW_BOUND && c.x <= UNIVERSE_HIGH_BOUND &&
		c.y >= UNIVERSE_LOW_BOUND && c.y <= UNIVERSE_HIGH_BOUND &&
		c.z >= UNIVERSE_LOW_BOUND && c.z <= UNIVERSE_HIGH_BOUND
}

type Location struct {
	isBoulder  bool // The location is a boulder, due to its coordinates being in the input
	coordinate Coordinate
	neighbours []*Location
}

func CountExposedSides(lines []string) (sides int) {
	boulders := linkNeighbours(parseInput(lines))

	for _, boulder := range boulders {
		sides += BOULDER_SIDES - len(boulder.neighbours)
	}

	return
}

func CountSidesExposedToFreeAir(lines []string) (sides int) {
	boulders := parseInput(lines)
	freeAirLocations := findFreeAirLocations(boulders)
	locations := appendWithTrappedAir(boulders, freeAirLocations)
	locations = linkNeighbours(locations)

	for _, location := range locations {
		if location.isBoulder {
			sides += BOULDER_SIDES - len(location.neighbours)
		}
	}

	return
}

// Visit all locations in the universe.
// If a location is neither a boulder or free air, it must be trapped air.
// Append the trapped-air locations to the list of boulders, since we need to treat them the
// same a real boulders.
func appendWithTrappedAir(boulders map[int]Location, freeAirLocations map[int]Location) map[int]Location {
	var out = make(map[int]Location)

	for x := UNIVERSE_LOW_BOUND; x <= UNIVERSE_HIGH_BOUND; x++ {
		for y := UNIVERSE_LOW_BOUND; y <= UNIVERSE_HIGH_BOUND; y++ {
			for z := UNIVERSE_LOW_BOUND; z <= UNIVERSE_HIGH_BOUND; z++ {
				coor := Coordinate{x, y, z}
				idxCoor := coor.toIndex()
				boulder, exists := boulders[idxCoor]
				_, isFreeAir := freeAirLocations[idxCoor]

				if exists {
					out[idxCoor] = boulder
				}
				if !exists && !isFreeAir {
					// Location is trapped air; append it to the list of boulders
					out[idxCoor] = Location{
						isBoulder:  false,
						coordinate: coor,
					}
				}
			}
		}
	}

	return out
}

// Find all free air locations, via Dijkstra pathfinding.
// Start searching from a location that certainly is free. From there, visit
// each reachable location. Location are unreachable, when its a boulder or falls outside
// the universe.
func findFreeAirLocations(boulders map[int]Location) map[int]Location {
	var freeAirLocations = make(map[int]Location)

	reachedLocs := make(map[int]bool)
	var path = make(map[int][]Coordinate)
	path[0] = []Coordinate{{0, 0, 0}}

	for length := 0; len(path[length]) > 0; length++ {
		for _, currentCoor := range path[length] {
			if currentCoor.withinUniverse() {
				// Current location is within the existing universe
				idxCurrCoor := currentCoor.toIndex()
				if _, existsBoulder := boulders[idxCurrCoor]; !existsBoulder {
					// Current location is not a boulder, and reachable from free air: ergo: it's also free air
					freeAirLocations[idxCurrCoor] = Location{
						isBoulder:  false,
						coordinate: currentCoor,
					}

					for _, rels := range relativeCoordinates() {
						// Add all unreached blocks around the current one to the next path
						nextCoor := Coordinate{
							x: currentCoor.x + rels[0],
							y: currentCoor.y + rels[1],
							z: currentCoor.z + rels[2],
						}
						idxTempCoor := nextCoor.toIndex()
						if !reachedLocs[idxTempCoor] {
							reachedLocs[idxTempCoor] = true
							path[length+1] = append(path[length+1], nextCoor)
						}
					}
				}
			}
		}
	}

	return freeAirLocations
}

// Link the current location to all its neighbouring locations, wether being it a boulder or trapped air.
func linkNeighbours(locations map[int]Location) map[int]Location {
	for idx, location := range locations {
		for _, rels := range relativeCoordinates() {
			idxNeighbour :=
				Coordinate{
					x: location.coordinate.x + rels[0],
					y: location.coordinate.y + rels[1],
					z: location.coordinate.z + rels[2],
				}.toIndex()

			if neighbour, exists := locations[idxNeighbour]; exists {
				location.neighbours = append(location.neighbours, &neighbour)
			}
		}

		locations[idx] = location
	}

	return locations
}

func parseInput(lines []string) map[int]Location {
	var boulders = make(map[int]Location)

	for _, line := range lines {
		boulder := parseLine(line)
		boulders[boulder.coordinate.toIndex()] = boulder
	}

	return boulders
}

func parseLine(line string) (boulder Location) {
	locs := strings.Split(line, ",")
	boulder.isBoulder = true
	boulder.coordinate.x = utils.ConvStrToI(locs[0])
	boulder.coordinate.y = utils.ConvStrToI(locs[1])
	boulder.coordinate.z = utils.ConvStrToI(locs[2])

	return
}

func relativeCoordinates() [][]int {
	return [][]int{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}
}
