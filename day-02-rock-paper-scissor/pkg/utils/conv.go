package utils

import (
	"strconv"
)

// Convert a string to an int, without the nasty error-check
func ConvStrToI(s string) (i int) {
	i, err := strconv.Atoi(s)
	check(err)
	return
}
