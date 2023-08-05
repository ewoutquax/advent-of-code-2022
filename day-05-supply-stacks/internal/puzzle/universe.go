package puzzle

import (
	"regexp"

	utils "github.com/ewoutquax/aoc-go-utils"
)

func (u *Universe) execInstructionsV1(instructions []string) {
	for _, instruction := range instructions {
		u.execInstructionV1(instruction)
	}

	return
}

func (u *Universe) execInstructionsV2(instructions []string) {
	for _, instruction := range instructions {
		u.execInstructionV2(instruction)
	}

	return
}

func (u *Universe) execInstructionV1(instruction string) {
	var idxSource, idxTarget int

	r, _ := regexp.Compile(`\Amove (\d+) from (\d+) to (\d+)\z`)
	matches := r.FindStringSubmatch(instruction)

	repeat := utils.ConvStrToI(matches[1])
	idxSource = utils.ConvStrToI(matches[2])
	idxTarget = utils.ConvStrToI(matches[3])

	for count := 0; count < repeat; count += 1 {
		stackSource := u.stack[idxSource]
		stackTarget := u.stack[idxTarget]

		crateCurrent := stackSource[len(stackSource)-1]

		u.stack[idxSource] = stackSource[0 : len(stackSource)-1]
		u.stack[idxTarget] = append(stackTarget, crateCurrent)
	}

	return
}

func (u *Universe) execInstructionV2(instruction string) {
	r, _ := regexp.Compile(`\Amove (\d+) from (\d+) to (\d+)\z`)
	matches := r.FindStringSubmatch(instruction)

	var size int = utils.ConvStrToI(matches[1])
	var idxSource int = utils.ConvStrToI(matches[2])
	var idxTarget int = utils.ConvStrToI(matches[3])

	stackSource := u.stack[idxSource]
	stackTarget := u.stack[idxTarget]

	stackCurrent := stackSource[len(stackSource)-(size):]

	u.stack[idxSource] = stackSource[0 : len(stackSource)-size]
	u.stack[idxTarget] = append(stackTarget, stackCurrent...)

	return
}

func (u Universe) generateSolution() (solution string) {
	for idx := 1; idx <= len(u.stack); idx += 1 {
		stack := u.stack[idx]
		crate := stack[len(stack)-1]
		solution += string(crate)
	}

	return
}
