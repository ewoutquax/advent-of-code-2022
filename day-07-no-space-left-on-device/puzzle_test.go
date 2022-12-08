package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	rootDir := parseOutput(terminalOutput())

	assert.Equal(t, "root", rootDir.name)
	assert.Equal(t, 2, len(rootDir.ptrChildDirs))

	var dirSize int = getDirSize(rootDir)
	assert.Equal(t, 48381165, dirSize)

	var dirs []dir = collectDirs(rootDir)
	assert.Equal(t, 4, len(dirs))

	assert.Equal(t, 95437, solvePart1(100000, terminalOutput()))
}

func TestPart2Examples(t *testing.T) {
	rootDir := parseOutput(terminalOutput())

	size := freeableSize(70000000, 30000000, rootDir)
	assert.Equal(t, 24933642, size)
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 1206825, solvePart1(100000, []string{}))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 9608311, solvePart2(70000000, 30000000, []string{}))
}

func terminalOutput() []string {
	return []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	}
}
