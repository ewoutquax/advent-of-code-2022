package main

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

	assert.Equal(t, 24000, solvePart1(input))
}

func TestPart1Solve(t *testing.T) {
	assert.Equal(t, 71502, solvePart1([][]string{{}}))
}

 func TestPart2Examples(t *testing.T) {
	var input = [][]string{
		{"1000", "2000", "3000"},
		{"4000"},
		{"5000", "6000"},
		{"7000", "8000", "9000"},
		{"10000"},
	}
   assert.Equal(t, 45000, solvePart2(input))
 }

func TestPart2Solve(t *testing.T) {
  assert.Equal(t, 208191, solvePart2([][]string{{}}))
}
