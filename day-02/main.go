package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

var delimiters = []rune {
	':',
	',',
	' ',
	';',
}

func split(r rune) bool {
	return slices.Index(delimiters, r) >= 0
}

func main() {
	bytes,err := os.ReadFile("input.txt")
	check(err)

	input := string(bytes)
	lines := strings.Split(input, "\n")

	start := time.Now()
	task1(lines)

	duration1:= time.Since(start)
	start = time.Now()
	
	task2(lines)
	duration2	:= time.Since(start)

	fmt.Println("Task 1:", duration1)
	fmt.Println("Task 2:", duration2)
}
