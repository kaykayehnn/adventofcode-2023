package main

import (
	"fmt"
	"strconv"
	"strings"
)

func task1(lines []string) {
	time := make([]int, 0)
	distance := make([]int, 0)
	isFirst := true
	for _, el := range lines {
		tokens := strings.Split(el, " ")
		for _, el := range tokens[1:] {
			if el == "" {
				continue
			}

			parsed, _ := strconv.Atoi(el)
			if isFirst {
				time = append(time, parsed)
			} else {
				distance = append(distance, parsed)
			}
		}

		isFirst = false
	}

	totalWays := 1
	for i := range len(time) {
		totalWays *= countWays(time[i], distance[i])
	}

	fmt.Println(totalWays)
}
