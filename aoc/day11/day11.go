package day11

import (
	"aoc/utils"
	"strconv"
)

var input = map[int]int{}

func init() {
	strings := utils.ReadListOfStrings("res/day11.in")
	for _, str := range strings {
		number, _ := strconv.Atoi(str)
		input[number]++
	}
}

func transform(stones map[int]int) map[int]int {
	var next = map[int]int{}
	for key, value := range stones {
		if key == 0 {
			next[1] += value
		} else if x := strconv.Itoa(key); len(x)%2 == 0 {
			firstHalf, _ := strconv.Atoi(x[:len(x)/2])
			secondHalf, _ := strconv.Atoi(x[len(x)/2:])
			next[firstHalf] += value
			next[secondHalf] += value
		} else {
			next[key*2024] += value
		}
	}
	return next
}

func run(blinks int) (result int) {
	var stones = input
	for i := 0; i < blinks; i++ {
		stones = transform(stones)
	}
	for _, value := range stones {
		result += value
	}
	return result
}

func Part1() (result int) {
	return run(25)
}

func Part2() (result int) {
	return run(75)
}
