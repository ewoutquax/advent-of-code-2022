package puzzle

import (
	"testing"

	"github.com/ewoutquax/advent-of-code-2022/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	monkeys := parseInput(myInput())
	assert.Equal(t, 3, len(monkeys))

	var names []string
	for name := range monkeys {
		names = append(names, name)
	}
	assert.True(t, utils.InSlice(names, "root"))

	rootMonkey := monkeys["root"]
	assert.False(t, rootMonkey.hasRawValue)
	assert.Equal(t, "root", rootMonkey.name)
	assert.Equal(t, Add, rootMonkey.operation)
	assert.Equal(t, "aa", rootMonkey.left.name)
	assert.Equal(t, "bb", rootMonkey.right.name)
}

func TestMyMonkeyGetValue(t *testing.T) {
	monkeys := parseInput(myInput())
	rootMonkey := monkeys["root"]

	assert.Equal(t, 3, rootMonkey.getValue())
}

func TestMonkeyGetValue(t *testing.T) {
	monkeys := parseInput(input())
	rootMonkey := monkeys["root"]

	assert.Equal(t, 152, rootMonkey.getValue())
}

func TestNeededHumanValue(t *testing.T) {
	monkeys := parseInput(input())
	monkeys["root"].setNeededHumanValue()

	assert.Equal(t, 301, monkeys["humn"].getValue())
}

func myInput() []string {
	return []string{
		"root: aa + bb",
		"aa: 1",
		"bb: 2",
	}
}

func input() []string {
	return []string{
		"root: pppw + sjmn",
		"dbpl: 5",
		"cczh: sllz + lgvd",
		"zczc: 2",
		"ptdq: humn - dvpt",
		"dvpt: 3",
		"lfqf: 4",
		"humn: 5",
		"ljgn: 2",
		"sjmn: drzm * dbpl",
		"sllz: 4",
		"pppw: cczh / lfqf",
		"lgvd: ljgn * ptdq",
		"drzm: hmdt - zczc",
		"hmdt: 32",
	}
}
