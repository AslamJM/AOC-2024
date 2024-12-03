package day2

import (
	"aoc-2024/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const fileName = "test"

func parseInput() [][]int {
	s := utils.ReadFile(fileName)
	lines := utils.GetLines(s)

	var sl [][]int

	for _, l := range lines {
		numStrs := strings.Split(l, " ")
		var nums []int

		for _, str := range numStrs {
			n, err := strconv.Atoi(strings.Trim(str, " "))
			if err != nil {
				log.Fatal(err)
			}

			nums = append(nums, n)
		}

		sl = append(sl, nums)
	}

	return sl
}

func testSafety(a []int) bool {

	i := 0
	order := "asc"

	if len(a) == 1 {
		return true
	}

	if a[0] > a[1] {
		order = "desc"
	}

	for i < len(a)-1 {

		if a[i] == a[i+1] {
			return false
		}

		if utils.Abs(a[i]-a[i+1]) > 3 {
			return false
		}

		switch order {
		case "asc":
			if a[i] > a[i+1] {
				return false
			}
		case "desc":
			if a[i] < a[i+1] {
				return false
			}
		}
		i++
	}

	return true
}

func removeElement(a []int, i int) []int {
	var x = []int{}
	x = append(x, a[:i]...)
	x = append(x, a[i+1:]...)
	return x
}

func tolerable(a []int) bool {
	if testSafety(a) {
		return true
	}

	i := 0

	for i < len(a) {
		rm := removeElement(a, i)
		if testSafety(rm) {
			return true
		}
		i++
	}

	return false
}

func Part1() {
	reports := parseInput()

	safe := 0

	for _, nums := range reports {
		if testSafety(nums) {
			safe++
		}
	}

	fmt.Println(safe)

}

func Part2() {
	reports := parseInput()
	safe := 0

	for _, nums := range reports {
		if tolerable(nums) {
			safe++
		}
	}

	fmt.Println(safe)
}
