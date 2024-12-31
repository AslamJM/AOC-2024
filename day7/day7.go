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

func permutations(l int, permMap *map[int][]rune, curr []rune, ops []rune) {

	mp := *permMap

	if mp[l] != nil {
		return
	}
	if len(curr) == l {
		mp[l] = append([]rune(nil), curr...)
		curr = []rune{}
		return
	}

	for _, op := range ops {
		curr = append(curr, op)
		permutations(l, permMap, curr, ops)
		curr = curr[:len(curr)-1]
	}
}

func Part1() {

	input := parseInput()

	ops := []rune{'*', '+'}
	permMap := make(map[int][]rune)

	permutations(len(input), &permMap, []rune{}, ops)

	for length, sequence := range permMap {
		fmt.Printf("Length %d: %s\n", length, string(sequence))
	}
}
