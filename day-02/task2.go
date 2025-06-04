package main

import (
	"fmt"
	"strconv"
	"strings"
)

func task2(lines []string) {
	sum := 0
	for _, line := range lines {
		ballsMap := map[string]int {}
		tokens := strings.FieldsFunc(line, split)
		// Skip the `Game N` string
		tokens = tokens[2:]

		for i := 0; i < len(tokens); i+=2 {
			color := tokens[i+1]
			balls, _ := strconv.Atoi(tokens[i])
			cap, ok := ballsMap[color]
			if !ok {
				ballsMap[color] = balls
			} else {
				ballsMap[color] = max(cap, balls)
			}
		}

		product := 1
		for _, v := range ballsMap {
			product *= v
		}
		sum += product
	}

	fmt.Println(sum)
}
