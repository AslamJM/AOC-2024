package day6

import (
	"aoc-2024/utils"
	"fmt"
	"slices"
	"strings"
)

const fileName = "input"

func parseInput() [][]string {
	s := utils.ReadFile(fileName)
	lines := utils.GetLines(s)

	var out [][]string

	for _, l := range lines {
		pos := strings.Split(l, "")
		out = append(out, pos)
	}

	return out
}

func findStart() (int, int) {
	s := parseInput()

	for i, l := range s {
		for j, c := range l {
			if c == "^" {
				return j, i
			}
		}
	}

	return 0, 0
}

func isExit(dir string, h, w, x, y int) bool {
	if dir == "^" || dir == "v" {
		if h == y {
			return true
		} else {
			return false
		}
	}

	if dir == ">" || dir == "<" {
		if w == x {
			return true
		} else {
			return false
		}
	}

	return false
}

func move(input *[][]string, x, y int) (int, int) {
	inp := *input
	dir := inp[y][x]

	switch dir {
	case "^":
		if inp[y-1][x] == "#" {
			inp[y][x] = ">"
			return x, y
		} else {
			inp[y-1][x] = "^"
			return x, y - 1
		}
	case ">":
		if inp[y][x+1] == "#" {
			inp[y][x] = "v"
			return x, y
		} else {
			inp[y][x+1] = ">"
			return x + 1, y
		}
	case "v":
		if inp[y+1][x] == "#" {
			inp[y][x] = "<"
			return x, y
		} else {
			inp[y+1][x] = "v"
			return x, y + 1
		}
	case "<":
		if inp[y][x-1] == "#" {
			inp[y][x] = "^"
			return x, y
		} else {
			inp[y][x-1] = "<"
			return x - 1, y
		}
	default:
		return x, y

	}
}

func canLoop(input *[][]string, loopPositions *[][2]int, x, y int) [2]int {
	inp := *input
	dir := inp[y][x]

	switch dir {
	case "^":
		if inp[y-1][x] == "." && !slices.Contains(*loopPositions, [2]int{x, y - 1}) {
			return [2]int{x, y - 1}
		}
	case ">":
		if inp[y][x+1] == "." && !slices.Contains(*loopPositions, [2]int{x + 1, y}) {
			return [2]int{x + 1, y}
		}
	case "v":
		if inp[y+1][x] == "." && !slices.Contains(*loopPositions, [2]int{x, y + 1}) {
			return [2]int{x, y + 1}
		}
	case "<":
		if inp[y][x-1] == "." && !slices.Contains(*loopPositions, [2]int{x - 1, y}) {
			return [2]int{x - 1, y}
		}
	}

	return [2]int{}
}

func Part1() {
	input := parseInput()
	x, y := findStart()
	h, w := len(input)-1, len(input[0])-1

	var visited [][2]int = [][2]int{
		{x, y},
	}

	for !isExit(input[y][x], h, w, x, y) {
		x, y = move(&input, x, y)

		if !slices.Contains(visited, [2]int{x, y}) {
			visited = append(visited, [2]int{x, y})
		}
	}

	fmt.Println(len(visited))

}

func Part2() {
	input := parseInput()
	x, y := findStart()
	h, w := len(input)-1, len(input[0])-1

	var visited [][2]int = [][2]int{
		{x, y},
	}

	var loopPositions [][2]int = [][2]int{}

	for !isExit(input[y][x], h, w, x, y) {
		x, y = move(&input, x, y)

		if !slices.Contains(visited, [2]int{x, y}) {
			visited = append(visited, [2]int{x, y})
		} else {
			if pos := canLoop(&input, &loopPositions, x, y); pos != [2]int{} && !slices.Contains(loopPositions, pos) {
				visited = [][2]int{}
				loopPositions = append(loopPositions, pos)
				x, y = findStart()
			}
		}
	}

	fmt.Println(len((loopPositions)))
}
