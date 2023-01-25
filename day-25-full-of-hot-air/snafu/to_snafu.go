package snafu

import (
	"strconv"
	"strings"
)

func ToSnafu(number int) string {
	return strings.Join(collectParts(number), "")
}

func collectParts(remainder int) []string {
	var out string
	var toConvert int

	if remainder == 0 {
		return []string{""}
	}

	toConvert = remainder % 5

	if toConvert == 3 {
		remainder += 2
		out = "="
	} else if toConvert == 4 {
		remainder += 1
		out = "-"
	} else {
		remainder -= toConvert
		out = strconv.Itoa(toConvert)
	}

	return append(collectParts(remainder/5), out)
}
