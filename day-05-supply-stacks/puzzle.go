package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Result of part-1: ", solvePart1(9))
	fmt.Println("Result of part-2: ", solvePart2(9))
}

func solvePart1(nrCrates int) string {
	blocks := read_file_as_blocks()

	var crates []string = blocks[0]
	var instructions []string = blocks[1]

	parsedCrates := parseCrates(crates, nrCrates)
	movedCrates := execInstructions(parsedCrates, instructions)

	return generateSolution(movedCrates)
}

func solvePart2(nrCrates int) string {
	blocks := read_file_as_blocks()

	var crates []string = blocks[0]
	var instructions []string = blocks[1]

	parsedCrates := parseCrates(crates, nrCrates)
	movedCrates := execInstructionsV2(parsedCrates, instructions)

	return generateSolution(movedCrates)
}

func generateSolution(crates map[int]string) (solution string) {
	for idx := 1; idx <= len(crates); idx += 1 {
		stack := crates[idx]
		crate := string(stack[len(stack)-1])
		solution += crate
	}

	return
}

func execInstructions(crates map[int]string, instructions []string) map[int]string {
	for _, instruction := range instructions {
		crates = execInstruction(crates, instruction)
	}

	return crates
}

func execInstructionsV2(crates map[int]string, instructions []string) map[int]string {
	for _, instruction := range instructions {
		crates = execInstructionV2(crates, instruction)
	}

	return crates
}

func execInstruction(crates map[int]string, instruction string) map[int]string {
	r, _ := regexp.Compile(`\Amove (\d+) from (\d+) to (\d+)\z`)
	matches := r.FindStringSubmatch(instruction)

	var repeat int = conv_str_to_i(matches[1])
	var idx_source int = conv_str_to_i(matches[2])
	var idx_target int = conv_str_to_i(matches[3])

	for count := 0; count < repeat; count += 1 {
		stack_source := crates[idx_source]
		stack_target := crates[idx_target]

		crate := string(stack_source[len(stack_source)-1])

		crates[idx_source] = string(stack_source[0 : len(stack_source)-1])
		crates[idx_target] = stack_target + crate
	}

	return crates
}

func execInstructionV2(crates map[int]string, instruction string) map[int]string {
	r, _ := regexp.Compile(`\Amove (\d+) from (\d+) to (\d+)\z`)
	matches := r.FindStringSubmatch(instruction)

	var size int = conv_str_to_i(matches[1])
	var idx_source int = conv_str_to_i(matches[2])
	var idx_target int = conv_str_to_i(matches[3])

	stack_source := crates[idx_source]
	stack_target := crates[idx_target]

  currentStack := string(stack_source[len(stack_source)-(size):len(stack_source)])

	crates[idx_source] = string(stack_source[0 : len(stack_source)-size])
	crates[idx_target] = stack_target + currentStack

	return crates
}

func parseCrates(lines []string, nrCrates int) (result map[int]string) {
	result = make(map[int]string)

	for _, line := range lines {
		if string(line[1]) == "1" {
			continue
		}

		for idx := 1; idx <= nrCrates; idx += 1 {
			crate := string(line[idx*4-3])

			if crate != " " {
				result[idx] = crate + string(result[idx])
			}
		}
	}

	return
}

func read_file_as_blocks() (blocks [][]string) {
	var block_inputs []string = strings.Split(read_file(), "\n\n")

	for _, block_input := range block_inputs {
		blocks = append(blocks, strings.Split(block_input, "\n"))
	}
	return
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
