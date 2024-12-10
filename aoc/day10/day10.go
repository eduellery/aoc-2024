package day10

import (
	"aoc/utils"
	"image"
)

var grid = map[image.Point]rune{}

func init() {
	lines := utils.ReadLines("res/day10.in")
	for y, line := range lines {
		for x, value := range line {
			grid[image.Point{x, y}] = value
		}
	}
}

func dfs(grid map[image.Point]rune, point image.Point, seen map[image.Point]bool) (score int) {
	if grid[point] == '9' {
		if seen[point] {
			return 0
		} else {
			seen[point] = true
			return 1
		}
	}
	for _, direction := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		if newPoint := point.Add(direction); grid[newPoint] == grid[point]+1 {
			score += dfs(grid, newPoint, seen)
		}
	}
	return score
}

func Part1() int {
	result := 0
	for point := range grid {
		if grid[point] == '0' {
			result += dfs(grid, point, map[image.Point]bool{})
		}
	}
	return result
}

func Part2() int {
	return 2
}
