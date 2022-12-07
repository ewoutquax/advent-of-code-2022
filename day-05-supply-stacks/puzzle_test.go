package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	assert.Equal(t, parsedCrates(), parseCrates(inputCrates(), 3))
	assert.Equal(t, movedCrates(), execInstructions(parsedCrates(), instructions()))
	assert.Equal(t, "CMZ", generateSolution(movedCrates()))
}

func TestPart2Examples(t *testing.T) {
	assert.Equal(t, movedCratesV2(), execInstructionsV2(parsedCrates(), instructions()))
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, "HBTMTBSDC", solvePart1(9))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, "PQTJRSHWS", solvePart2(9))
}

func inputCrates() []string {
	return []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
	}
}

func instructions() []string {
	return []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
}

func parsedCrates() map[int]string {
	return map[int]string{
		1: "ZN",
		2: "MCD",
		3: "P",
	}
}

func movedCrates() map[int]string {
	return map[int]string{
		1: "C",
		2: "M",
		3: "PDNZ",
	}
}

func movedCratesV2() map[int]string {
	return map[int]string{
		1: "M",
		2: "C",
		3: "PZND",
	}
}
