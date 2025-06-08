package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var adjacent = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, +1},
	{0, -1},
	{0, +1},
	{1, -1},
	{1, 0},
	{1, +1},
}

var nums = map[byte]int{
	'0': 1,
	'1': 1,
	'2': 1,
	'3': 1,
	'4': 1,
	'5': 1,
	'6': 1,
	'7': 1,
	'8': 1,
	'9': 1,
}

var EMPTY_SPACE byte = '.'

func isOutOfBounds(lines []string, i, j int) bool {
	return i < 0 || j < 0 || i >= len(lines) || j >= len(lines[0])
}

func isNum(char byte) bool {
	_, isNum := nums[char]
	return isNum
}

func isNextToSymbol(lines []string, i, j int) bool {
	for _, coord := range adjacent {
		newI := i + coord[0]
		newJ := j + coord[1]

		if isOutOfBounds(lines, newI, newJ) {
			continue
		}

		char := lines[newI][newJ]
		if !isNum(char) && char != EMPTY_SPACE {
			return true
		}
	}

	return false
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	check(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")

	start := time.Now()
	task1(lines)

	duration1 := time.Since(start)
	start = time.Now()

	task2(lines)
	duration2 := time.Since(start)

	fmt.Println("Task 1:", duration1)
	fmt.Println("Task 2:", duration2)
}
