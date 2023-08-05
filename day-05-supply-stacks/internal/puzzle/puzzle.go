package puzzle

func ResolveInstructionsV1(blocks [][]string, nrCratesWide int) string {
	var crates []string = blocks[0]
	var instructions []string = blocks[1]

	universe := parseInput(crates, nrCratesWide)
	universe.execInstructionsV1(instructions)

	return universe.generateSolution()
}

func ResolveInstructionsV2(blocks [][]string, nrCratesWide int) string {
	var crates []string = blocks[0]
	var instructions []string = blocks[1]

	universe := parseInput(crates, nrCratesWide)
	universe.execInstructionsV2(instructions)

	return universe.generateSolution()
}

func parseInput(lines []string, nrCrates int) (universe Universe) {
	universe.stack = make(map[int][]Crate)

	for _, line := range lines {
		if string(line[1]) == "1" {
			// indication that all crate-lines have been parsed
			continue
		}

		for idx := 1; idx <= nrCrates; idx += 1 {
			char := string(line[idx*4-3])

			if char != " " {
				crate := []Crate{Crate(char)}
				universe.stack[idx] = append(crate, universe.stack[idx]...)
			}
		}
	}

	return
}
