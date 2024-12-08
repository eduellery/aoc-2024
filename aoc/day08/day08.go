package day08

import (
	"aoc/utils"
	"image"
)

var width = 0
var height = 0
var antennas = map[rune][]image.Point{}

func init() {
	lines := utils.ReadLines("res/day08.in")

	for y, line := range lines {
		for x, value := range line {
			if value != '.' {
				list := antennas[value]
				antennas[value] = append(list, image.Point{x, y})
			}
		}
	}

	height = len(lines)
	width = len(lines[0])
}

func countValidAntennas(locationA, locationB image.Point, antinodes map[image.Point]int) {
	x1, x2 := locationA.X, locationB.X
	y1, y2 := locationA.Y, locationB.Y
	dx := utils.Abs(x1 - x2)
	dy := utils.Abs(y1 - y2)
	if x1 < x2 {
		normalizedXCount(x1, x2, y1, y2, dx, dy, antinodes)
	} else {
		normalizedXCount(x2, x1, y2, y1, dx, dy, antinodes)
	}
}

func normalizedXCount(x1, x2, y1, y2, dx, dy int, antinodes map[image.Point]int) {
	if y1 < y2 {
		normalizedCoordinatesCount(x1-dx, y1-dy, x2+dx, y2+dy, antinodes)
	} else {
		normalizedCoordinatesCount(x1-dx, y1+dy, x2+dx, y2-dy, antinodes)
	}
}

func normalizedCoordinatesCount(x3, y3, x4, y4 int, antinodes map[image.Point]int) {
	if x3 >= 0 && x3 < width && y3 >= 0 && y3 < height {
		antinodes[image.Point{x3, y3}] = antinodes[image.Point{x3, y3}] + 1
	}
	if x4 >= 0 && x4 < width && y4 >= 0 && y4 < height {
		antinodes[image.Point{x4, y4}] = antinodes[image.Point{x3, y3}] + 1
	}
}

func Part1() int {
	var antinodes = map[image.Point]int{}
	for _, locations := range antennas {
		if len(locations) < 2 {
			continue
		}
		for i, locationA := range locations {
			for _, locationB := range locations[i+1:] {
				countValidAntennas(locationA, locationB, antinodes)
			}
		}
	}
	return len(antinodes)
}
