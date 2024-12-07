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

func check(acc int, key int, values []int) bool {
	if len(values) == 0 || acc > key {
		return false
	}
	if len(values) == 1 {
		return key == acc*values[0] || key == acc+values[0]
	}
	return check(acc+values[0], key, values[1:]) || check(acc*values[0], key, values[1:])
}

func Part1() int {
	part1 := 0
	for key, values := range content {
		if check(values[0]+values[1], key, values[2:]) || check(values[0]*values[1], key, values[2:]) {
			part1 += key
		}
	}
	return part1
}
