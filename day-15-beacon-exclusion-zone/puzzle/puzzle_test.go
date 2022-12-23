package puzzle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	sensor := parseLine(input()[0])

	assert.Equal(t, 2, sensor.location.x)
	assert.Equal(t, 18, sensor.location.y)
	assert.Equal(t, -2, sensor.nearestBeacon.x)
	assert.Equal(t, 15, sensor.nearestBeacon.y)
	assert.Equal(t, 7, sensor.ManhattanDistance(sensor.nearestBeacon))
}

func TestExcludesOnRow(t *testing.T) {
	var sensor Sensor
	var result []int

	sensor = parseLine("Sensor at x=2, y=0: closest beacon is at x=2, y=10")
	result = excludesOnRow(sensor, 10)
	assert.Equal(t, []int{2}, result)

	sensor = parseLine("Sensor at x=3, y=0: closest beacon is at x=2, y=10")
	result1 := excludesOnRow(sensor, 10)
	assert.Equal(t, []int{2, 3, 4}, result1)

	sensor = parseLine("Sensor at x=1, y=0: closest beacon is at x=2, y=10")
	result2 := excludesOnRow(sensor, 10)
	assert.Equal(t, []int{0, 1, 2}, result2)
}

func TestAllExcludesOnRow(t *testing.T) {
	result := AllExcludesOnRow(input(), 10)
	assert.Equal(t, 26, result)
}

func TestFindIntersection(t *testing.T) {
	sensor1 := parseLine(input()[0])

	intersection := findIntersection(sensor1, Line1, sensor1, Line2)
	assert.Equal(t, 2, intersection.x)
	assert.Equal(t, 26, intersection.y)
}

func TestFindAllIntersections(t *testing.T) {
	var sensors []Sensor

	for _, line := range input() {
		sensors = append(sensors, parseLine(line))
	}

	intersections := findAllIntersections(sensors)

	assert.Equal(t, 14*13*2*2, len(intersections))
}

func TestFindOpenLocation(t *testing.T) {
	intersections := findOpenLocation(input())

	assert.Equal(t, 1, len(intersections))
}

func TestFindOpenLocationHash(t *testing.T) {
	hash := FindOpenLocationHash(input())

	assert.Equal(t, 56000011, hash)
}

func input() []string {
	return []string{
		"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
		"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
		"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
		"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
		"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
		"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
		"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
		"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
		"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
		"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
		"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
		"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
		"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
		"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
	}
}
