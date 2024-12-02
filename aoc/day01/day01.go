package day01

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var part1 int
var part2 int

func Day01() {
	lines, err := utils.ReadLines("res/day01.in")
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	var list1, list2 []int

	for _, line := range lines {
		locations := strings.Fields(line)

		num1, _ := strconv.Atoi(locations[0])
		num2, _ := strconv.Atoi(locations[1])

		if len(locations) >= 2 {
			list1 = append(list1, num1)
			list2 = append(list2, num2)
		}
	}

	sort.Ints(list1)
	sort.Ints(list2)

	frequencyMap1 := make(map[int]int)
	frequencyMap2 := make(map[int]int)

	for i := 0; i < len(list1); i++ {
		current := list1[i] - list2[i]
		if current < 0 {
			current = -current
		}
		part1 += current
		frequencyMap1[list1[i]]++
		frequencyMap2[list2[i]]++
	}

	for key, count := range frequencyMap1 {
		part2 += key * count * frequencyMap2[key]
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
