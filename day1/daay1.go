package day1

import (
	"aoc-2024/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

const (
	inputFileName = "day-1-p-1"
)

func Part1() {
	s := utils.ReadFile(inputFileName)
	lines := utils.GetLines(s)

	var l []int
	var r []int

	for _, line := range lines {
		nums := strings.Split(line, "  ")
		l1, err := strconv.Atoi(strings.Trim(nums[0], " "))

		if err != nil {
			log.Fatal(err)
		}

		r1, err := strconv.Atoi(strings.Trim(nums[1], " "))

		if err != nil {
			log.Fatal(err)
		}

		l = append(l, l1)
		r = append(r, r1)
	}

	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	sum := 0

	for i, num := range l {
		dis := num - r[i]

		if dis > 0 {
			sum += dis
		} else {
			sum -= dis
		}
	}

	fmt.Println(sum)

}

func countMap(a []int) map[int]int {
	m := make(map[int]int)

	for _, n := range a {
		m[n]++
	}

	return m
}

func Part2() {
	s := utils.ReadFile(inputFileName)
	lines := utils.GetLines(s)

	var l []int
	var r []int

	for _, line := range lines {
		nums := strings.Split(line, "  ")
		l1, err := strconv.Atoi(strings.Trim(nums[0], " "))

		if err != nil {
			log.Fatal(err)
		}

		r1, err := strconv.Atoi(strings.Trim(nums[1], " "))

		if err != nil {
			log.Fatal(err)
		}

		l = append(l, l1)
		r = append(r, r1)
	}

	rMap := countMap(r)

	sum := 0

	for _, num := range l {
		sum += num * rMap[num]
	}

	fmt.Println(sum)

}
