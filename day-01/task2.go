package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func task2(lines []string) {
	nums := []string {"_zero_", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	numsRegex := strings.Join(nums[:], "|")

	firstRegex := regexp.MustCompile(fmt.Sprintf(`.*?(%s|\d)`, numsRegex))
	lastRegex := regexp.MustCompile(fmt.Sprintf(`.*(%s|\d)`, numsRegex))
	sum := 0
	for _, element := range lines {
		firstMatch := firstRegex.FindStringSubmatch(element)
		lastMatch := lastRegex.FindStringSubmatch(element)

		firstNumIndex := slices.Index(nums, firstMatch[1])
		if firstNumIndex == -1 {
			firstNumIndex, _ = strconv.Atoi(firstMatch[1])
		}

		lastNumIndex := slices.Index(nums, lastMatch[1])
		if lastNumIndex == -1 {
			lastNumIndex, _ = strconv.Atoi(lastMatch[1])
		}
		
		number := firstNumIndex * 10 + lastNumIndex
		sum += number
	}

	fmt.Println(sum)
}
