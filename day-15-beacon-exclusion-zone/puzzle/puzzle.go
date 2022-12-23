package puzzle

import (
	"fmt"
	"regexp"

	"aoc.com/2022/day-15/utils"
)

// The covered area of a sensor is a diamond-shape
// The unmarked location will be on the lines around the area.
// To narrow down the possible locations, we'll calculate the intersections of all the lines of all sensors
// and of each found intersection, we'll check if it's outside the manhattan-distance of each sensor

// ^      ^
// |     / \
// |    1   2
// |   /     \
// |  <       >
// |   \     /
// |    4   3
// |     \ /
// y      V
// |
// +-x------->

// All the vectors 2 and 4 of all sensors are parralel, and won't have intersections.
// Some for vectors 1 and 3.
// Calculate the rest

// Vector is side1 = maxY - sensor.y / sensor.x - minX
// Vector is side2 = sensor.y - maxY / maxX - sensor.x
// Vector is side3 = sensor.y - minY / maxX - sensor.x
// Vector is side4 = minY - sensor.y / sensor.x - minX
// Lift for each side ('b' is the lift): y=ax+b => b = y - ax

// Intersections are found via:
//           y = a1.x + b1 =>
//           y = a2.x + b2 =>
//   a1.x + b1 = a2.x + b2 =>
// a1.x - a2.x = b2 - b1   =>
// (a1 - a2).x = b2 - b1   =>
//           x = (b2 - b1) / (a1 - a2)

type Vector uint

const (
	Line1 Vector = iota
	Line2
	Line3
	Line4
)

type Location struct {
	x int
	y int
}

type Sensor struct {
	location      Location
	nearestBeacon Location
}

func FindOpenLocationHash(lines []string) int {
	location := findOpenLocation(lines)

	return location[0].x*4000000 + location[0].y
}

func findOpenLocation(lines []string) (validLocations []Location) {
	var sensors []Sensor

	for _, line := range lines {
		sensors = append(sensors, parseLine(line))
	}

	locations := findAllIntersections(sensors)

	for _, location := range locations {
		outOfRange := true
		for idx := 0; outOfRange && idx < len(sensors); idx += 1 {
			sensor := sensors[idx]
			outOfRange = sensor.ManhattanDistance(sensor.nearestBeacon) < sensor.ManhattanDistance(location)
		}

		if outOfRange && location.x >= 0 && location.x <= 4000000 && location.y >= 0 && location.y <= 4000000 {
			validLocations = append(validLocations, location)
		}
	}

	fmt.Println("Valid locations: ", validLocations)

	return validLocations
}

func findAllIntersections(sensors []Sensor) (intersections []Location) {
	for _, sensor1 := range sensors {
		for _, sensor2 := range sensors {
			if sensor1 != sensor2 {
				intersections = append(intersections, findIntersection(sensor1, Line1, sensor2, Line2))
				intersections = append(intersections, findIntersection(sensor1, Line1, sensor2, Line4))
				intersections = append(intersections, findIntersection(sensor1, Line3, sensor2, Line2))
				intersections = append(intersections, findIntersection(sensor1, Line3, sensor2, Line4))
			}
		}
	}

	return
}

func findIntersection(sensor1 Sensor, line1 Vector, sensor2 Sensor, line2 Vector) (l Location) {
	vector1 := getVector(sensor1, line1)
	vector2 := getVector(sensor2, line2)

	l.x = (sensor2.Lift(line2) - sensor1.Lift(line1)) / (vector1 - vector2)
	l.y = int(vector1*l.x + sensor1.Lift(line1))

	return
}

func AllExcludesOnRow(lines []string, row int) int {
	var sensors []Sensor
	var excludes []int

	for _, line := range lines {
		sensors = append(sensors, parseLine(line))
	}

	for _, s := range sensors {
		excludes = append(excludes, excludesOnRow(s, row)...)
	}

	uniques := utils.Unique(excludes)
	for _, s := range sensors {
		if s.nearestBeacon.y == row {
			uniques = remove(uniques, s.nearestBeacon.x)
		}
	}

	return len(uniques)
}

func excludesOnRow(sensor Sensor, row int) (excludes []int) {
	distanceY := utils.Abs(sensor.location.y - row)

	md := sensor.ManhattanDistance(sensor.nearestBeacon)
	for x := sensor.location.x - md; x <= sensor.location.x+md; x += 1 {
		if utils.Abs(x-sensor.location.x)+distanceY <= sensor.ManhattanDistance(sensor.nearestBeacon) {
			excludes = append(excludes, x)
		}
	}

	return
}

func parseLine(line string) (sensor Sensor) {
	r, _ := regexp.Compile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	matches := r.FindStringSubmatch(line)

	sensor.location.x = utils.ConvStrToI(matches[1])
	sensor.location.y = utils.ConvStrToI(matches[2])
	sensor.nearestBeacon.x = utils.ConvStrToI(matches[3])
	sensor.nearestBeacon.y = utils.ConvStrToI(matches[4])

	return
}

func (s Sensor) ManhattanDistance(target Location) int {
	return utils.Abs(s.location.x-target.x) + utils.Abs(s.location.y-target.y)
}

func remove(s []int, e int) []int {
	var idx int = -1

	for i, current := range s {
		if current == e {
			idx = i
		}
	}

	if idx != -1 {
		s[idx] = s[len(s)-1]
		s = s[:len(s)-1]
	}

	return s
}

func getVector(s Sensor, line Vector) (v int) {
	switch line {
	case Line1:
		v = s.Vector1()
	case Line2:
		v = s.Vector2()
	case Line3:
		v = s.Vector3()
	case Line4:
		v = s.Vector4()
	}

	return
}

// Lift for each side is: y=ax+b => b = y - ax
func (s Sensor) Lift(line Vector) (l int) {
	switch line {
	case Line1:
		l = s.MaxY() - s.Vector1()*s.location.x
	case Line2:
		l = s.location.y - s.Vector2()*s.MaxX()
	case Line3:
		l = s.location.y - s.Vector3()*s.MaxX()
	case Line4:
		l = s.MinY() - s.Vector4()*s.location.x
	}
	return
}

// Vector is side1 = maxY - sensor.y / sensor.x - minX
func (s Sensor) Vector1() int {
	return (s.MaxY() - s.location.y) / (s.location.x - s.MinX())
}

// Vector is side2 = sensor.y - maxY / maxX - sensor.x
func (s Sensor) Vector2() int {
	return (s.location.y - s.MaxY()) / (s.MaxX() - s.location.x)
}

// Vector is side3 = sensor.y - minY / maxX - sensor.x
func (s Sensor) Vector3() int {
	return (s.location.y - s.MinY()) / (s.MaxX() - s.location.x)
}

// Vector is side4 = minY - sensor.y / sensor.x - minX
func (s Sensor) Vector4() int {
	return (s.MinY() - s.location.y) / (s.location.x - s.MinX())
}

func (s Sensor) MinX() int {
	return s.location.x - s.ManhattanDistance(s.nearestBeacon) - 1
}

func (s Sensor) MaxX() int {
	return s.location.x + s.ManhattanDistance(s.nearestBeacon) + 1
}

func (s Sensor) MinY() int {
	return s.location.y - s.ManhattanDistance(s.nearestBeacon) - 1
}

func (s Sensor) MaxY() int {
	return s.location.y + s.ManhattanDistance(s.nearestBeacon) + 1
}
