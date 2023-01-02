package puzzle

import (
	"strings"

	"aoc.com/2022/day-06/utils"
)

func FindFirstMarker(input string, size int) int {
	for idx := size; idx < len(input); idx += 1 {
		message := strings.Split(input[idx-size:idx], "")

		if len(utils.Unique(message)) == size {
			return idx
		}
	}

	panic("No marker found")
}
