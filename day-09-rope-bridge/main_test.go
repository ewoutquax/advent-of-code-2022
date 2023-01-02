package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	count := solvePart1()
	assert.Equal(t, 6057, count)
}

func TestSolvePart2(t *testing.T) {
	count := solvePart2()
	assert.Equal(t, 2514, count)
}
