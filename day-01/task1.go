package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func task1(lines []string) {
	firstRegex := regexp.MustCompile(`.*?(\d)`)
	lastRegex := regexp.MustCompile(`.*(\d)`)
	sum := 0
	for _, element := range lines {
		firstMatch := firstRegex.FindStringSubmatch(element)
		lastMatch := lastRegex.FindStringSubmatch(element)

		firstNumber, err := strconv.Atoi(firstMatch[1])
		check(err)
		secondNumber, err := strconv.Atoi(lastMatch[1])
		check(err)

		number := firstNumber*10 + secondNumber
		sum += number
	}

	fmt.Println(sum)
}
