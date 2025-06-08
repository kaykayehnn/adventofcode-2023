package main

import (
	"fmt"
	"strconv"
)

var GEAR = '*'

func task2(lines []string) {
	sum := 0
	for i, line := range lines {
		for j, char := range line {
			if char == GEAR {
				numIndexes := make([][]int, 0)

				foundNumI := -1
				for _, coords := range adjacent {
					newI := i + coords[0]
					newJ := j + coords[1]
					if isOutOfBounds(lines, newI, newJ) {
						continue
					}

					char := lines[newI][newJ]
					foundNum := isNum(char)
					if !foundNum {
						foundNumI = -1
					}
					if foundNum {
						if foundNumI != newI || coords[0] == 0 {
							numIndexes = append(numIndexes, []int{newI, newJ})
							foundNumI = newI
						}
					}
				}

				if len(numIndexes) == 2 {
					product := 1
					// find the nums
					for _, numIndex := range numIndexes {
						numRow := numIndex[0]
						numStart := numIndex[1]
						numEnd := numIndex[1]

						for !isOutOfBounds(lines, numRow, numStart-1) && isNum(lines[numRow][numStart-1]) {
							numStart--
						}
						for !isOutOfBounds(lines, numRow, numEnd+1) && isNum(lines[numRow][numEnd+1]) {
							numEnd++
						}

						num, _ := strconv.Atoi(lines[numRow][numStart : numEnd+1])
						product *= num
					}

					sum += product
				}
			}
		}
	}

	fmt.Println("Task 2: xxxxxx")
}
