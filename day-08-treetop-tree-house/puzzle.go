package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NORTH = "north"
const EAST = "east"
const SOUTH = "south"
const WEST = "west"

type tree struct {
	height    int
	direction map[string]direction
}
type direction struct {
	tree    *tree
	is_leaf bool
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
	return distance(tree, tree.height, NORTH, true) *
		distance(tree, tree.height, EAST, true) *
		distance(tree, tree.height, SOUTH, true) *
		distance(tree, tree.height, WEST, true)
}

func distance(tree *tree, threshold int, direction string, current bool) int {
	if !current && tree.height >= threshold {
		return 0
	} else if tree.direction[direction].is_leaf {
		return 0
	} else {
		return 1 + distance(tree.direction[direction].tree, threshold, direction, false)
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
	return allBelowHeight(tree, tree.height, NORTH, true) ||
		allBelowHeight(tree, tree.height, EAST, true) ||
		allBelowHeight(tree, tree.height, SOUTH, true) ||
		allBelowHeight(tree, tree.height, WEST, true)
}

func allBelowHeight(tree *tree, threshold int, direction string, current bool) bool {
	if !current && tree.height >= threshold {
		return false
	} else if tree.direction[direction].is_leaf {
		return true
	} else {
		return allBelowHeight(tree.direction[direction].tree, threshold, direction, false)
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

	tree.direction = make(map[string]direction)
	tree.direction[EAST] = direction{is_leaf: true}
	tree.direction[SOUTH] = direction{is_leaf: true}

	if x == 0 {
		tree.direction[WEST] = direction{is_leaf: true}
	} else {
		location_west := 100*y + (x - 1)
		tree_west := trees[location_west]
		tree_west.direction[EAST] = direction{is_leaf: false, tree: &tree}
		tree.direction[WEST] = direction{is_leaf: false, tree: tree_west}
	}

	if y == 0 {
		tree.direction[NORTH] = direction{is_leaf: true}
	} else {
		location_north := 100*(y-1) + x
		tree_north := trees[location_north]
		tree_north.direction[SOUTH] = direction{is_leaf: false, tree: &tree}
		tree.direction[NORTH] = direction{is_leaf: false, tree: tree_north}
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
