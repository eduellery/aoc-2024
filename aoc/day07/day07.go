package day07

import (
	"aoc/utils"
	"strconv"
	"strings"
)

var content = map[int][]int{}

func init() {
	lines := utils.ReadLines("res/day07.in")
	for _, line := range lines {
		equation := strings.Split(line, ":")
		sum, _ := strconv.Atoi(equation[0])
		var list []int
		for _, value := range strings.Split(equation[1], " ") {
			i, _ := strconv.Atoi(value)
			list = append(list, i)
		}
		content[sum] = list
	}
}

func check(acc int, key int, values []int, concat bool) bool {
	if len(values) == 0 || acc > key {
		return false
	}
	if len(values) == 1 {
		if concat && key == conv(acc, values[0]) {
			return true
		}
		return key == acc*values[0] || key == acc+values[0]
	}

	if concat && check(conv(acc, values[0]), key, values[1:], concat) {
		return true
	} else {
		return check(acc+values[0], key, values[1:], concat) || check(acc*values[0], key, values[1:], concat)
	}
}

func conv(a, b int) int {
	str := strconv.Itoa(a) + strconv.Itoa(b)
	i, _ := strconv.Atoi(str)
	return i
}

func Part1() int {
	part1 := 0
	for key, values := range content {
		if check(values[0]+values[1], key, values[2:], false) || check(values[0]*values[1], key, values[2:], false) {
			part1 += key
		}
	}
	return part1
}

func Part2() int {
	part2 := 0
	for key, values := range content {
		if check(values[0]+values[1], key, values[2:], true) || check(values[0]*values[1], key, values[2:], true) {
			part2 += key
		}
	}
	return part2
}
