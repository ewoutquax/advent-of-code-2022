package utils

import (
	"strconv"
)

// Convert a string to an int, without the nasty error-check
func ConvStrToI(s string) (i int) {
	i, err := strconv.Atoi(s)
	Check(err)
	return
}

func Unique(a []string) (out []string) {
	var keys = make(map[string]bool)

	for _, val := range a {
		if !keys[val] {
			keys[val] = true
			out = append(out, val)
		}
	}

	return
}
