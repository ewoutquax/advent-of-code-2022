package puzzle

import "strings"

const checkDirection = "NSWE"

func CountEmptyGroundTilesAfterRounds(maxRounds int, lines []string) int {
	movedElves := moveMultipleRounds(maxRounds, parseInput(lines))

	return countEmptyGroundTiles(movedElves)
}

func FirstRoundWithoutMovement(lines []string) (nrRound int) {
	var anyMovement bool = true

	elves := parseInput(lines)

	for nrRound = 0; anyMovement; nrRound++ {
		anyMovement = false
		for _, elf := range elves {
			if elf.shouldMove(elves) {
				anyMovement = true
				break
			}

		}

		if anyMovement {
			setMoveToPosition(nrRound, &elves)
			elves = finalizePositions(elves)
		}
	}

	return
}

func countEmptyGroundTiles(elves map[string]*Elf) (emptyGroundTiles int) {
	var minX, minY int = 9999, 9999
	var maxX, maxY int

	for _, elf := range elves {
		if minX > elf.position.x {
			minX = elf.position.x
		}
		if maxX < elf.position.x {
			maxX = elf.position.x
		}
		if minY > elf.position.y {
			minY = elf.position.y
		}
		if maxY < elf.position.y {
			maxY = elf.position.y
		}
	}

	return ((maxX - minX + 1) * (maxY - minY + 1)) - len(elves)
}

func moveMultipleRounds(maxRounds int, elves map[string]*Elf) map[string]*Elf {
	for nrRound := 0; nrRound < maxRounds; nrRound++ {
		setMoveToPosition(nrRound, &elves)
		elves = finalizePositions(elves)
	}

	return elves
}

func finalizePositions(elves map[string]*Elf) map[string]*Elf {
	var countLocation = make(map[string]int)
	var movedElves = make(map[string]*Elf)

	for _, elf := range elves {
		countLocation[elf.moveTo.toKey()]++
	}

	for _, elf := range elves {
		if countLocation[elf.moveTo.toKey()] == 1 {
			elf.position = *elf.moveTo
		}

		movedElves[elf.position.toKey()] = elf
		elf.moveTo = nil
	}

	return movedElves
}

func setMoveToPosition(nrRound int, elvesByLocation *map[string]*Elf) {
	var directionToCheck string

	for _, elf := range *elvesByLocation {
		if elf.shouldMove(*elvesByLocation) {
			for idx := 0; elf.moveTo == nil && idx < 4; idx++ {
				directionToCheck = string(checkDirection[(nrRound+idx)%4])

				switch directionToCheck {
				case "N":
					if elf.checkNorth(elvesByLocation) {
						elf.setMoveToNorth()
					}
				case "E":
					if elf.checkEast(elvesByLocation) {
						elf.setMoveToEast()
					}
				case "S":
					if elf.checkSouth(elvesByLocation) {
						elf.setMoveToSouth()
					}
				case "W":
					if elf.checkWest(elvesByLocation) {
						elf.setMoveToWest()
					}
				}
			}
		}

		if elf.moveTo == nil {
			elf.setMoveToCurrentLocation()
		}
	}
}

func parseInput(lines []string) map[string]*Elf {
	elvesByLocation := make(map[string]*Elf)

	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if char == "#" {
				elf := Elf{
					position: Position{x, y},
				}
				elvesByLocation[elf.position.toKey()] = &elf
			}
		}
	}

	return elvesByLocation
}
