package puzzle

import (
	"fmt"
	"regexp"
	"strings"

	"aoc.com/2022/day-19/utils"
)

type Blueprint struct {
	nr                    int
	priceOreBotOre        int
	priceClayBotOre       int
	priceObsidianBotOre   int
	priceObsidianBotClay  int
	priceGeodeBotOre      int
	priceGeodeBotObsidian int
}

type State struct {
	minute int

	nrOreBots      int
	nrClayBots     int
	nrObsidianBots int
	nrGeodeBots    int

	blockBuildOreBot      bool
	blockBuildClayBot     bool
	blockBuildObsidianBot bool
	blockBuildGeodeBot    bool

	amtOre      int
	amtClay     int
	amtObsidian int
	amtGeode    int
}

func (s State) canBuildOreBot(b Blueprint) bool {
	return !s.blockBuildOreBot &&
		!s.canBuildGeodeBot(b) &&
		s.amtOre >= b.priceOreBotOre &&
		(s.nrOreBots < b.priceOreBotOre ||
			s.nrOreBots < b.priceClayBotOre ||
			s.nrOreBots < b.priceObsidianBotOre ||
			s.nrOreBots < b.priceGeodeBotOre)
}
func (s State) canBuildClayBot(b Blueprint) bool {
	return !s.blockBuildClayBot &&
		!s.canBuildGeodeBot(b) &&
		s.amtOre >= b.priceClayBotOre &&
		s.nrClayBots < b.priceObsidianBotClay

}
func (s State) canBuildObsidianBot(b Blueprint) bool {
	return !s.blockBuildObsidianBot &&
		!s.canBuildGeodeBot(b) &&
		s.amtOre >= b.priceObsidianBotOre && s.amtClay >= b.priceObsidianBotClay &&
		s.nrObsidianBots < b.priceGeodeBotObsidian
}
func (s State) canBuildGeodeBot(b Blueprint) bool {
	return !s.blockBuildGeodeBot && s.amtOre >= b.priceGeodeBotOre && s.amtObsidian >= b.priceGeodeBotObsidian
}

func (s State) key() string {
	return strings.Join(
		[]string{
			fmt.Sprintf("%03d", s.minute),
			fmt.Sprintf("%03d", s.nrOreBots),
			fmt.Sprintf("%03d", s.nrClayBots),
			fmt.Sprintf("%03d", s.nrObsidianBots),
			fmt.Sprintf("%03d", s.nrGeodeBots),
			boolToString(s.blockBuildOreBot),
			boolToString(s.blockBuildClayBot),
			boolToString(s.blockBuildObsidianBot),
			boolToString(s.blockBuildGeodeBot),
			fmt.Sprintf("%03d", s.amtOre),
			fmt.Sprintf("%03d", s.amtClay),
			fmt.Sprintf("%03d", s.amtObsidian),
			fmt.Sprintf("%03d", s.amtGeode),
		}, "/")
}

func (s State) value() string {
	return ""
}

func SumBlueprintQualities(lines []string) (total int) {
	for _, blueprint := range parseInput(lines) {
		total += findBlueprintQuality(blueprint)
	}

	return
}

func findBlueprintQuality(blueprint Blueprint) int {
	var states = make(map[int]map[string]string)
	var maxGeodes = 0

	state := State{
		nrOreBots:             1,
		blockBuildOreBot:      false,
		blockBuildClayBot:     false,
		blockBuildObsidianBot: false,
		blockBuildGeodeBot:    false,
	}

	states[1] = make(map[string]string)
	states[1][state.key()] = state.value()

	for minute := 1; minute <= 24; minute++ {
		fmt.Print("Blueprint ", blueprint.nr, ": minute ", minute, ": processing ", len(states[minute]), " scenarios\n")

		states[minute+1] = make(map[string]string)

		for key, value := range states[minute] {
			state := buildStateFromKeyAndValue(key, value)
			state.minute = minute
			newStates := execMinute(state, blueprint)

			for _, newState := range newStates {
				if maxGeodes <= newState.amtGeode+2 {
					if maxGeodes < newState.amtGeode {
						maxGeodes = newState.amtGeode
					}
					states[minute+1][newState.key()] = newState.value()
				}
			}
		}
	}

	fmt.Print("Max geodes for blueprint '", blueprint.nr, "': ", maxGeodes, "\n\n")

	return maxGeodes * blueprint.nr
}

func execMinute(state State, blueprint Blueprint) (outStates []State) {
	buildOreBot := state.canBuildOreBot(blueprint)
	buildClayBot := state.canBuildClayBot(blueprint)
	buildObsidianBot := state.canBuildObsidianBot(blueprint)
	buildGeodeBot := state.canBuildGeodeBot(blueprint)

	state.amtOre += state.nrOreBots
	state.amtClay += state.nrClayBots
	state.amtObsidian += state.nrObsidianBots
	state.amtGeode += state.nrGeodeBots

	if buildOreBot {
		outStates = append(outStates, doBuildOreBot(state, blueprint))
	}
	if buildClayBot {
		outStates = append(outStates, doBuildClayBot(state, blueprint))
	}
	if buildObsidianBot {
		outStates = append(outStates, doBuildObsidianBot(state, blueprint))
	}
	if buildGeodeBot {
		outStates = append(outStates, doBuildGeodeBot(state, blueprint))
	}

	// For each bot we can build, we also try the scenario where we didn't
	// And block the build in the next turn(s), to prevent obsolete scenarios which are less efficient
	if buildOreBot {
		state.blockBuildOreBot = true
	}
	if buildClayBot {
		state.blockBuildClayBot = true
	}
	if buildObsidianBot {
		state.blockBuildObsidianBot = true
	}
	if buildGeodeBot {
		state.blockBuildGeodeBot = true
	}

	// Return all new scenarios
	return append(outStates, state)
}

func doBuildOreBot(s State, b Blueprint) State {
	s.nrOreBots++
	s.amtOre -= b.priceOreBotOre

	s.blockBuildOreBot = false
	s.blockBuildClayBot = false
	s.blockBuildObsidianBot = false
	s.blockBuildGeodeBot = false

	return s
}

func doBuildClayBot(s State, b Blueprint) State {
	s.nrClayBots++
	s.amtOre -= b.priceClayBotOre

	s.blockBuildOreBot = false
	s.blockBuildClayBot = false
	s.blockBuildObsidianBot = false
	s.blockBuildGeodeBot = false

	return s
}

func doBuildObsidianBot(s State, b Blueprint) State {
	s.nrObsidianBots++
	s.amtOre -= b.priceObsidianBotOre
	s.amtClay -= b.priceObsidianBotClay

	s.blockBuildOreBot = false
	s.blockBuildClayBot = false
	s.blockBuildObsidianBot = false
	s.blockBuildGeodeBot = false

	return s
}

func doBuildGeodeBot(s State, b Blueprint) State {
	s.nrGeodeBots++
	s.amtOre -= b.priceGeodeBotOre
	s.amtObsidian -= b.priceGeodeBotObsidian

	s.blockBuildOreBot = false
	s.blockBuildClayBot = false
	s.blockBuildObsidianBot = false
	s.blockBuildGeodeBot = false

	return s
}

func buildStateFromKeyAndValue(key string, value string) (state State) {
	keyParts := strings.Split(key, "/")

	state.minute = utils.ConvStrToI(keyParts[0])
	state.nrOreBots = utils.ConvStrToI(keyParts[1])
	state.nrClayBots = utils.ConvStrToI(keyParts[2])
	state.nrObsidianBots = utils.ConvStrToI(keyParts[3])
	state.nrGeodeBots = utils.ConvStrToI(keyParts[4])
	state.blockBuildOreBot = stringToBool(keyParts[5])
	state.blockBuildClayBot = stringToBool(keyParts[6])
	state.blockBuildObsidianBot = stringToBool(keyParts[7])
	state.blockBuildGeodeBot = stringToBool(keyParts[8])
	state.amtOre = utils.ConvStrToI(keyParts[9])
	state.amtClay = utils.ConvStrToI(keyParts[10])
	state.amtObsidian = utils.ConvStrToI(keyParts[11])
	state.amtGeode = utils.ConvStrToI(keyParts[12])

	return
}

func parseInput(lines []string) (blueprints []Blueprint) {
	for _, line := range lines {
		blueprints = append(blueprints, parseLine(line))
	}

	return
}

func parseLine(line string) (blueprint Blueprint) {
	r, _ := regexp.Compile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian\.`)
	matches := r.FindStringSubmatch(line)

	blueprint.nr = utils.ConvStrToI(matches[1])
	blueprint.priceOreBotOre = utils.ConvStrToI(matches[2])
	blueprint.priceClayBotOre = utils.ConvStrToI(matches[3])
	blueprint.priceObsidianBotOre = utils.ConvStrToI(matches[4])
	blueprint.priceObsidianBotClay = utils.ConvStrToI(matches[5])
	blueprint.priceGeodeBotOre = utils.ConvStrToI(matches[6])
	blueprint.priceGeodeBotObsidian = utils.ConvStrToI(matches[7])

	return
}

func boolToString(b bool) string {
	if b {
		return "T"
	} else {
		return "F"
	}
}

func stringToBool(s string) bool {
	return s == "T"
}
