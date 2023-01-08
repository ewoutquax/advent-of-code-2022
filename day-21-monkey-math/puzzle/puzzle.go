package puzzle

import (
	"strings"

	"github.com/ewoutquax/advent-of-code-2022/utils"
)

type Operation uint8

const (
	Add Operation = iota
	Subtract
	Multiply
	Divide
)

type Monkey struct {
	name         string    // Name of the monkeu; used in a test, to see if the correct monkey is used
	hasRawValue  bool      // Used to determine wether the monkey has a raw value, or a math-calculation. In case of a math-calculation, solve it, store the value in this struct, and flip the boolean
	value        int       // The given raw value, or the solution of the math-operation
	referencedBy *Monkey   // When finding the path to the human, this field is used to backtrack the path to the rootMonkey
	nameLeft     string    // Name of the monkey with the left value of the math-operation
	nameRight    string    // Name of the monkey with the left value of the math-operation
	left         *Monkey   // Monkey with the left value of the math-operation
	operation    Operation // The operation to perform on the left- and rightvalue
	right        *Monkey   // Monkey with the left value of the math-operation
}

func (m *Monkey) getValue() int {
	var value int

	if !m.hasRawValue {
		leftValue := m.left.getValue()
		rightValue := m.right.getValue()

		switch m.operation {
		case Add:
			value = leftValue + rightValue
		case Subtract:
			value = leftValue - rightValue
		case Multiply:
			value = leftValue * rightValue
		case Divide:
			value = leftValue / rightValue
		}

		m.hasRawValue = true
		m.value = value
	}

	return m.value
}

func (m *Monkey) setNeededHumanValue() {
	path := findPathToMonkey(m, "humn")

	setRequiredValueForPathMonkey(path, 0)
}

func setRequiredValueForPathMonkey(path []*Monkey, requiredValue int) {
	var knownValue int
	var nextRequiredValue int

	currentMonkey := path[0]

	if len(path) == 1 {
		currentMonkey.value = requiredValue
	} else {
		nextMonkey := path[1]

		if currentMonkey.left.name == nextMonkey.name {
			knownValue = currentMonkey.right.getValue()
		} else {
			knownValue = currentMonkey.left.getValue()
		}

		if currentMonkey.name == "root" {
			nextRequiredValue = knownValue
		} else if currentMonkey.operation == Add && currentMonkey.left.name == nextMonkey.name {
			nextRequiredValue = requiredValue - currentMonkey.right.getValue()
		} else if currentMonkey.operation == Add && currentMonkey.right.name == nextMonkey.name {
			nextRequiredValue = requiredValue - currentMonkey.left.getValue()
		} else if currentMonkey.operation == Subtract && currentMonkey.left.name == nextMonkey.name {
			nextRequiredValue = requiredValue + currentMonkey.right.getValue()
		} else if currentMonkey.operation == Subtract && currentMonkey.right.name == nextMonkey.name {
			nextRequiredValue = currentMonkey.left.getValue() - requiredValue
		} else if currentMonkey.operation == Multiply && currentMonkey.left.name == nextMonkey.name {
			nextRequiredValue = requiredValue / currentMonkey.right.getValue()
		} else if currentMonkey.operation == Multiply && currentMonkey.right.name == nextMonkey.name {
			nextRequiredValue = requiredValue / currentMonkey.left.getValue()
		} else if currentMonkey.operation == Divide && currentMonkey.left.name == nextMonkey.name {
			nextRequiredValue = requiredValue * currentMonkey.right.getValue()
		} else if currentMonkey.operation == Divide && currentMonkey.right.name == nextMonkey.name {
			nextRequiredValue = currentMonkey.left.getValue() / requiredValue
		} else {
			panic("Don't know how to continue")
		}

		setRequiredValueForPathMonkey(path[1:], nextRequiredValue)
	}
}

func findPathToMonkey(startMonkey *Monkey, nameTarget string) (foundPath []*Monkey) {
	var targetMonkey *Monkey
	var paths = make(map[int][]*Monkey)
	firstPath := []*Monkey{startMonkey}
	paths[0] = firstPath

	for length := 0; targetMonkey == nil; length++ {
		for _, monkey := range paths[length] {
			if monkey.name == nameTarget {
				targetMonkey = monkey
			} else {
				if !monkey.hasRawValue && monkey.left.referencedBy == nil {
					monkey.left.referencedBy = monkey
					paths[length+1] = append(paths[length+1], monkey.left)
				}
				if !monkey.hasRawValue && monkey.right.referencedBy == nil {
					monkey.right.referencedBy = monkey
					paths[length+1] = append(paths[length+1], monkey.right)
				}
			}
		}
	}

	currentMonkey := targetMonkey
	for currentMonkey != startMonkey {
		foundPath = append([]*Monkey{currentMonkey}, foundPath...)
		currentMonkey = currentMonkey.referencedBy
	}
	foundPath = append([]*Monkey{startMonkey}, foundPath...)

	return
}

func GetValueRootMonkey(input []string) int {
	monkeys := parseInput(input)
	rootMonkey := monkeys["root"]

	return rootMonkey.getValue()
}

func GetNeededValueHuman(input []string) int {
	monkeys := parseInput(input)
	monkeys["root"].setNeededHumanValue()

	return monkeys["humn"].getValue()
}

func parseInput(lines []string) map[string]*Monkey {
	var out = make(map[string]*Monkey, len(lines))

	for _, line := range lines {
		monkey := parseLine(line)
		out[monkey.name] = &monkey
	}

	linkMonkeys(out)

	return out
}

func linkMonkeys(monkeys map[string]*Monkey) map[string]*Monkey {
	for _, monkey := range monkeys {
		if !monkey.hasRawValue {
			leftMonkey := monkeys[monkey.nameLeft]
			rightMonkey := monkeys[monkey.nameRight]
			monkey.left = leftMonkey
			monkey.right = rightMonkey
		}
	}

	return monkeys
}

func parseLine(line string) (monkey Monkey) {
	parts := strings.Split(line, ": ")

	monkey.name = parts[0]
	if !strings.Contains(parts[1], " ") {
		monkey.value = utils.ConvStrToI(parts[1])
		monkey.hasRawValue = true
	} else {
		monkey.hasRawValue = false
		monkey.nameLeft, monkey.operation, monkey.nameRight = parseMath(parts[1])
	}

	return
}

func parseMath(mathPart string) (nameLeft string, operation Operation, nameRight string) {
	parts := strings.Split(mathPart, " ")

	nameLeft = parts[0]
	nameRight = parts[2]

	switch parts[1] {
	case "+":
		operation = Add
	case "-":
		operation = Subtract
	case "*":
		operation = Multiply
	case "/":
		operation = Divide
	default:
		panic("Unknown operation")
	}

	return
}
