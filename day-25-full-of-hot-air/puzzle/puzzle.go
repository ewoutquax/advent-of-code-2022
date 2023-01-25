package puzzle

import "github.com/ewoutquax/advent-of-code-2022/day-25/snafu"

func SumSnafusAsSnafu(snafus []string) string {
	return snafu.ToSnafu(sumSnafus(snafus))
}

func sumSnafus(snafus []string) (total int) {
	for _, mySnafu := range snafus {
		total += snafu.ToDecimal(mySnafu)
	}

	return
}
