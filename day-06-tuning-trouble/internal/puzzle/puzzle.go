package puzzle

import (
	"strings"

	utils "github.com/ewoutquax/aoc-go-utils"
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
