package day05

import (
	"aoc/utils"
	"slices"
	"strconv"
	"strings"
)

var instructions string
var block string

func init() {
	content := utils.ReadContent("res/day05.in")
	split := strings.Split(strings.TrimSpace(content), "\n\n")
	instructions = split[0]
	block = split[1]
}

func compare(a, b string) int {
	for _, instruction := range strings.Split(instructions, "\n") {
		if numbers := strings.Split(instruction, "|"); numbers[0] == a && numbers[1] == b {
			return -1
		}
	}
	return 0
}

func run(sorted bool) int {
	r := 0
	for _, pages := range strings.Split(block, "\n") {
		if page := strings.Split(pages, ","); slices.IsSortedFunc(page, compare) == sorted {
			slices.SortFunc(page, compare)
			n, _ := strconv.Atoi(page[len(page)/2])
			r += n
		}
	}
	return r
}

func Part1() int {
	return run(true)
}

func Part2() int {
	return run(false)
}
