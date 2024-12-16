package day5

import (
	"aoc-2024/utils"
	"strconv"
	"strings"
)

const fileName = "input"

func parseInput() ([]string, []string) {
	s := utils.ReadFile(fileName)

	chuncks := utils.GetLines(s)

	var cfg []string
	var breakPoint int

	for i, c := range chuncks {
		l := strings.Trim(c, " ")
		if l == "" {
			breakPoint = i
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
			panic("not a number")
		}
		numR, err := strconv.Atoi(strings.Trim(digits[1], " "))
		if err != nil {
			panic("not a number")
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

func Part1() {
	cfg, upds := parseInput()
	cfgMap := makeConfigMap(cfg)
	sum := 0

	for _, l := range upds {
		nums := toIntArr(l)

	}

}
