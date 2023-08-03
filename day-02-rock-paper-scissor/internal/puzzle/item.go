package puzzle

func (i Item) determineForOutcome(requiredOutcome string) (other Item) {
	if requiredOutcome == "X" {
		// Lose turn
		other = i.requiredToLose()
	} else if requiredOutcome == "Z" {
		// Win turn
		other = i.requiredToWin()
	} else {
		// draw
		other = i
	}

	return
}

func (mine Item) scoreRound(his Item) (score int) {
	if his == mine {
		score = 3
	} else if his.isRock() && mine.isPaper() ||
		his.isPaper() && mine.isScissor() ||
		his.isScissor() && mine.isRock() {
		score = 6
	}

	return
}

func (i Item) requiredToWin() Item {
	return map[Item]Item{
		Rock:    Paper,
		Paper:   Scissor,
		Scissor: Rock,
	}[i]
}

func (i Item) requiredToLose() Item {
	return map[Item]Item{
		Rock:    Scissor,
		Paper:   Rock,
		Scissor: Paper,
	}[i]
}

func (i Item) value() int {
	return map[Item]int{
		Rock:    1,
		Paper:   2,
		Scissor: 3,
	}[i]
}

func (i Item) isRock() bool {
	return i == Rock
}

func (i Item) isPaper() bool {
	return i == Paper
}

func (i Item) isScissor() bool {
	return i == Scissor
}
