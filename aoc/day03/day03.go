package day03

import (
	"aoc/utils"
	"regexp"
	"strconv"
	"strings"
)

var content string

func init() {
	content = utils.ReadContent("res/day03.in")
}

func mul(content string) (sum int) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, pair := range re.FindAllStringSubmatch(content, -1) {
		num1, _ := strconv.Atoi(pair[1])
		num2, _ := strconv.Atoi(pair[2])
		sum += num1 * num2
	}
	return sum
}

func Part1() int {
	return mul(content)
}

func Part2() int {
	re := regexp.MustCompile(`(?s)(?:^|do\(\)).*?(?:don't\(\)|$)`)
	return mul(strings.Join(re.FindAllString(string(content), -1), ""))
}
