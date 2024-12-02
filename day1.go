package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func day1Solution() {
	input := strings.Split(ReadInput("day1.input"), "\n")
	left := make([]int, len(input))
	right := make([]int, len(input))
	for i, pair := range input {
		if pair == "" {
			continue
		}
		splited := strings.Split(pair, "   ")
		// fmt.Println(splited)
		left[i], _ = strconv.Atoi(splited[0])
		right[i], _ = strconv.Atoi(splited[1])
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDist := 0
	for i := 0; i < len(left); i++ {
		totalDist += AbsInt(left[i] - right[i])
	}

	fmt.Println("part1: ", totalDist)

	leftCounter := make(map[int]int)
	for _, l := range left {
		if _, ok := leftCounter[l]; ok {
			leftCounter[l]++
		} else {
			leftCounter[l] = 1
		}
	}

	rightCounter := make(map[int]int)
	for _, r := range right {
		if _, ok := rightCounter[r]; ok {
			rightCounter[r]++
		} else {
			rightCounter[r] = 1
		}
	}

	leftSimilarityScore := 0
	for _, l := range left {
		leftSimilarityScore += l * rightCounter[l]
	}

	rightSimilarityScore := 0
	for _, r := range right {
		rightSimilarityScore += r * leftCounter[r]
	}

	fmt.Println("part2 - leftSimilarityScore: ", leftSimilarityScore)
	fmt.Println("part2 - sum:", leftSimilarityScore+rightSimilarityScore)
}
