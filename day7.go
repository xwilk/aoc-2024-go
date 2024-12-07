package main

import (
	"math"
	"strconv"
	"strings"
	"time"
)

type Operation int

const (
	Plus Operation = iota
	Multiply
	Concatonate
)

func find(input []string, operations []Operation) int {
	finalResult := 0

	for _, line := range input {
		data := strings.Split(line, ":")
		value, _ := strconv.Atoi(data[0])
		var numbers []int
		var numbersStrings []string
		for _, n := range strings.Split(data[1], " ") {
			num, err := strconv.Atoi(n)
			if err == nil {
				numbers = append(numbers, num)
				numbersStrings = append(numbersStrings, n)
			}
		}
		results := []int{numbers[0]}
		for i := 1; i < len(numbers); i++ {
			newResults := []int{}
			for operation := range operations {
				for _, result := range results {
					if operation == 0 {
						newResults = append(newResults, result+numbers[i])
					} else if operation == 1 {
						newResults = append(newResults, result*numbers[i])
					} else if operation == 2 {
						digits := NumOfDigits(numbers[i])
						multiplier := int(math.Pow10(digits))
						newResults = append(newResults, multiplier*result+numbers[i])
					}
				}
			}
			results = newResults
		}

		for _, result := range results {
			if result == value {
				finalResult += value
				break
			}
		}
	}
	return finalResult
}

func day7Solution() (int, int, []time.Duration) {
	measurements := []time.Duration{}
	input := strings.Split(ReadInput("day7.input"), "\n")
	input = input[:len(input)-1]

	start := time.Now()
	part1 := find(input, []Operation{Plus, Multiply})
	measurements = append(measurements, time.Since(start))
	start = time.Now()
	part2 := find(input, []Operation{Plus, Multiply, Concatonate})
	measurements = append(measurements, time.Since(start))

	return part1, part2, measurements
}
