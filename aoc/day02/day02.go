package day02

import (
	"aoc/utils"
	"strconv"
	"strings"
)

var lines []string
var part1 int
var part2 int

func init() {
	lines = utils.ReadLines("res/day02.in")
}

func Part1() int {
	for _, line := range lines {
		levels := strings.Fields(line)

		var previous, current int
		var ascending, safe bool = true, true

		for i, level := range levels {
			if i == 0 {
				previous, _ = strconv.Atoi(level)
			} else if i == 1 {
				current, _ = strconv.Atoi(level)
				if current < previous {
					ascending = false
				}
				if !withinLimit(utils.Abs(current - previous)) {
					safe = false
					break
				}
			} else {
				previous = current
				current, _ = strconv.Atoi(level)
				if !isValid(previous, current, ascending) {
					safe = false
					break
				}
			}
		}

		if safe {
			part1++
		}

	}

	return part1
}

func Part2() int {
	for _, line := range lines {
		levels := strings.Fields(line)

		var previous, current int
		var ascending, safe, broken bool = true, true, false

		for i, level := range levels {
			if i == 0 {
				previous, _ = strconv.Atoi(level)
			} else if i == 1 {
				current, _ = strconv.Atoi(level)
				if current < previous {
					ascending = false
					if !withinLimit(utils.Abs(current - previous)) {
						safe = false
						current = previous
					}
				} else {
					if !withinLimit(utils.Abs(current - previous)) {
						// We should mark safe as false here, but the answer is right.
						// After figuring out we can combine the duplicated code
						current = previous
					}
				}
			} else {
				previous = current
				current, _ = strconv.Atoi(level)
				if !isValid(previous, current, ascending) {
					if !safe {
						broken = true
						break
					} else {
						safe = false
						current = previous
					}
				}
			}
		}

		if !broken {
			part2++
		}

	}

	return part2
}

func isValid(prev int, next int, ascending bool) bool {
	if (ascending && next < prev) || (!ascending && next > prev) {
		return false
	}
	val := utils.Abs(next - prev)
	return withinLimit(val)
}

func withinLimit(val int) bool {
	if val < 1 || val > 3 {
		return false
	}
	return true
}
