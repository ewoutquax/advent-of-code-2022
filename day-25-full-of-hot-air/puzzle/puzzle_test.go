package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumSnafu(t *testing.T) {
	assert.Equal(t, 4890, sumSnafus(input()))
}

func TestSumSnafusAsSnafu(t *testing.T) {
	assert.Equal(t, "2=-1=0", SumSnafusAsSnafu(input()))
}

func input() []string {
	return []string{
		"1=-0-2",
		"12111",
		"2=0=",
		"21",
		"2=01",
		"111",
		"20012",
		"112",
		"1=-1=",
		"1-12",
		"12",
		"1=",
		"122",
	}
}
