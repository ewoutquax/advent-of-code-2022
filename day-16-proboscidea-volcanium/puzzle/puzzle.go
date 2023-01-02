package puzzle

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"aoc.com/2022/day-16/utils"
)

type Cave struct {
	name          string
	rate          int
	rawNeighbours string
	neighbours    []*Cave
	originCave    *Cave
}

func FindBestPath(lines []string) (highScore int) {
	caves := parseLines(lines)

	var series []string = []string{"AA"}
	for _, cave := range caves {
		if cave.rate > 0 {
			series = append(series, cave.name)
		}
	}
	sort.Strings(series)
	fmt.Println("series (caves with rate > 0): ", series)

	var copySeries = make([]string, len(series))
	copy(copySeries, series)

	for currentPath := copySeries; len(currentPath) > 0; currentPath = nextPermutationFromCurrent(series, currentPath) {
		if currentPath[0] != "AA" {
			break
		}
		nameLastCave, score := scoreForPath(currentPath, caves)
		fmt.Print("\nScore for path '", currentPath, "': ", score)
		if highScore < score {
			highScore = score
		}

		if nameLastCave != currentPath[len(currentPath)-1] {
			idxLastCave := idxInSlice(currentPath, nameLastCave)

			var myRange = make([]string, idxLastCave+1)
			copy(myRange, currentPath[:idxLastCave+1])

			currentPath = nextPermutationFromCurrent(revertSlice(series), myRange)

			// fmt.Print("\nSkip range '", myRange, "', and continue with path: '", currentPath)
		}
	}

	return
}

func scoreForPath(path []string, caves map[string]*Cave) (lastCaveName string, score int) {
	var minutes, flowRate int
	var currentCaveName string = "AA"

	for _, destinationCaveName := range path {
		if destinationCaveName == "AA" {
			continue
		}
		steps := findPath(currentCaveName, destinationCaveName, caves)
		for idx := 0; idx < len(steps)-1; idx += 1 {
			// fmt.Println("\nCurrent cave", steps[idx])
			if minutes < 30 {
				minutes += 1
				score += flowRate

				// fmt.Println("\nMinute:", minutes)
				// fmt.Println("Moving to cave:", destinationCaveName)
				// fmt.Println("Increase score with", flowRate, "to", score)
			} else {
				lastCaveName = currentCaveName
				// fmt.Println("\n\nIn 30 minutes, we made it to cave:", lastCaveName, "while the whole path is:", path)
				return
			}
		}
		currentCaveName = destinationCaveName
		if minutes < 30 {
			minutes += 1
			score += flowRate

			// fmt.Println("\nMinute:", minutes)
			// fmt.Println("Arrived in cave", destinationCaveName)
			// fmt.Println("Increase score with", flowRate, "to", score)

			flowRate += caves[destinationCaveName].rate
			// fmt.Println("Open valve, and increase flowrate to:", flowRate)
		}
	}

	for _ = minutes; minutes < 30; minutes += 1 {
		// fmt.Println("\nMinute:", minutes)
		// fmt.Println("Don't make a move")
		// fmt.Println("Increase score with", flowRate, "to", score)

		score += flowRate
	}

	lastCaveName = path[len(path)-1]
	return
}

func findPath(startName string, endName string, caves map[string]*Cave) []string {
	// fmt.Println("findPath: find path between", startName, "and", endName)
	var nextCave *Cave
	var startCave *Cave = caves[startName]
	visitedCaves := make(map[string]bool)

	paths := make(map[int][]*Cave)
	paths[0] = []*Cave{startCave}

	var endReached bool = false
	for pathLength := 0; !endReached; pathLength += 1 {
		for _, currentCave := range paths[pathLength] {
			// fmt.Println("findPath: currentCave:", currentCave.name)
			// fmt.Println("rawNeighbours:", currentCave.rawNeighbours)

			if currentCave.name == endName {
				// fmt.Println("Reached the target cave!")
				endReached = true
			} else {
				neighbours := strings.Split(currentCave.rawNeighbours, ", ")

				for _, nextCaveName := range neighbours {
					if visitedCaves[nextCaveName] {
						// fmt.Println("findPath: skip already visited cave:", nextCaveName)
					} else {
						// fmt.Println("findPath: adding to path unvisited cave:", nextCaveName)

						visitedCaves[nextCaveName] = true
						nextCave = caves[nextCaveName]
						paths[pathLength+1] = append(paths[pathLength+1], nextCave)
						nextCave.originCave = currentCave
					}
				}
			}
		}
	}

	return buildPath(startName, caves[endName])
}

func buildPath(startName string, cave *Cave) []string {
	// fmt.Println("buildPath: cave:", cave)

	if cave.name == startName {
		return []string{startName}
	} else {
		return append(buildPath(startName, cave.originCave), cave.name)
	}
}

func parseLines(lines []string) map[string]*Cave {
	var caves = make(map[string]*Cave)

	for _, line := range lines {
		cave := parseLine(line)
		caves[cave.name] = &cave
	}

	return caves
}

func parseLine(line string) (cave Cave) {
	r, _ := regexp.Compile(`Valve ([A-Z]{2}) has flow rate=(\d+); tunnel[s]? lead[s]? to valve[s]? (.*)$`)
	matches := r.FindStringSubmatch(line)

	cave.name = matches[1]
	cave.rate = utils.ConvStrToI(matches[2])
	cave.rawNeighbours = matches[3]

	return
}

func nextPermutationFromCurrent(series []string, current []string) []string {
	if len(series) == len(current) {
		var copyCurrent = make([]string, len(current))
		copy(copyCurrent, current)

		prevValue := -1
		for idx := len(copyCurrent) - 1; idx >= 0; idx -= 1 {
			curValue := idxInSlice(series, copyCurrent[idx])
			if prevValue < curValue {
				prevValue = curValue
			} else {
				bumpValue := copyCurrent[idx]
				trim := copyCurrent[idx:]
				sort.Strings(trim)
				idxTrim := idxInSlice(trim, bumpValue)
				newCurrent := append(copyCurrent[:idx], trim[idxTrim+1])
				return nextPermutationFromCurrent(series, newCurrent)
			}
		}
	} else if len(current) > 0 {
		var copySeries = make([]string, len(series))
		copy(copySeries, series)

		next := current[0]

		index := idxInSlice(copySeries, next)
		restSeries := append(copySeries[:index], copySeries[index+1:]...)

		return append([]string{next}, nextPermutationFromCurrent(restSeries, current[1:])...)
	} else {
		return series
	}

	return []string{}
}

func idxInSlice(slice []string, target string) int {
	for idx, elm := range slice {
		if elm == target {
			return idx
		}
	}

	return -1
}

func revertSlice(slice []string) []string {
	copySlice := make([]string, 0)
	for idx := len(slice) - 1; idx >= 0; idx -= 1 {
		copySlice = append(copySlice, slice[idx])
	}

	return copySlice
}

func buildValidPaths(caves map[string]*Cave) [][]string {
	var validCaveNames []string

	for _, cave := range caves {
		if cave.rate > 0 && cave.name != "AA" {
			validCaveNames = append(validCaveNames, cave.name)
		}
	}

	fmt.Println("Build permutations from", len(validCaveNames), "caves")
	return buildPermutations(validCaveNames)
}

func buildPermutations(names []string) (myPermutations [][]string) {
	if len(names) == 1 {
		myPermutations = [][]string{names}
	} else {
		for idx := 0; idx < len(names); idx += 1 {
			var copyStrings = make([]string, len(names))
			copy(copyStrings, names)
			otherPermutations := buildPermutations(append(copyStrings[:idx], copyStrings[idx+1:]...))

			for _, otherPermutation := range otherPermutations {
				myPermutations = append(myPermutations, append([]string{names[idx]}, otherPermutation...))
			}
		}
	}

	return
}
