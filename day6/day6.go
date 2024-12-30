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

func canLoop(input *[][]string, x, y int) bool {
	inp := *input
	dir := inp[y][x]

	switch dir {
	case "^":
		if inp[y+1][x] == "." {
			return true
		}
	case ">":
		if inp[y][x+1] == "." {
			return true
		}
	case "v":
		if inp[y-1][x] == "." {
			return true
		}
	case "<":
		if inp[y][x-1] == "." {
			return true
		}
	}

	return false
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
			if canLoop(&input, x, y) && !slices.Contains(loopPositions, [2]int{x, y}) {
				loopPositions = append(loopPositions, [2]int{x, y})
				x, y = findStart()
			}
		}
	}

	fmt.Println(len((loopPositions)))
}
