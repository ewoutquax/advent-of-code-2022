package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Solve(t *testing.T) {
	c := make(chan int)
	go solvePart1(c)
	result := <-c

	assert.Equal(t, 4418, result)
}

func TestPart2Solve(t *testing.T) {
	c := make(chan int)
	go solvePart2(c)
	result := <-c

	assert.Equal(t, 2486, result)
}
