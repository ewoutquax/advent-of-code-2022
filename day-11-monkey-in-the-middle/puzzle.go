package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	id                  int
	items               []int
	worryOperation      string
	worryUseItemValue   bool
	worryValue          int
	testDivisible       int
	testTrue            int
	testFalse           int
	itemsInspectedCount int
}

func main() {
	fmt.Println("Result of part-1: ", solvePart1([][]string{{}}))
	fmt.Println("Result of part-2: ", solvePart2([][]string{{}}))
}

func solvePart1(blocks [][]string) int {
	if len(blocks) == 1 {
		blocks = readFileAsBlocks()
	}
	monkeys := parseMonkeys(blocks)
	for round := 0; round < 20; round += 1 {
		monkeys = playRound(monkeys, "V1")
	}

	return calculateMonkeyBusiness(monkeys)
}

func solvePart2(blocks [][]string) int {
	if len(blocks) == 1 {
		blocks = readFileAsBlocks()
	}
	monkeys := parseMonkeys(blocks)
	for round := 0; round < 10_000; round += 1 {
		monkeys = playRound(monkeys, "V2")
	}

	return calculateMonkeyBusiness(monkeys)
}

func calculateMonkeyBusiness(monkeys map[int]monkey) (business int) {
	var countInspection []int
	for _, monkey := range monkeys {
		countInspection = append(countInspection, monkey.itemsInspectedCount)
	}

	sort.Ints(countInspection)

	return countInspection[len(countInspection)-2] * countInspection[len(countInspection)-1]
}

func playRound(monkeys map[int]monkey, version string) map[int]monkey {
	var targetMonkeyId, newItemValue int
	var lcm int = 1
	for _, monkey := range monkeys {
		lcm *= monkey.testDivisible
	}

	for idx := 0; idx < len(monkeys); idx += 1 {
		monkey := monkeys[idx]

		for _, item := range monkey.items {
			if version == "V1" {
				targetMonkeyId, newItemValue = determineTargetMonkeyIdV1(item, monkey)
			} else {
				targetMonkeyId, newItemValue = determineTargetMonkeyIdV2(item, monkey, lcm)
			}

			targetMonkey := monkeys[targetMonkeyId]
			targetMonkey.items = append(monkeys[targetMonkeyId].items, newItemValue)
			monkeys[targetMonkeyId] = targetMonkey
		}
		monkey.itemsInspectedCount += len(monkey.items)
		monkey.items = []int{}
		monkeys[monkey.id] = monkey
	}

	return monkeys
}

func determineTargetMonkeyIdV1(item int, sourceMonkey monkey) (targetMonkeyId int, newItemValue int) {
	var concreteValue int

	if sourceMonkey.worryUseItemValue {
		concreteValue = item
	} else {
		concreteValue = sourceMonkey.worryValue
	}

	if sourceMonkey.worryOperation == "MULTIPLY" {
		newItemValue = (item * concreteValue) / 3
	} else {
		newItemValue = (item + concreteValue) / 3
	}

	if newItemValue%sourceMonkey.testDivisible == 0 {
		targetMonkeyId = sourceMonkey.testTrue
	} else {
		targetMonkeyId = sourceMonkey.testFalse
	}

	return
}

func determineTargetMonkeyIdV2(item int, sourceMonkey monkey, lcm int) (targetMonkeyId int, newItemValue int) {
	var concreteValue int

	if sourceMonkey.worryUseItemValue {
		concreteValue = item
	} else {
		concreteValue = sourceMonkey.worryValue
	}

	if sourceMonkey.worryOperation == "MULTIPLY" {
		newItemValue = (item * concreteValue) % lcm
	} else {
		newItemValue = (item + concreteValue) % lcm
	}

	if newItemValue%sourceMonkey.testDivisible == 0 {
		targetMonkeyId = sourceMonkey.testTrue
	} else {
		targetMonkeyId = sourceMonkey.testFalse
	}

	return
}

func parseMonkeys(blocks [][]string) map[int]monkey {
	out := make(map[int]monkey)
	for _, block := range blocks {
		monkey := parseMonkey(block)
		out[monkey.id] = monkey
	}

	return out
}

func parseMonkey(lines []string) (out monkey) {
	var lineParts []string
	// id
	lineParts = strings.Split(lines[0], " ")
	lineParts = strings.Split(lineParts[1], ":")
	out.id = convStrToI(lineParts[0])

	// items
	lineParts = strings.Split(lines[1], ": ")
	lineParts = strings.Split(lineParts[1], ", ")
	for _, itemId := range lineParts {
		out.items = append(out.items, convStrToI(itemId))
	}

	// worryOperation + worryValue
	lineParts = strings.Split(lines[2], " old ")
	lineParts = strings.Split(lineParts[1], " ")
	if lineParts[0] == "*" {
		out.worryOperation = "MULTIPLY"
	} else {
		out.worryOperation = "ADD"
	}
	if lineParts[1] == "old" {
		out.worryUseItemValue = true
	} else {
		out.worryUseItemValue = false
		out.worryValue = convStrToI(lineParts[1])
	}

	// testDivisible
	lineParts = strings.Split(lines[3], " by ")
	out.testDivisible = convStrToI(lineParts[1])

	// testTrue
	lineParts = strings.Split(lines[4], " monkey ")
	out.testTrue = convStrToI(lineParts[1])

	// testFalse
	lineParts = strings.Split(lines[5], " monkey ")
	out.testFalse = convStrToI(lineParts[1])

	return
}

func readFileAsBlocks() (blocks [][]string) {
	var block_inputs []string = strings.Split(readFile(), "\n\n")

	for _, block_input := range block_inputs {
		blocks = append(blocks, strings.Split(block_input, "\n"))
	}
	return
}

func readFile() string {
	raw, err := os.ReadFile("input.txt")
	check(err)
	return strings.TrimSuffix(string(raw), "\n")
}

func convStrToI(s string) (i int) {
	i, err := strconv.Atoi(s)
	check(err)
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
