package main

import (
	"strconv"
	"strings"
	"time"
)

func analyzeReport(report []int) int {
	last := -1
	direction := 0
	i := 0

	for _, lvl := range report {
		if last == -1 {
			last = lvl
			continue
		}
		if lvl == last {
			break
		}
		distance := AbsInt(last - lvl)
		if distance < 1 || distance > 3 {
			break
		}

		if direction == 0 {
			if lvl > last {
				direction = 1
			} else {
				direction = -1
			}
		}

		if direction == 1 && lvl < last {
			break
		}
		if direction == -1 && lvl > last {
			break
		}

		i++
		last = lvl
	}

	return i
}

func retry(lvls []int, badIndex int) int {
	newLvls := RemoveIndex(lvls, badIndex)
	return analyzeReport(newLvls)
}

func day2Solution() (int, int, []time.Duration) {
	input := ReadInput("day2.input")
	part1 := 0
	part2 := 0

	for _, report := range strings.Split(input, "\n") {
		if report == "" {
			continue
		}

		lvlsString := strings.Split(report, " ")
		lvls := make([]int, len(lvlsString))
		for i, lvl := range lvlsString {
			lvls[i], _ = strconv.Atoi(lvl)
		}

		index := analyzeReport(lvls)

		if index == len(lvls)-1 {
			part1++
			part2++
		} else {
			if retry(lvls, index) == len(lvls)-2 {
				part2++
				continue
			}
			if retry(lvls, index+1) == len(lvls)-2 {
				part2++
				continue
			}
			if index > 0 && retry(lvls, index-1) == len(lvls)-2 {
				part2++
				continue
			}
		}
	}

	return part1, part2, []time.Duration{}
}
