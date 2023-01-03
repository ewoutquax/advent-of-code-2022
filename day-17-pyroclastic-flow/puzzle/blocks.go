package puzzle

// Position of an edge, relative to the topleft corner of the block
type RelativePosition struct {
	position Position
}

type Block struct {
	height int                // Height of the block, used to determine the initial y-position of the block
	rocks  []RelativePosition // Position of each piece of rock; needed to populate the cave, so we can stack the blocks
}

func getBlock(blocks []Block, nrBlock int) Block {
	return blocks[nrBlock%len(blocks)]
}

func parseBlocks() (blocks []Block) {
	blocks = []Block{
		parseBlock(block1()),
		parseBlock(block2()),
		parseBlock(block3()),
		parseBlock(block4()),
		parseBlock(block5()),
	}
	return
}

func parseBlock(lines []string) (block Block) {
	block.height = len(lines)

	for relY, line := range lines {
		block.rocks = append(block.rocks, parseRocks(line, relY)...)
	}

	return
}

func parseRocks(line string, relY int) []RelativePosition {
	var positions = make([]RelativePosition, 0)

	for x := 0; x < len(line); x += 1 {
		if string(line[x]) == "#" {
			positions = append(positions, RelativePosition{position: Position{x: x, y: -relY}})
		}
	}

	return positions
}

func block1() []string {
	return []string{"####"}
}
func block2() []string {
	return []string{
		" # ",
		"###",
		" # ",
	}

}
func block3() []string {
	return []string{
		"  #",
		"  #",
		"###",
	}
}
func block4() []string {
	return []string{
		"#",
		"#",
		"#",
		"#",
	}
}
func block5() []string {
	return []string{
		"##",
		"##",
	}
}
