package main

import (
	"fmt"
	"strconv"
	"strings"
)

func task2(lines []string) {
	time := strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", "")
	distance := strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", "")
	parsedTime, _ := strconv.Atoi(time)
	parsedDistance, _ := strconv.Atoi(distance)

	totalWays := countWays(parsedTime, parsedDistance)

	fmt.Println(totalWays)
}
