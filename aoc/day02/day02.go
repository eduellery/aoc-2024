package day02

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

var part1 int
var part2 int

func Day02() {
	lines, err := utils.ReadLines("res/day02.in")
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

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
					if !withinLimit(current, previous) {
						safe = false
						break
					}
				} else {
					if !withinLimit(previous, current) {
						safe = false
						break
					}
				}
			} else {
				previous = current
				current, _ = strconv.Atoi(level)
				if ascending {
					if !withinLimit(previous, current) {
						safe = false
						break
					}
				} else {
					if !withinLimit(current, previous) {
						safe = false
						break
					}
				}
			}
		}

		if safe {
			part1++
		}

	}

	fmt.Println("Part 1:", part1)

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
					if !withinLimit(current, previous) {
						safe = false
						current = previous
					}
				} else {
					if !withinLimit(previous, current) {
						current = previous
					}
				}
			} else {
				previous = current
				current, _ = strconv.Atoi(level)
				if ascending {
					if !withinLimit(previous, current) {
						if !safe {
							broken = true
							break
						} else {
							safe = false
							current = previous
						}
					}
				} else {
					if !withinLimit(current, previous) {
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
		}

		if !broken {
			part2++
		}

	}

	fmt.Println("Part 2:", part2)
}

func withinLimit(min int, max int) bool {
	if max-min > 3 {
		return false
	}
	if max-min < 1 {
		return false
	}
	return true
}
