package puzzle

import (
	"aoc.com/2022/day-20/utils"
)

const ENCRYPTION_KEY int = 811589153

type Number struct {
	value int
	prev  *Number
	next  *Number
}

func GrooveHash(input []int) int {
	mixedList := mixNumbers(parseInput(input, 1))

	return sumGrooves(mixedList)
}

func GrooveHashHard(input []int) int {
	list := parseInput(input, ENCRYPTION_KEY)

	for run := 0; run < 10; run++ {
		list = mixNumbers(list)
	}

	return sumGrooves(list)
}

func sumGrooves(list []*Number) int {
	grooves := grooveNumbers(list)

	return grooves[0] + grooves[1] + grooves[2]
}

func grooveNumbers(list []*Number) (grooves [3]int) {
	var startNumber *Number

	for _, number := range list {
		if number.value == 0 {
			startNumber = number
			break
		}
	}

	for step := 0; step < 3; step++ {
		for idxNext := 0; idxNext < 1000; idxNext += 1 {
			startNumber = startNumber.next
		}
		grooves[step] = startNumber.value
	}

	return
}

func mixNumbers(list []*Number) []*Number {
	for _, currentNumber := range list {
		// fmt.Println("Mix number at index:", idx)
		for ctrMove := utils.Abs(currentNumber.value) % (len(list) - 1); ctrMove > 0; ctrMove-- {
			if currentNumber.value > 0 {
				// Move right (to next)
				currentNumber.prev.next = currentNumber.next
				currentNumber.next.prev = currentNumber.prev
				currentNumber.next.next.prev = currentNumber

				tmpNextNext := currentNumber.next.next
				currentNumber.next.next = currentNumber
				currentNumber.prev = currentNumber.next
				currentNumber.next = tmpNextNext
			} else {
				// Move left (to prev)
				currentNumber.next.prev = currentNumber.prev
				currentNumber.prev.next = currentNumber.next
				currentNumber.prev.prev.next = currentNumber

				tmpPrevPrev := currentNumber.prev.prev
				currentNumber.prev.prev = currentNumber
				currentNumber.next = currentNumber.prev
				currentNumber.prev = tmpPrevPrev
			}
		}
	}

	return list
}

func parseInput(lines []int, multiplier int) []*Number {
	list := make([]*Number, len(lines))

	for idx, value := range lines {
		number := Number{value: value * multiplier}
		list[idx] = &number
		if idx > 0 {
			prevNumber := list[idx-1]
			number.prev = prevNumber
			prevNumber.next = &number
		}
	}

	lastNumber := list[len(list)-1]
	lastNumber.next = list[0]
	list[0].prev = lastNumber

	return list
}
