package utils

import (
	"log"
	"os"
	"strings"
)

const (
	fildir = "inputs/"
)

func ReadFile(fileName string) string {
	input, err := os.ReadFile(fildir + fileName)

	if err != nil {
		log.Fatal(err)
	}

	return string(input)
}

func GetLines(s string) []string {
	return strings.Split(s, "\r\n")
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
