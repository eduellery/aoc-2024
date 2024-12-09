package day09

import (
	"aoc/utils"
)

var nums []int

func init() {
	listOfNumbers := utils.ReadLines("res/day09.in")[0]

	for _, char := range listOfNumbers {
		num := int(char - '0')
		nums = append(nums, num)
	}
}

func Part1() int {
	index := 0
	counter := 0
	j := len(nums) - 1
	available := nums[j]
	for i, num := range nums {
		if j <= i {
			for available > 0 {
				counter += index * j / 2
				index++
				available--
			}
			break
		}
		if i%2 == 0 {
			for num > 0 {
				counter += index * i / 2
				num--
				index++
			}
		} else {
			for num > 0 {
				if available <= 0 {
					j -= 2
					available = nums[j]
				}
				for available > 0 {
					counter += index * j / 2
					index++
					available--
					num--
					if num <= 0 {
						break
					}
				}
			}
		}
	}
	return counter
}
