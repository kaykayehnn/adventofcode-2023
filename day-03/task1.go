package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func task1(lines []string) {
	regex := regexp.MustCompile(`\d+`)

	sum := 0
	for i, line := range lines {
		matches := regex.FindAllStringIndex(line, -1)
		for _, match := range matches {
			for j := match[0]; j < match[1]; j++ {
				if isNextToSymbol(lines, i, j) {
					num, _ := strconv.Atoi(line[match[0]:match[1]])
					sum += num
					break
				}
			}
		}
	}

	fmt.Println("Task 1: xxxxxx")
}
