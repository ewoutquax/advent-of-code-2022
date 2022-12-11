package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	var expectedMonkey monkey = monkey{
		id:                  0,
		items:               []int{79, 98},
		worryOperation:      "MULTIPLY",
		worryUseItemValue:   false,
		worryValue:          19,
		testDivisible:       23,
		testTrue:            2,
		testFalse:           3,
		itemsInspectedCount: 0,
	}

	assert.Equal(t, expectedMonkey, parseMonkey(input()[0]))

	var parsedMonkeys map[int]monkey = parseMonkeys(input())
	assert.Equal(t, expectedMonkey, parsedMonkeys[0])
	assert.Equal(t, 4, len(parsedMonkeys))

	targetMonkeyId, newItemValue := determineTargetMonkeyIdV1(79, parsedMonkeys[0])
	assert.Equal(t, 3, targetMonkeyId)
	assert.Equal(t, 500, newItemValue)

	expectedItems := [][]int{{20, 23, 27, 26}, {2080, 25, 167, 207, 401, 1046}, {}, {}}
	monkeys := playRound(parsedMonkeys, "V1")
	var collectedItems [][]int
	for idx := 0; idx < len(monkeys); idx += 1 {
		collectedItems = append(collectedItems, monkeys[idx].items)
	}
	assert.Equal(t, expectedItems, collectedItems)
}

func TestPart1Play20Rounds(t *testing.T) {
	var monkeys map[int]monkey = parseMonkeys(input())
	for round := 0; round < 20; round += 1 {
		monkeys = playRound(monkeys, "V1")
		collectedItems := make(map[int][]int)
		for idx := 0; idx < len(monkeys); idx += 1 {
			collectedItems[idx] = monkeys[idx].items
		}

		fmt.Println("Collected item after round '", round, "': ", collectedItems)
	}

	var foundItems [][]int
	for idx := 0; idx < len(monkeys); idx += 1 {
		foundItems = append(foundItems, monkeys[idx].items)
	}

	expectedItems := [][]int{{10, 12, 14, 26, 34}, {245, 93, 53, 199, 115}, {}, {}}
	assert.Equal(t, expectedItems, foundItems)

	count := solvePart1(input())
	assert.Equal(t, 10605, count)
}

func TestPart2Examples(t *testing.T) {
	count := solvePart2(input())
	assert.Equal(t, 2713310158, count)

}

func TestSolvePart1(t *testing.T) {
	count := solvePart1([][]string{{}})
	assert.Equal(t, 90882, count)
}

func TestSolvePart2(t *testing.T) {
	count := solvePart2([][]string{{}})
	assert.Equal(t, 30893109657, count)
}

func input() [][]string {
	return [][]string{{
		"Monkey 0:",
		"  Starting items: 79, 98",
		"  Operation: new = old * 19",
		"  Test: divisible by 23",
		"    If true: throw to monkey 2",
		"    If false: throw to monkey 3",
	}, {
		"Monkey 1:",
		"  Starting items: 54, 65, 75, 74",
		"  Operation: new = old + 6",
		"  Test: divisible by 19",
		"    If true: throw to monkey 2",
		"    If false: throw to monkey 0",
	}, {
		"Monkey 2:",
		"  Starting items: 79, 60, 97",
		"  Operation: new = old * old",
		"  Test: divisible by 13",
		"    If true: throw to monkey 1",
		"    If false: throw to monkey 3",
	}, {
		"Monkey 3:",
		"  Starting items: 74",
		"  Operation: new = old + 3",
		"  Test: divisible by 17",
		"    If true: throw to monkey 0",
		"    If false: throw to monkey 1",
	},
	}
}
