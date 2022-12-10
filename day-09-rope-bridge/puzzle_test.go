package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Examples(t *testing.T) {
	count := solvePart1(input())
	assert.Equal(t, 13, count)
}

func TestPart2Examples(t *testing.T) {
	count := solvePart2(inputPart2())
	assert.Equal(t, 36, count)
}

func TestSolvePart1(t *testing.T) {
	count := solvePart1([]string{})
	assert.Equal(t, 6057, count)
}

func TestSolvePart2(t *testing.T) {
	count := solvePart2([]string{})
	assert.Equal(t, 2514, count)
}

func input() []string {
	return []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}
}

func inputPart2() []string {
	return []string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	}
}
