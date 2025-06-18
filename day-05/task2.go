package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func mapValue(allMaps [][][]int, value int) int {
	mappedValue := value
	for _, myMap := range allMaps {
		for _, mapping := range myMap {
			destination := mapping[0]
			start := mapping[1]
			length := mapping[2]
			if mappedValue >= start && mappedValue < start+length {
				offset := mappedValue - start
				mappedValue = destination + offset
				break
			}
		}
	}

	return mappedValue
}

func task2(lines []string) {

	allMaps := make([][][]int, 0)
	for _, el := range lines[1:] {
		if el == "" {
			continue
		}
		if strings.Contains(el, "map") {
			allMaps = append(allMaps, make([][]int, 0))
		} else {
			nums := strings.Split(el, " ")
			parsedNums := make([]int, 0)
			for _, el := range nums {
				parsed, _ := strconv.Atoi(el)
				parsedNums = append(parsedNums, parsed)
			}
			allMaps[len(allMaps)-1] = append(allMaps[len(allMaps)-1], parsedNums)
		}
	}

	seeds := strings.Split(strings.ReplaceAll(lines[0], "seeds: ", ""), " ")
	parsedSeeds := make([]int, 0)
	for _, el := range seeds {
		parsed, _ := strconv.Atoi(el)
		parsedSeeds = append(parsedSeeds, parsed)
	}

	lowestLocation := math.MaxInt64
	for i := 0; i < len(parsedSeeds); i += 2 {
		value := parsedSeeds[i]
		for j := 0; j < parsedSeeds[i+1]; j++ {
			lowestLocation = min(lowestLocation, mapValue(allMaps, value+j))
		}
	}
	fmt.Println(lowestLocation)
}
