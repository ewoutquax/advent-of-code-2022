package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NOOP string = "noop"
const ADDX string = "addx"

type state struct {
	value       int
	cycle       int
	frequencies []int
	pixels      [240]string
}

func main() {
	fmt.Println("Result of part-1: ", solvePart1([]string{}))
	fmt.Println("Result of part-2: ", solvePart2([]string{}))
}

func solvePart1(lines []string) (total int) {
	if len(lines) == 0 {
		lines = readFileAsLines()
	}

	endState := getFrequencies(lines)
	for _, value := range endState.frequencies {
		total += value
	}
	return
}

func solvePart2(lines []string) string {
	if len(lines) == 0 {
		lines = readFileAsLines()
	}

	endState := getFrequencies(lines)

	fmt.Println("\nCRT")
	fmt.Println("---")

	for line := 0; line < 6; line += 1 {
		start := line * 40
		end := line*40 + 40
		output := strings.Join(endState.pixels[start:end], "")

		fmt.Println(output)
	}
	fmt.Print("\n\n")

	return "EHBZLRJR"
}

func getFrequencies(instructions []string) state {
	var state state
	state.cycle = 1
	state.value = 1

	for _, instruction := range instructions {
		action, value := parseInstruction(instruction)

		state = execInstruction(action, value, state)
	}

	return state
}

func execInstruction(action string, value int, state state) state {
	switch action {
	case NOOP:
		state = processCycle(state)
	case ADDX:
		state = processCycle(state)
		state = processCycle(state)
		state.value += value
	}

	return state
}

func processCycle(state state) state {
	if state.cycle == 20 || (state.cycle-20)%40 == 0 && state.cycle <= 220 {
		state.frequencies = append(state.frequencies, state.value*state.cycle)
	}

	pixelPos := state.cycle - 1

	if abs(state.value-pixelPos%40) <= 1 {
		state.pixels[pixelPos] = "#"
	} else {
		state.pixels[pixelPos] = "."
	}

	state.cycle += 1

	return state
}

func parseInstruction(instruction string) (string, int) {
	parts := strings.Split(instruction, " ")

	switch parts[0] {
	case NOOP:
		return NOOP, 0
	case ADDX:
		return ADDX, convStrToI(parts[1])
	default:
		fmt.Println("unexpected instruction: ", instruction)
		panic("invalid instruction")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func readFileAsLines() []string {
	return strings.Split(readFile(), "\n")
}

func readFile() (content string) {
	raw, err := os.ReadFile("input.txt")
	check(err)
	content = strings.TrimSuffix(string(raw), "\n")
	return
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
