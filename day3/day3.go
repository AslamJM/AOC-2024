package day3

import (
	"aoc-2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseInput() string {
	return utils.ReadFile("test")
}

func mul(s string) int {
	digits := s[4 : len(s)-1]

	nums := strings.Split(digits, ",")
	n1, err := strconv.Atoi(nums[0])
	if err != nil {
		return 0
	}

	n2, err := strconv.Atoi(nums[1])
	if err != nil {
		return 0
	}

	return n1 * n2
}

func totalMult(s string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllString(s, -1)

	total := 0

	for _, match := range matches {
		total += mul(match)
	}

	return total
}

func Part1() {
	input := parseInput()
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := re.FindAllString(input, -1)

	total := 0

	for _, match := range matches {
		total += mul(match)
	}

	fmt.Println(total)
}

func Part2() {
	input := parseInput()

	segments := strings.Split(input, "do")

	total := 0

	for _, seg := range segments {
		if strings.HasPrefix(seg, "n't") {
			continue
		}

		total += totalMult(seg)
	}

	fmt.Println(total)
}
