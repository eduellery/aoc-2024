package day04

import (
	"aoc/utils"
	"image"
	"strings"
)

var grid = map[image.Point]rune{}

func init() {
	lines := utils.ReadLines("res/day04.in")

	for y, line := range lines {
		for x, letter := range line {
			grid[image.Point{x, y}] = letter
		}
	}

}

func adj(point image.Point, length int) []string {
	delta := []image.Point{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
		{-1, -1}, {1, -1}, {1, 1}, {-1, 1},
	}

	words := make([]string, len(delta))

	for i, direction := range delta {
		for n := range length {
			words[i] += string(grid[point.Add(direction.Mul(n))])
		}
	}
	return words
}

func Part1() int {
	part1 := 0
	for point := range grid {
		part1 += strings.Count(strings.Join(adj(point, 4), " "), "XMAS")
	}
	return part1
}
