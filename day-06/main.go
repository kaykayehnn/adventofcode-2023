package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func countWays(time, distance int) int {
	total := 0

	for i := 1; i < time; i++ {
		if i*(time-i) > distance {
			total++
		}
	}

	return total
}

func check(e error) {
	if e != nil {
		panic(e)
	}
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
