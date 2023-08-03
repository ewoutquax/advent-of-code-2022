package puzzle

type Player struct {
	position Position // Position of the player inside the basin
}

// Get a string-representation of the position of the player
func (p Player) toKey() string {
	return p.position.toKey()
}

type Basin struct {
	nrStep            int             // Nr of current step, which dictates the location of blizzards and players
	players           []Player        // List with players with their location, by the current step
	blizzards         []*Blizzard     // List with references to all the blizzards inside the basin, by the current step
	wallLocations     map[string]bool // Keys of the locations of all found walls
	minX              int             // Location of the wall on the east side of the basin
	maxX              int             // Location of the wall on the west side of the basin
	minY              int             // Location of the wall on the north side of the basin
	maxY              int             // Location of the wall on the south side of the basin
	playerReachedGoal bool            // Did a single player from the current list reached the end-position
	moveToEnd         bool
	moveToBeginning   bool
	moveToEndAgain    bool
}

type DrawChar struct {
	char  string
	count int
}

func (b *Basin) moveBlizzards() *Basin {
	// var movedBlizzards = make([]*Blizzard, len(b.blizzards))

	// for idx, blizzard := range b.blizzards {
	for _, blizzard := range b.blizzards {
		switch blizzard.direction {
		case Up:
			if blizzard.position.y <= b.minY {
				blizzard.position.y = b.maxY
			} else {
				blizzard.position.y--
			}
		case Right:
			if blizzard.position.x >= b.maxX {
				blizzard.position.x = b.minX
			} else {
				blizzard.position.x++
			}
		case Down:
			if blizzard.position.y >= b.maxY {
				blizzard.position.y = b.minY
			} else {
				blizzard.position.y++
			}
		case Left:
			if blizzard.position.x <= b.minX {
				blizzard.position.x = b.maxX
			} else {
				blizzard.position.x--
			}
		default:
			panic("Blizzard has unknown direction")
		}
	}

	return b
}

func (b Basin) blizzardLocations() map[string]bool {
	var locs = make(map[string]bool)

	for _, blizzard := range b.blizzards {
		locs[blizzard.toKey()] = true
	}

	return locs
}

func (b Basin) blizzardLocations2() map[string]string {
	var locs = make(map[string]string)
	var value string

	for _, blizzard := range b.blizzards {
		switch blizzard.direction {
		case Up:
			value = "U"
		case Right:
			value = "R"
		case Down:
			value = "D"
		case Left:
			value = "L"
		}

		locs[blizzard.toKey()] = value
	}

	return locs
}

func (b *Basin) movePlayers() *Basin {
	var movedPlayers []Player
	var movedPlayerLocations = make(map[string]bool)
	blizzardLocations := b.blizzardLocations()

	var refLocs []Position = []Position{
		{-1, +0},
		{+0, -1},
		{+0, +0},
		{+0, +1},
		{+1, +0},
	}

	for _, player := range b.players {
		for _, refloc := range refLocs {
			movedPlayer := Player{position: Position{
				x: player.position.x + refloc.x,
				y: player.position.y + refloc.y,
			}}
			if !blizzardLocations[movedPlayer.toKey()] &&
				!b.wallLocations[movedPlayer.toKey()] &&
				!movedPlayerLocations[movedPlayer.toKey()] &&
				movedPlayer.position.y >= 0 &&
				movedPlayer.position.y <= b.maxY+1 {

				movedPlayers = append(movedPlayers, movedPlayer)
				movedPlayerLocations[movedPlayer.position.toKey()] = true

				if b.moveToEnd && movedPlayer.position.y > b.maxY {
					b.playerReachedGoal = true
				}

				if b.moveToBeginning && movedPlayer.position.y < b.minY {
					b.playerReachedGoal = true
				}

			}
		}
	}

	b.players = movedPlayers
	return b
}
