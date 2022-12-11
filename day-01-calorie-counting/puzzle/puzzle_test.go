package puzzle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	var input = [][]string{
		{"1000", "2000", "3000"},
		{"4000"},
		{"5000", "6000"},
		{"7000", "8000", "9000"},
		{"10000"},
	}

	assert.Equal(t, 24000, TotalTopN(input, 1))
}

func TestPart2Examples(t *testing.T) {
	var input = [][]string{
		{"1000", "2000", "3000"},
		{"4000"},
		{"5000", "6000"},
		{"7000", "8000", "9000"},
		{"10000"},
	}
	assert.Equal(t, 45000, TotalTopN(input, 3))
}
