package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
)

func task1(lines []string) {
	parseRegex := regexp.MustCompile(`Card\s+\d+:\s([\d\s]+)\| ([\d\s]+)`)
	numRegex := regexp.MustCompile(`\d+`)

	sum := 0
	for _, line := range lines {
		matches := parseRegex.FindStringSubmatch(line)
		winningNums := numRegex.FindAllString(matches[1], -1)
		myNums := numRegex.FindAllString(matches[2], -1)

		count := 0
		for _, num := range myNums {
			if slices.Index(winningNums, num) != -1 {
				count++
			}
		}

		if count > 0 {
			sum += int(math.Pow(2, float64(count-1)))
		}
	}

	fmt.Println(sum)
}
