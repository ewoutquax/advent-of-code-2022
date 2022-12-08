package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tree struct {
	height        int
	north         *tree
	east          *tree
	south         *tree
	west          *tree
	is_leaf_north bool
	is_leaf_east  bool
	is_leaf_south bool
	is_leaf_west  bool
}

func main() {
	fmt.Println("Result of part-1: ", solvePart1([]string{}))
	fmt.Println("Result of part-2: ", solvePart2([]string{}))
}

func solvePart1(lines []string) int {
	if len(lines) == 0 {
		lines = read_file_as_lines()
	}

	trees := parseInput(lines)

	return countVisibleTrees(trees)
}

func solvePart2(lines []string) int {
	if len(lines) == 0 {
		lines = read_file_as_lines()
	}

	trees := parseInput(lines)

	return maxScenicScore(trees)
}

func maxScenicScore(trees map[int]*tree) (max int) {
	for _, tree := range trees {
		score := getScenicScore(tree)
		if max < score {
			max = score
		}
	}

	return
}

func getScenicScore(tree *tree) int {
	return distanceNorth(tree, tree.height, true) *
		distanceEast(tree, tree.height, true) *
		distanceSouth(tree, tree.height, true) *
		distanceWest(tree, tree.height, true)
}

func distanceNorth(tree *tree, threshold int, current bool) int {
	if !current && tree.height >= threshold {
		return 0
	} else if tree.is_leaf_north {
		return 0
	} else {
		return 1 + distanceNorth(tree.north, threshold, false)
	}
}

func distanceEast(tree *tree, threshold int, current bool) int {
	if !current && tree.height >= threshold {
		return 0
	} else if tree.is_leaf_east {
		return 0
	} else {
		return 1 + distanceEast(tree.east, threshold, false)
	}
}

func distanceSouth(tree *tree, threshold int, current bool) int {
	if !current && tree.height >= threshold {
		return 0
	} else if tree.is_leaf_south {
		return 0
	} else {
		return 1 + distanceSouth(tree.south, threshold, false)
	}
}

func distanceWest(tree *tree, threshold int, current bool) int {
	if !current && tree.height >= threshold {
		return 0
	} else if tree.is_leaf_west {
		return 0
	} else {
		return 1 + distanceWest(tree.west, threshold, false)
	}
}

func countVisibleTrees(trees map[int]*tree) (count int) {
	for _, tree := range trees {
		if isVisible(tree) {
			count += 1
		}
	}

	return
}

func isVisible(tree *tree) bool {
	return allBelowHeightNorth(tree, tree.height, true) ||
		allBelowHeightEast(tree, tree.height, true) ||
		allBelowHeightSouth(tree, tree.height, true) ||
		allBelowHeightWest(tree, tree.height, true)
}

func allBelowHeightNorth(tree *tree, threshold int, current bool) bool {
	if !current && tree.height >= threshold {
		return false
	} else if tree.is_leaf_north {
		return true
	} else {
		return allBelowHeightNorth(tree.north, threshold, false)
	}
}

func allBelowHeightEast(tree *tree, threshold int, current bool) bool {
	if !current && tree.height >= threshold {
		return false
	} else if tree.is_leaf_east {
		return true
	} else {
		return allBelowHeightEast(tree.east, threshold, false)
	}
}

func allBelowHeightSouth(tree *tree, threshold int, current bool) bool {
	if !current && tree.height >= threshold {
		return false
	} else if tree.is_leaf_south {
		return true
	} else {
		return allBelowHeightSouth(tree.south, threshold, false)
	}
}

func allBelowHeightWest(tree *tree, threshold int, current bool) bool {
	if !current && tree.height >= threshold {
		return false
	} else if tree.is_leaf_west {
		return true
	} else {
		return allBelowHeightWest(tree.west, threshold, false)
	}
}

func parseInput(lines []string) map[int]*tree {
	trees := make(map[int]*tree)

	for x, line := range lines {
		for y, height := range strings.Split(line, "") {
			trees = addTree(trees, x, y, conv_str_to_i(height))
		}
	}

	// for loc, tree := range trees {
	// 	fmt.Println("tree on location ", loc, ":", tree)
	// }

	return trees
}

func addTree(trees map[int]*tree, x int, y int, height int) map[int]*tree {
	var tree tree
	tree.height = height
	tree.is_leaf_south = true
	tree.is_leaf_east = true

	if x == 0 {
		tree.is_leaf_west = true
	} else {
		tree.is_leaf_west = false

		location_west := 100*y + (x - 1)
		tree_west := trees[location_west]
		tree_west.is_leaf_east = false
		tree_west.east = &tree
		tree.west = tree_west
	}

	if y == 0 {
		tree.is_leaf_north = true

	} else {
		tree.is_leaf_north = false

		location_north := 100*(y-1) + x
		tree_north := trees[location_north]
		tree_north.is_leaf_south = false
		tree_north.south = &tree
		tree.north = tree_north
	}

	location := 100*y + x
	trees[location] = &tree

	return trees
}

func read_file_as_blocks() (blocks [][]string) {
	var block_inputs []string = strings.Split(read_file(), "\n\n")

	for _, block_input := range block_inputs {
		blocks = append(blocks, strings.Split(block_input, "\n"))
	}
	return
}

func read_file_as_lines() []string {
	return strings.Split(read_file(), "\n")
}

func read_file() (content string) {
	raw, err := os.ReadFile("input.txt")
	check(err)
	content = strings.TrimSuffix(string(raw), "\n")
	return
}

func conv_str_to_i(s string) (i int) {
	i, err := strconv.Atoi(s)
	check(err)
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
