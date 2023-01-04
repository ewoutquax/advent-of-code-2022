package main

import (
	"fmt"

	"aoc.com/2022/day-18/puzzle"
	"aoc.com/2022/day-18/utils"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	go solvePart1(chan1)
	go solvePart2(chan2)
	part1 := <-chan1
	part2 := <-chan2

	fmt.Println("Solution of part1:", part1)
	fmt.Println("Solution of part2:", part2)
}

func solvePart1(c chan int) {
	c <- puzzle.CountExposedSides(utils.ReadFileAsLines())
}

func solvePart2(c chan int) {
	c <- puzzle.CountSidesExposedToFreeAir(utils.ReadFileAsLines())
}
