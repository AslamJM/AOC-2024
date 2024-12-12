package day4

import (
	"aoc-2024/utils"
	"fmt"
	"strings"
)

const filname = "test"

func parseInput() [][]string {
	s := utils.ReadFile(filname)
	lines := utils.GetLines(s)

	var letters [][]string

	for _, l := range lines {
		chars := strings.Split(l, "")
		letters = append(letters, chars)
	}

	return letters
}

func Part1() {
	input := parseInput()

	cols, rows := len(input), len(input[0])
	count := 0

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if input[i][j] == "X" {
				// check right tr br
				if j <= rows-4 {
					if words := input[i][j+1 : j+4]; strings.Join(words, "") == "MAS" {
						count++
					}

					// tr
					if i >= 3 {
						if strings.Join([]string{
							input[i-1][j+1], input[i-2][j+2], input[i-3][j+3],
						}, "") == "MAS" {
							count++
						}

						// br
						if i <= cols-4 {

							if strings.Join([]string{
								input[i+1][j+1], input[i+2][j+2], input[i+3][j+3],
							}, "") == "MAS" {
								count++
							}
						}
					}

				}

				// check left tl bl
				if j >= 3 {
					if words := input[i][j-3 : j]; strings.Join(words, "") == "SAM" {
						count++
					}

					//tl
					if i >= 3 {
						if strings.Join([]string{
							input[i-1][j-1], input[i-2][j-2], input[i-3][j-3],
						}, "") == "MAS" {
							count++
						}

						// bl
						if i <= cols-4 {
							if strings.Join([]string{
								input[i+1][j-1], input[i+2][j-2], input[i+3][j-3],
							}, "") == "MAS" {
								count++
							}
						}
					}
				}

				// check top
				if i >= 3 {
					if strings.Join([]string{
						input[i-1][j], input[i-2][j], input[i-3][j],
					}, "") == "MAS" {
						count++
					}
				}

				// check bottom
				if i <= cols-4 {
					if strings.Join([]string{
						input[i+1][j], input[i+2][j], input[i+3][j],
					}, "") == "MAS" {
						count++
					}
				}

			}
		}
	}

	fmt.Println(count)
}
