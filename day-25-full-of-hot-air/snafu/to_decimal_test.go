package snafu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDecimal(t *testing.T) {
	checks := map[string]int{
		"2=-01":         976,
		"1":             1,
		"2":             2,
		"1=":            3,
		"1-":            4,
		"10":            5,
		"11":            6,
		"12":            7,
		"2=":            8,
		"2-":            9,
		"20":            10,
		"1=0":           15,
		"1-0":           20,
		"1=11-2":        2022,
		"1-0---0":       12345,
		"1121-1110-1=0": 314159265,
	}

	for snafu, decimal := range checks {
		assert.Equal(t, decimal, ToDecimal(snafu))
	}
}
