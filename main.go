package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var functionsMap = map[int]func() (int, int){
	1: day1Solution,
	2: day2Solution,
	3: day3Solution,
}

func parseDayInput() int {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("Error - no arguments provided")
		os.Exit(1)
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Error - invalid argument provided")
		os.Exit(1)
	}

	return day
}

func main() {
	day := parseDayInput()
	f, exist := functionsMap[day]
	if exist {
		part1Solution, part2Solution := f()
		fmt.Println("part1: ", part1Solution)
		fmt.Println("part2: ", part2Solution)
	}
}
