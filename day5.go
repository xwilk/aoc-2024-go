package main

import (
	"strconv"
	"strings"
	"time"
)

func parseInput(filename string) (map[int][]int, [][]int) {
	input := strings.Split(ReadInput(filename), "\n\n")
	rulesStr := strings.Split(input[0], "\n")
	rulesStr = rulesStr[:len(rulesStr)-1]
	rules := make(map[int][]int, len(rulesStr))
	for _, r := range rulesStr {
		p := strings.Split(r, "|")
		n1, _ := strconv.Atoi(p[0])
		n2, _ := strconv.Atoi(p[1])
		rules[n1] = append(rules[n1], n2)
	}

	updatesStr := strings.Split(input[1], "\n")
	updates := make([][]int, len(updatesStr))
	for i, u := range updatesStr {
		chars := strings.Split(u, ",")
		update := make([]int, len(chars))
		for j, c := range chars {
			update[j], _ = strconv.Atoi(c)
		}
		updates[i] = update
	}

	return rules, updates
}

func day5Solution() (int, int, []time.Duration) {
	rules, updates := parseInput("day5.input")
	part1 := 0
	part2 := 0

	var incorrectUpdates [][]int
	for _, u := range updates {
		visited := make(map[int]bool)
		visited[u[0]] = true
		correct := true
		for _, n := range u[1:] {
			for _, r := range rules[n] {
				if ok := visited[r]; ok {
					correct = false
					break
				}
			}

			if !correct {
				break
			}

			visited[n] = true
		}

		if correct {
			part1 += u[len(u)/2]
		} else {
			incorrectUpdates = append(incorrectUpdates, u)
		}
	}

	for _, update := range incorrectUpdates {
		i := 0
		sorted := false
		for !sorted {
			n := update[i]

			error := false
			for j := i + 1; j < len(update); j++ {
				for _, r := range rules[update[j]] {
					if n == r {
						var newU []int
						newU = append(newU, update[:i]...)
						newU = append(newU, update[j])
						newU = append(newU, update[i:j]...)
						newU = append(newU, update[j+1:]...)
						update = newU

						error = true
						break
					}
				}
			}

			if !error {
				i++
				if i == len(update) {
					sorted = true
				}
			}
		}

		part2 += update[len(update)/2]
	}

	return part1, part2, []time.Duration{}
}
