package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type nToBlinks struct {
	num        int
	blinks     int
	iterations int
}

var cache = map[nToBlinks]int{}

func blink(n, i, t int) int {
	ntb := nToBlinks{n, i, t}
	if cached, exists := cache[ntb]; exists {
		return cached
	}
	if i == t {
		return 1
	}

	if n == 0 {
		res := blink(1, i+1, t)
		cache[ntb] = res
		return res
	} else if digits := NumOfDigits(n); digits%2 == 0 {
		multiplier := int(math.Pow10(digits / 2))
		res := blink(n/multiplier, i+1, t) + blink(n%multiplier, i+1, t)
		cache[ntb] = res
		return res
	} else {
		res := blink(n*2024, i+1, t)
		cache[ntb] = res
		return res
	}
}

func blinks(nums []int, t int) int {
	fmt.Printf("nums: %v, iterations: %v\n", nums, t)
	total := 0
	for _, n := range nums {
		total += blink(n, 0, t)
		fmt.Printf("n: %v, total: %v\n", n, total)
	}
	return total
}

func day11Solution() (int, int, []time.Duration) {
	input := ReadInput("day11.input")
	input = strings.TrimSuffix(input, "\n")
	numbersStrings := strings.Split(input, " ")
	numbers := make([]int, len(numbersStrings))
	for i, ns := range numbersStrings {
		numbers[i], _ = strconv.Atoi(ns)
	}

	part1 := blinks(numbers, 25)
	part2 := blinks(numbers, 75)
	return part1, part2, []time.Duration{}
}
