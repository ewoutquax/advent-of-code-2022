package puzzle

type Elf struct {
	position Position
	moveTo   *Position
}

func (e *Elf) setMoveToNorth() { e.moveTo = &Position{x: e.position.x + 0, y: e.position.y - 1} }
func (e *Elf) setMoveToEast()  { e.moveTo = &Position{x: e.position.x + 1, y: e.position.y + 0} }
func (e *Elf) setMoveToSouth() { e.moveTo = &Position{x: e.position.x + 0, y: e.position.y + 1} }
func (e *Elf) setMoveToWest()  { e.moveTo = &Position{x: e.position.x - 1, y: e.position.y + 0} }
func (e *Elf) setMoveToCurrentLocation() {
	e.moveTo = &Position{x: e.position.x + 0, y: e.position.y + 0}
}

func (elf Elf) checkNorth(elvePositions *map[string]*Elf) bool {
	var positions []Position = []Position{
		{x: elf.position.x - 1, y: elf.position.y - 1},
		{x: elf.position.x + 0, y: elf.position.y - 1},
		{x: elf.position.x + 1, y: elf.position.y - 1},
	}

	return emptyPossiblePositions(positions, elvePositions)
}

func (elf Elf) checkEast(elvePositions *map[string]*Elf) bool {
	var positions []Position = []Position{
		{x: elf.position.x + 1, y: elf.position.y - 1},
		{x: elf.position.x + 1, y: elf.position.y + 0},
		{x: elf.position.x + 1, y: elf.position.y + 1},
	}

	return emptyPossiblePositions(positions, elvePositions)
}

func (elf Elf) checkSouth(elvePositions *map[string]*Elf) bool {
	var positions []Position = []Position{
		{x: elf.position.x - 1, y: elf.position.y + 1},
		{x: elf.position.x + 0, y: elf.position.y + 1},
		{x: elf.position.x + 1, y: elf.position.y + 1},
	}

	return emptyPossiblePositions(positions, elvePositions)
}

func (elf Elf) checkWest(elvePositions *map[string]*Elf) bool {

	var positions []Position = []Position{
		{x: elf.position.x - 1, y: elf.position.y - 1},
		{x: elf.position.x - 1, y: elf.position.y + 0},
		{x: elf.position.x - 1, y: elf.position.y + 1},
	}

	return emptyPossiblePositions(positions, elvePositions)
}

func (e Elf) shouldMove(elves map[string]*Elf) bool {
	var positions []Position = []Position{
		{x: e.position.x - 1, y: e.position.y - 1},
		{x: e.position.x - 1, y: e.position.y + 0},
		{x: e.position.x - 1, y: e.position.y + 1},
		{x: e.position.x + 0, y: e.position.y - 1},
		{x: e.position.x + 0, y: e.position.y + 1},
		{x: e.position.x + 1, y: e.position.y - 1},
		{x: e.position.x + 1, y: e.position.y + 0},
		{x: e.position.x + 1, y: e.position.y + 1},
	}

	for _, position := range positions {
		if elves[position.toKey()] != nil {
			return true
		}
	}

	return false
}

func emptyPossiblePositions(possiblePositions []Position, elvePositions *map[string]*Elf) bool {
	for _, possiblePosition := range possiblePositions {
		tmp := *elvePositions
		if tmp[possiblePosition.toKey()] != nil {

			return false
		}
	}

	return true
}
