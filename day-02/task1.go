package main

import (
	"fmt"
	"strconv"
	"strings"
)

func task1(lines []string) {
  ballsMap := map[string]int {
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	sum := 0
	for _, line := range lines {
		tokens := strings.FieldsFunc(line, split)
		gameNumber,_ := strconv.Atoi(tokens[1])
		// Skip the `Game N` string
		tokens = tokens[2:]
		possible := true

		for i:=0; i < len(tokens); i+=2{
			color := tokens[i+1]
			balls, _ := strconv.Atoi(tokens[i])
			cap := ballsMap[color]
			
			possible = balls <= cap
			if !possible {
				break
			}
		}

		if possible {
			sum += gameNumber
		}
	}

	fmt.Println(sum)
}
