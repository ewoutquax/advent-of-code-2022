package puzzle

import (
	"fmt"
	"strconv"
	"strings"
)

type RowHeights struct {
	row map[int]int
}

type RelativeHeight struct {
	identifier   string
	nrBlock      int
	heighestRock int
}

type Cave struct {
	heighestRock    int
	rocks           map[int]bool
	heights         RowHeights
	relativeHeights []RelativeHeight
}

// Information about the block that is currently falling
type CurrentBlock struct {
	position Position // topleft position of the current block
	block    Block    // The block that is currently falling
}

func HeightStackedBlocks(movements string, nrBlocks int) int {
	cave := stackBlocks(movements, nrBlocks)

	fmt.Println("cave.relativeHeights:", cave.relativeHeights)

	return cave.heighestRock
}

func stackBlocks(movements string, nrBlocks int) (cave Cave) {
	var canBlockFall bool
	var nrMovement int
	var blocks []Block = parseBlocks()

	cave.rocks = make(map[int]bool)
	cave.heights.row = make(map[int]int)

	for nrBlock := 0; nrBlock < nrBlocks; nrBlock += 1 {
		if nrBlock%len(blocks) == 0 && nrMovement%len(movements) == 1 {
			relHeight := cave.heights.convToIdentifier(cave.heighestRock)
			relHeight.nrBlock = nrBlock
			relHeight.heighestRock = cave.heighestRock

			cave.relativeHeights = append(cave.relativeHeights, relHeight)
		} else if nrBlock%len(blocks) == 0 {
			fmt.Println("modulo nrMovement upon starting with block0:", nrMovement%len(movements))
		}

		currentBlock := CurrentBlock{
			block: getBlock(blocks, nrBlock),
			position: Position{
				y: cave.heighestRock + getBlock(blocks, nrBlock).height + 3,
				x: 3,
			},
		}

		canBlockFall = true
		for canBlockFall {
			movementValue := getMovementValue(movements, nrMovement)
			nrMovement += 1

			if blockCanMoveByJet(cave, currentBlock, movementValue) {
				currentBlock.position.x += movementValue
			}

			canBlockFall = blockCanFall(cave, currentBlock)
			if canBlockFall {
				currentBlock.position.y -= 1
			}
		}

		placeBlockInCave(&cave, currentBlock)
		// drawCave(cave)
	}

	return
}

func drawCave(cave Cave) {
	fmt.Println("\nCAVE")
	for y := cave.heighestRock + 1; y > 0; y -= 1 {
		fmt.Print("|")
		for x := 1; x < 8; x += 1 {
			pos := Position{x: x, y: y}
			if cave.rocks[pos.toValue()] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}
	fmt.Print("---------\n\n")
}

func placeBlockInCave(cave *Cave, block CurrentBlock) {
	for _, relPos := range block.block.rocks {
		pos := Position{
			x: block.position.x + relPos.position.x,
			y: block.position.y + relPos.position.y,
		}

		if cave.heights.row[pos.x] < pos.y {
			cave.heights.row[pos.x] = pos.y
		}

		cave.rocks[pos.toValue()] = true
	}

	if cave.heighestRock < block.position.y {
		cave.heighestRock = block.position.y
	}
}

func blockCanMoveByJet(cave Cave, block CurrentBlock, movementValue int) bool {
	block.position.x += movementValue
	return blockRocksFitInCave(cave, block)
}

func blockCanFall(cave Cave, block CurrentBlock) bool {
	block.position.y -= 1
	return blockRocksFitInCave(cave, block)
}

func blockRocksFitInCave(cave Cave, block CurrentBlock) bool {
	for _, relPos := range block.block.rocks {
		pos := Position{
			x: block.position.x + relPos.position.x,
			y: block.position.y + relPos.position.y,
		}

		if pos.x <= 0 || pos.x >= 8 || pos.y <= 0 || cave.rocks[pos.toValue()] {
			return false
		}
	}

	return true
}

func (heights RowHeights) convToIdentifier(baseHeight int) RelativeHeight {
	var strRow []string
	for idx := 1; idx <= 7; idx += 1 {
		strRow = append(strRow, strconv.Itoa(heights.row[idx]-baseHeight))
	}

	return RelativeHeight{identifier: strings.Join(strRow, ",")}
}

func getMovementValue(movements string, nrMovement int) (value int) {
	movement := string(movements[nrMovement%len(movements)])

	switch movement {
	case "<":
		value = -1
	case ">":
		value = 1
	default:
		panic("Unknown movement")
	}

	return
}
