package puzzle

import (
	"sort"
	"strings"

	utils "github.com/ewoutquax/aoc-go-utils"
)

func CountCompleteOverlaps(lines []string) (count int) {
	for _, line := range lines {
		if completeOverlap(line) {
			count += 1
		}
	}

	return
}

func CountPartlyOverlaps(lines []string) (count int) {
	for _, line := range lines {
		if hasOverlap(line) {
			count += 1
		}
	}

	return
}

func completeOverlap(line string) (result bool) {
	parts := strings.Split(line, ",")

	leftSections := expandSections(string(parts[0]))
	rightSections := expandSections(string(parts[1]))

	allSections := utils.Unique(append(leftSections, rightSections...))
	sort.Ints(allSections)

	return equalSlices(allSections, leftSections) ||
		equalSlices(allSections, rightSections)
}

func hasOverlap(line string) bool {
	parts := strings.Split(line, ",")

	var leftSections []int = expandSections(string(parts[0]))
	var rightSections []int = expandSections(string(parts[1]))

	allSections := append(leftSections, rightSections...)
	uniqueSections := utils.Unique(allSections)

	return len(allSections) != len(uniqueSections)
}

func expandSections(section string) (sections []int) {
	from_to := strings.Split(section, "-")

	from := utils.ConvStrToI(from_to[0])
	to := utils.ConvStrToI(from_to[1])

	for idx := from; idx <= to; idx += 1 {
		sections = append(sections, idx)
	}

	return
}

// Check two slices for identical content.
// Items are expected to be in the same order, so best to supply sorted slices.
func equalSlices(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
