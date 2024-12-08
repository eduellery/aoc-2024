package day08

import (
	"aoc/utils"
	"image"
)

var antennas = map[rune][]image.Point{}
var bounds = map[image.Point]bool{}

func init() {
	lines := utils.ReadLines("res/day08.in")

	for y, line := range lines {
		for x, value := range line {
			bounds[image.Point{x, y}] = true
			if value != '.' {
				list := antennas[value]
				antennas[value] = append(list, image.Point{x, y})
			}
		}
	}
}

func Part1() int {
	var antinodes = map[image.Point]struct{}{}
	for _, locations := range antennas {
		if len(locations) < 2 {
			continue
		}
		for _, locationA := range locations {
			for _, locationB := range locations {
				if locationA == locationB {
					continue
				}
				// Add the dist between pointB and pointA to pointB,
				// which will create its counterpart (antinode)
				test := locationB.Add(locationB.Sub(locationA))
				if bounds[test] {
					antinodes[test] = struct{}{}
				}
			}
		}
	}
	return len(antinodes)
}

func Part2() int {
	var antinodes = map[image.Point]struct{}{}
	for _, locations := range antennas {
		if len(locations) < 2 {
			continue
		}
		for _, locationA := range locations {
			for _, locationB := range locations {
				if locationA == locationB {
					continue
				}
				// Given a location, keep adding the dist to another location
				// until it gets out of bounds. All generated locations are antinodes
				for test := locationB.Sub(locationA); bounds[locationB]; locationB = locationB.Add(test) {
					antinodes[locationB] = struct{}{}
				}
			}
		}
	}
	return len(antinodes)
}
