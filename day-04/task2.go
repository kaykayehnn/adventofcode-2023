package main

import (
	"fmt"
	"regexp"
	"slices"
)

func task2(lines []string) {
	parseRegex := regexp.MustCompile(`Card\s+\d+:\s([\d\s]+)\| ([\d\s]+)`)
	numRegex := regexp.MustCompile(`\d+`)

	winningCards := make([]int, 0)
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

		winningCards = append(winningCards, count)
	}

	// Process the cards
	totalCards := 0
	for i := len(winningCards) - 1; i >= 0; i-- {
		initialWonCards := winningCards[i]
		currentIterationCards := 1
		for j := range initialWonCards {
			currentIterationCards += winningCards[i+j+1]
		}

		winningCards[i] = currentIterationCards
		totalCards += currentIterationCards
	}

	fmt.Println(totalCards)
}
