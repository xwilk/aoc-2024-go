package main

import (
	"sort"
	"strconv"
	"strings"
)

func day1Solution() (int, int) {
	input := strings.Split(ReadInput("day1.input"), "\n")
	input = input[:len(input)-1]
	left := make([]int, len(input))
	right := make([]int, len(input))
	for i, pair := range input {
		splited := strings.Split(pair, "   ")
		left[i], _ = strconv.Atoi(splited[0])
		right[i], _ = strconv.Atoi(splited[1])
	}

	sort.Ints(left)
	sort.Ints(right)

	part1 := 0
	for i := 0; i < len(left); i++ {
		part1 += AbsInt(left[i] - right[i])
	}

	counter := MakeCounter(right)
	part2 := 0
	for _, n := range left {
		part2 += n * counter[n]
	}

	return part1, part2
}
