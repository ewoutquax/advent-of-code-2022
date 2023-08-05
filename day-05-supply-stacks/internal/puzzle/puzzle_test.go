package puzzle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCrates(t *testing.T) {
	universe := parseInput(inputCrates(), 3)
	assert.Equal(t, parsedCrates(), universe)
}

func TestPart1Examples(t *testing.T) {
	universe := parseInput(inputCrates(), 3)
	universe.execInstructionsV1(instructions())

	assert.Equal(t, movedCratesV1(), universe)
	assert.Equal(t, "CMZ", universe.generateSolution())
}

func TestPart2Examples(t *testing.T) {
	universe := parsedCrates()
	universe.execInstructionsV2(instructions())
	assert.Equal(t, movedCratesV2(), universe)
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

func parsedCrates() Universe {
	return Universe{
		stack: map[int][]Crate{
			1: {"Z", "N"},
			2: {"M", "C", "D"},
			3: {"P"},
		},
	}
}

func movedCratesV1() Universe {
	return Universe{
		stack: map[int][]Crate{
			1: {"C"},
			2: {"M"},
			3: {"P", "D", "N", "Z"},
		}}
}

func movedCratesV2() Universe {
	return Universe{
		stack: map[int][]Crate{
			1: {"M"},
			2: {"C"},
			3: {"P", "Z", "N", "D"},
		}}
}
