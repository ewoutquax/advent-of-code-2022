package puzzle

const SPAWN_LOCATION_X int = 500 // On which x-axis do new sands spawn
const SPAWN_LOCATION_Y int = 0   // On which y-axis do new sands spawn

type Sand struct {
	position Position
}

// Let the current sand fall, until it becomes a full rest, or drops into the void
func (s *Sand) fallUntilRest(u *Universe) {
	for !s.isAtRest(u) {
		switch true {
		case s.canFallDown(u):
			s.position.y += 1
		case s.canFallLeft(u):
			s.position.x -= 1
			s.position.y += 1
		case s.canFallRight(u):
			s.position.x += 1
			s.position.y += 1
		}
	}
}

// Will return true, if the postion left, right and direct below the current position are blocked positions
func (s Sand) isAtRest(u *Universe) bool {
	return s.isOnFloor(u) || !(s.canFallLeft(u) || s.canFallDown(u) || s.canFallRight(u))
}

// Check for a blocked position, directly under the position of the current sand
func (s Sand) canFallDown(u *Universe) bool {
	return !u.blockedPositions[Position{
		x: s.position.x,
		y: s.position.y + 1,
	}.toKey()]
}

// Check for a blocked position, on the left under the position of the current sand
func (s Sand) canFallLeft(u *Universe) bool {
	return !u.blockedPositions[Position{
		x: s.position.x - 1,
		y: s.position.y + 1,
	}.toKey()]
}

// Check for a blocked position, on the right under the position of the current sand
func (s Sand) canFallRight(u *Universe) bool {
	return !u.blockedPositions[Position{
		x: s.position.x + 1,
		y: s.position.y + 1,
	}.toKey()]
}

// Has the current sand fallen into the void / on to the floor?
func (s Sand) isOnFloor(u *Universe) bool {
	return s.position.y+1 >= u.floorY
}

// Is the current sand unable to fall, while still at the top row?
func (s Sand) isAtRestOnTop(u *Universe) bool {
	return s.isAtRest(u) && s.position.y == 0
}

func (s Sand) toKey() string {
	return s.position.toKey()
}

// Spawn a new sand at the default location
func spawnSand() Sand {
	return Sand{
		position: Position{
			x: SPAWN_LOCATION_X,
			y: SPAWN_LOCATION_Y,
		},
	}
}
