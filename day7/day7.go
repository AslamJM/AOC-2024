package day7

import (
	"aoc-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

const fileName = "input"

type calibration struct {
	result int
	nums   []int
}

func parseInput() []calibration {
	s := utils.ReadFile(fileName)
	lines := utils.GetLines(s)

	out := []calibration{}

	for _, l := range lines {
		first := strings.Split(l, ":")

		total, err := strconv.Atoi(strings.Trim(first[0], " "))
		if err != nil {
			panic("total is not a number")
		}

		numstring := strings.Split(strings.Trim(first[1], " "), " ")
		nums := []int{}

		for _, n := range numstring {
			num, err := strconv.Atoi(strings.Trim(n, " "))
			if err != nil {
				panic("n is not a number")
			}
			nums = append(nums, num)
		}

		out = append(out, calibration{total, nums})
	}

	return out

}

func permutations(l int, ops []rune) [][]rune {

	if l == 0 {
		return [][]rune{{}}
	}

	var result [][]rune

	one_less := permutations(l-1, ops)

	for _, per := range one_less {
		for _, op := range ops {
			result = append(result, append(per, op))
		}
	}

	return result
}

func isCalibrationCorrect(cal calibration, seq []rune) bool {

	nums := cal.nums
	res := cal.result

	sum := nums[0]

	for i, r := range seq {
		if r == '+' {
			sum += nums[i+1]
		}

		if r == '*' {
			sum *= nums[i+1]
		}
	}

	return sum == res
}

func Part1() {

	input := parseInput()

	ops := []rune{'*', '+'}
	permMap := make(map[int][][]rune)

	sum := 0

	for _, cal := range input {
		perms, ok := permMap[len(cal.nums)-1]

		if !ok {
			permMap[len(cal.nums)-1] = permutations(len(cal.nums)-1, ops)
			perms = permMap[len(cal.nums)-1]
		}

		for _, p := range perms {
			if isCalibrationCorrect(cal, p) {
				sum += cal.result
				break
			}
		}

	}

	fmt.Println(sum)

}
