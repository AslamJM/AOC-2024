package day5

import (
	"aoc-2024/utils"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const fileName = "te"

func parseInput() ([]string, []string) {
	s := utils.ReadFile(fileName)

	chuncks := utils.GetLines(s)

	var cfg []string
	var breakPoint int

	for i, c := range chuncks {
		l := strings.Trim(c, " ")
		if l == "" {
			breakPoint = i + 1
			break
		}
		cfg = append(cfg, l)
	}

	upds := chuncks[breakPoint:]

	return cfg, upds
}

type cfgNode struct {
	before []int
	after  []int
}

func makeConfigMap(cfg []string) map[int]cfgNode {
	var cfgMap = make(map[int]cfgNode)

	for _, c := range cfg {
		digits := strings.Split(c, "|")

		numL, err := strconv.Atoi(strings.Trim(digits[0], " "))
		if err != nil {
			panic("map: not a number")
		}
		numR, err := strconv.Atoi(strings.Trim(digits[1], " "))
		if err != nil {
			panic("map: not a number")
		}

		r := cfgMap[numR]
		l := cfgMap[numL]

		r.before = append(r.before, numL)
		l.after = append(l.after, numR)

		cfgMap[numL] = cfgNode{
			before: l.before,
			after:  l.after,
		}
		cfgMap[numR] = cfgNode{
			before: r.before,
			after:  r.after,
		}
	}

	return cfgMap
}

func toIntArr(s string) []int {
	nums := strings.Split(s, ",")

	var out = []int{}

	for _, d := range nums {
		n, err := strconv.Atoi(strings.Trim(d, " "))
		if err != nil {
			panic("not a number")
		}
		out = append(out, n)
	}

	return out
}

func isInOrder(nums []int, cfgMap map[int]cfgNode) bool {
	for i, n := range nums {
		// check before
		bf := nums[:i]

		if len(bf) > 0 {
			cfgAfter := cfgMap[n].after
			for _, k := range bf {
				if slices.Contains(cfgAfter, k) {
					return false
				}
			}
		}

		// check after
		af := nums[i+1:]
		if len(af) > 0 {
			cfgBefore := cfgMap[n].before

			for _, k := range af {
				if slices.Contains(cfgBefore, k) {
					return false
				}
			}
		}
	}

	return true
}

func Part1() {
	cfg, upds := parseInput()
	cfgMap := makeConfigMap(cfg)
	sum := 0

	for _, l := range upds {
		nums := toIntArr(l)

		if isInOrder(nums, cfgMap) {
			sum += nums[len(nums)/2]
		}

	}

	fmt.Println(sum)

}

func Part2() {
	cfg, upds := parseInput()
	cfgMap := makeConfigMap(cfg)
	sum := 0

	for _, l := range upds {
		nums := toIntArr(l)

		if !isInOrder(nums, cfgMap) {
			sort.Slice(nums, func(i, j int) bool {
				if slices.Contains(cfgMap[nums[j]].before, nums[i]) {
					return true
				}
				if slices.Contains(cfgMap[nums[j]].after, nums[i]) {
					return false
				}
				if slices.Contains(cfgMap[nums[i]].before, nums[j]) {
					return false
				}
				if slices.Contains(cfgMap[nums[i]].after, nums[j]) {
					return true
				}

				return true
			})

			sum += nums[len(nums)/2]
		}

	}

	fmt.Println(sum)
}
