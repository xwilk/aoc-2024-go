package main

import (
	"strings"
	"time"
)

var designCache = map[string]int{}

func isPossible(design string, patterns []string) int {
	if design == "" {
		return 1
	}

	if res, exists := designCache[design]; exists {
		return res
	}

	arrangements := 0
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			arrangements += isPossible(design[len(pattern):], patterns)
		}

	}
	designCache[design] = arrangements
	return arrangements
}

func day19Solution() (int, int, []time.Duration) {
	input := ReadInput("day19.input")
	input = strings.TrimSuffix(input, "\n")
	data := strings.Split(input, "\n\n")
	patternsStr := data[0]
	designsStr := data[1]
	patterns := []string{}
	for _, pattern := range strings.Split(patternsStr, ", ") {
		patterns = append(patterns, pattern)
	}
	designs := []string{}
	for _, design := range strings.Split(designsStr, "\n") {
		designs = append(designs, design)
	}

	part1 := 0
	part2 := 0
	for _, design := range designs {
		res := isPossible(design, patterns)
		if res > 0 {
			part1++
			part2 += res
		}
	}

	return part1, part2, []time.Duration{}
}
