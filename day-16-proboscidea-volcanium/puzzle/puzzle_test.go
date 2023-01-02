package puzzle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	var cave Cave = parseLine(input()[1])

	assert.Equal(t, "BB", cave.name)
	assert.Equal(t, 13, cave.rate)
	assert.Equal(t, "CC, AA", cave.rawNeighbours)
}

func TestParseInput(t *testing.T) {
	var caves map[string]*Cave = parseLines(input())

	assert.Equal(t, "AA", caves["AA"].name)
	assert.Equal(t, 0, caves["AA"].rate)
	assert.Equal(t, "DD, II, BB", caves["AA"].rawNeighbours)
	// assert.Equal(t, 3, len(caves["AA"].neighbours))
	// assert.Equal(t, "DD", caves["AA"].neighbours[0].name)
	// assert.Equal(t, "II", caves["AA"].neighbours[1].name)
	// assert.Equal(t, "BB", caves["AA"].neighbours[2].name)
}

func TestBuildValidPaths(t *testing.T) {
	caves := parseLines(input())
	paths := buildValidPaths(caves)

	assert.Equal(t, 720, len(paths))
}

func TestFindPath(t *testing.T) {
	caves := parseLines(input())

	path := findPath("AA", "JJ", caves)
	assert.Equal(t, []string{"AA", "II", "JJ"}, path)
}

func TestCalculateScoreByPath(t *testing.T) {
	caves := parseLines(input())
	path := []string{"DD", "BB", "JJ", "HH", "EE", "CC"}

	lastCaveName, score := scoreForPath(path, caves)

	assert.Equal(t, 1651, score)
	assert.Equal(t, "CC", lastCaveName)
}

func TestFindBestPath(t *testing.T) {
	bestScore := FindBestPath(input())
	assert.Equal(t, 1651, bestScore)
}

func TestBuildPermutation(t *testing.T) {
	list := buildPermutations([]string{"CC", "BB", "AA"})

	assert.Equal(t, [][]string{[]string{"CC", "BB", "AA"}, []string{"CC", "AA", "BB"}, []string{"BB", "CC", "AA"}, []string{"BB", "AA", "CC"}, []string{"AA", "CC", "BB"}, []string{"AA", "BB", "CC"}}, list)
}

func TestNextPermutationFromCurrent1(t *testing.T) {
	series := []string{"AA", "BB", "CC", "DD"}
	current := []string{}
	expected := []string{"AA", "BB", "CC", "DD"}

	next := nextPermutationFromCurrent(series, current)

	assert.Equal(t, expected, next, "Test for building the initial permutation, starting with nothing")
}

func TestNextPermutationFromCurrent2(t *testing.T) {
	series := []string{"AA", "BB", "CC", "DD"}
	current := []string{"AA", "CC"}
	expected := []string{"AA", "CC", "BB", "DD"}

	next := nextPermutationFromCurrent(series, current)
	fmt.Println("next permutation:", next)

	assert.Equal(t, expected, next, "Test for building a permutation from a partly series, needed to skip whole ranges of permutations")
}

func TestNextPermutationFromCurrent3(t *testing.T) {
	series := []string{"AA", "BB", "CC", "DD"}
	current := []string{"BB", "DD", "AA", "CC"}
	expected := []string{"BB", "DD", "CC", "AA"}

	next := nextPermutationFromCurrent(series, current)

	assert.Equal(t, expected, next, "Test for build the next permutation, with a 'basic' addition")
}

// Test for build the next permutation, with a larger number of elements in the shift
func TestNextPermutationFromCurrent4(t *testing.T) {
	series := []string{"AA", "BB", "CC", "DD"}
	current := []string{"BB", "CC", "DD", "AA"}
	expected1 := []string{"BB", "DD", "AA", "CC"}
	expected2 := []string{"BB", "DD", "CC", "AA"}
	expected3 := []string{"CC", "AA", "BB", "DD"}

	next1 := nextPermutationFromCurrent(series, current)
	next2 := nextPermutationFromCurrent(series, next1)
	next3 := nextPermutationFromCurrent(series, next2)

	assert.Equal(t, expected1, next1, "Test for build the next permutation, with a larger number of elements in the shift")
	assert.Equal(t, expected2, next2, "Test for build the next permutation, with a larger number of elements in the shift")
	assert.Equal(t, expected3, next3, "Test for build the next permutation, with a larger number of elements in the shift")
}

// Returns empty list of strings after the final permutation
func TestNextPermutationFromCurrent5(t *testing.T) {
	series := []string{"AA", "BB", "CC", "DD"}
	current := []string{"DD", "CC", "AA", "BB"}
	expected1 := []string{"DD", "CC", "BB", "AA"}
	expected2 := []string{}

	next1 := nextPermutationFromCurrent(series, current)
	next2 := nextPermutationFromCurrent(series, next1)

	assert.Equal(t, expected1, next1, "Return empty list of string, after final permutation")
	assert.Equal(t, expected2, next2, "Return empty list of string, after final permutation")
}

func TestNextPermutationRange(t *testing.T) {
	series := []string{"AA", "BB", "CC", "DD"}
	revertSeries := []string{"DD", "CC", "BB", "AA"}
	var current = make([]string, 4)
	current = []string{"BB", "CC"}
	expected1 := []string{"BB", "CC", "DD", "AA"}
	expected2 := []string{"BB", "DD", "AA", "CC"}

	next1 := nextPermutationFromCurrent(revertSeries, current)
	next2 := nextPermutationFromCurrent(series, next1)

	assert.Equal(t, expected1, next1, "Complete the range with the last possible permutation")
	assert.Equal(t, expected2, next2, "Continue after the new range, by continuing after the last possible permutation")
}

func input() []string {
	return []string{
		"Valve AA has flow rate=0; tunnels lead to valves DD, II, BB",
		"Valve BB has flow rate=13; tunnels lead to valves CC, AA",
		"Valve CC has flow rate=2; tunnels lead to valves DD, BB",
		"Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE",
		"Valve EE has flow rate=3; tunnels lead to valves FF, DD",
		"Valve FF has flow rate=0; tunnels lead to valves EE, GG",
		"Valve GG has flow rate=0; tunnels lead to valves FF, HH",
		"Valve HH has flow rate=22; tunnel leads to valve GG",
		"Valve II has flow rate=0; tunnels lead to valves AA, JJ",
		"Valve JJ has flow rate=21; tunnel leads to valve II",
	}
}
