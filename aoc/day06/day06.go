package day06

import (
	"aoc/utils"
	"image"
)

var originalGrid = map[image.Point]rune{}
var originalShip image.Point
var directions = []image.Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var seen = map[image.Point]int{}
var total = 0

func init() {
	lines := utils.ReadLines("res/day06.in")

	for y, line := range lines {
		for x, letter := range line {
			originalGrid[image.Point{y, x}] = letter
			if letter == '^' {
				originalShip = image.Point{y, x}
			}
			total++
		}
	}
}

func run(part2 bool, obstacle image.Point) (int, bool) {
	grid, ship := map[image.Point]rune{}, originalShip
	for key, value := range originalGrid {
		grid[key] = value
	}
	if part2 {
		grid[obstacle] = '#'
	}
	i, index, counter, loop := 0, 0, 1, true
	direction := directions[index%4]
	grid[ship] = 'X'
	for {
		if i > total {
			break
		}
		next := ship.Add(direction)
		if string(grid[next]) == "." {
			ship = next
			grid[ship], seen[ship] = 'X', 1
			counter++
		} else if string(grid[next]) == "#" {
			index++
			direction = directions[index%4]
			continue
		} else if string(grid[next]) == "X" {
			ship = next
			grid[ship] = 'X'
		} else {
			loop = false
			break
		}
		i++
	}
	return counter, loop
}

func Part1() int {
	result, _ := run(false, image.Point{0, 0})
	return result
}

func Part2() int {
	counter := 0
	for point := range seen {
		_, loop := run(true, point)
		if loop {
			counter++
		}
	}
	return counter
}
