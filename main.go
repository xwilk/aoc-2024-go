package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func TimeIt(fn func() (int, int, []time.Duration)) func() (int, int, []time.Duration) {
	return func() (int, int, []time.Duration) {
		start := time.Now()
		part1, part2, measurements := fn()
		elapsed := time.Since(start)
		fmt.Printf("Full execution time: %s\n", elapsed)
		return part1, part2, measurements
	}
}

var functionsMap = map[int]func() (int, int, []time.Duration){
	1:  day1Solution,
	2:  day2Solution,
	3:  day3Solution,
	4:  day4Solution,
	5:  day5Solution,
	6:  day6Solution,
	7:  day7Solution,
	8:  day8Solution,
	9:  day9Solution,
	10: day10Solution,
	11: day11Solution,
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
		part1Solution, part2Solution, measurements := TimeIt(f)()
		fmt.Println("part1: ", part1Solution)
		fmt.Println("part2: ", part2Solution)
		fmt.Println("measurements: ", measurements)
	}
}
