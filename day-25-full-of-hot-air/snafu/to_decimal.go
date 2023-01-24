package snafu

import (
	"math"
	"strings"
)

func ToDecimal(snafu string) (decimal int) {
	var power, multiplier int
	var char string

	parts := strings.Split(snafu, "")

	for idx := len(parts) - 1; idx >= 0; idx-- {
		power = len(parts) - idx - 1
		multiplier = int(math.Pow(float64(5), float64(power)))

		char = string(snafu[idx])
		value := snafuBase(char)

		decimal += value * multiplier
	}

	return
}
