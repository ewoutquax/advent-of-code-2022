package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	blueprint := parseLine(input()[0])

	assert.Equal(t, 1, blueprint.nr)
	assert.Equal(t, 4, blueprint.priceOreBotOre)
	assert.Equal(t, 2, blueprint.priceClayBotOre)
	assert.Equal(t, 3, blueprint.priceObsidianBotOre)
	assert.Equal(t, 14, blueprint.priceObsidianBotClay)
	assert.Equal(t, 2, blueprint.priceGeodeBotOre)
	assert.Equal(t, 7, blueprint.priceGeodeBotObsidian)
}

func TestFindBlueprintQuality(t *testing.T) {
	blueprint := parseLine(input()[0])

	quality := findBlueprintQuality(blueprint)
	assert.Equal(t, 9, quality)
}

func TestSumBlueprintQualities(t *testing.T) {
	sum := SumBlueprintQualities(input())
	assert.Equal(t, 33, sum)
}

func input() []string {
	return []string{
		"Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.",
		"Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.",
	}
}
