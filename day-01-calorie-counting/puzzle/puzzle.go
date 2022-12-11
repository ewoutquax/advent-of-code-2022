package puzzle

import (
	"aoc.com/2022/day-01/utils"
	"sort"
)

func TotalTopN(blocks [][]string, top_N int) (total int) {
	var calories []int
	for _, lines := range blocks {
		elf := 0
		for _, line := range lines {
			elf += utils.ConvStrToI(line)
		}
		calories = append(calories, elf)
	}

	sort.Ints(calories)

	for i := 1; i <= top_N; i += 1 {
		total += calories[len(calories)-i]
	}

	return
}
