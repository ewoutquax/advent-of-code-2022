package puzzle

type Universe struct {
	blockedPositions map[string]bool // Position blocked in this universe, either by rock or by sand
	nrSands          int             // Number of pieces of sand, dropped into the universe
	floorY           int             // Y-location of the floor of the universe
}

func (u *Universe) fillWithSandUntilFloor() {
	var sand Sand

	for u.nrSands = -1; !sand.isOnFloor(u); u.nrSands++ {
		sand = spawnSand()
		sand.fallUntilRest(u)

		u.blockedPositions[sand.toKey()] = true
	}
}

func (u *Universe) fillWithSandUntilTop() {
	var sand Sand

	for u.nrSands = 0; !sand.isAtRestOnTop(u); u.nrSands++ {
		sand = spawnSand()
		sand.fallUntilRest(u)

		u.blockedPositions[sand.toKey()] = true
	}
}

// Soawn a new sand, and make it fall untill it becomes a rest.
// After it has come to rest, add it to the list of blockedLocations
func (u *Universe) dropSandTillRest() *Universe {
	sand := spawnSand()
	sand.fallUntilRest(u)

	u.blockedPositions[sand.position.toKey()] = true

	return u
}
