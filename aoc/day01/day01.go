package day01

import (
	"aoc/utils"
	"slices"
	"strconv"
	"strings"
)

var list1 []int
var list2 []int
var count2 map[int]int

func init() {
	lines := utils.ReadLines("res/day01.in")

	count2 = make(map[int]int)

	for _, line := range lines {
		locations := strings.Fields(line)
		num1, _ := strconv.Atoi(locations[0])
		num2, _ := strconv.Atoi(locations[1])
		count2[num2]++
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	slices.Sort(list1)
	slices.Sort(list2)
}

func Part1() int {
	part1 := 0

	for i := range list1 {
		part1 += utils.Abs(list2[i] - list1[i])
	}

	return part1
}

func Part2() int {
	part2 := 0

	for i := range list1 {
		part2 += list1[i] * count2[list1[i]]
	}

	return part2
}
