package day11

import (
	"aoc/utils"
	"slices"
	"strconv"
)

var input []string

func init() {
	input = utils.ReadListOfStrings("res/day11.in")
}

func transform(numbers []string) []string {
	i := 0
	for i < len(numbers) {
		if numbers[i] == "0" {
			numbers[i] = "1"
		} else if x := len(numbers[i]); x%2 == 0 {
			firstHalf, _ := strconv.Atoi(numbers[i][:x/2])
			secondHalf, _ := strconv.Atoi(numbers[i][x/2:])
			numbers = slices.Insert(numbers, i, strconv.Itoa(firstHalf))
			i++
			numbers[i] = strconv.Itoa(secondHalf)
		} else {
			num, _ := strconv.Atoi(numbers[i])
			numbers[i] = strconv.Itoa(num * 2024)
		}
		i++
	}
	return numbers
}

func Part1() int {
	var numbers = input
	for i := 0; i < 25; i++ {
		numbers = transform(numbers)
	}
	return len(numbers)
}
