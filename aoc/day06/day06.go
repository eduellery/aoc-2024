package day06

import (
	"aoc/utils"
	"image"
)

var grid = map[image.Point]rune{}
var ship image.Point
var index int
var directions = []image.Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func init() {
	lines := utils.ReadLines("res/day06.in")

	for y, line := range lines {
		for x, letter := range line {
			grid[image.Point{y, x}] = letter
			if letter == '^' {
				ship = image.Point{y, x}
			}
		}
	}
}

func run() int {
	counter := 1
	direction := directions[index%4]
	grid[ship] = 'X'
	for {
		next := ship.Add(direction)
		if string(grid[next]) == "." {
			ship = next
			grid[ship] = 'X'
			counter++
		} else if string(grid[next]) == "#" {
			index++
			direction = directions[index%4]
			continue
		} else if string(grid[next]) == "X" {
			ship = next
			grid[ship] = 'X'
		} else {
			break
		}
	}
	return counter
}

func Part1() int {
	return run()
}
