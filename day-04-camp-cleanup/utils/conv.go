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

func Unique(intSlice []int) (list []int) {
	var keys = make(map[int]bool)

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return
}
