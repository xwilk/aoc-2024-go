package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day3Solution() (int, int) {
	input := ReadInput("day3.input")
	part1 := 0
	part2 := 0

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	instructions := re.FindAllStringSubmatch(input, -1)
	barrier := true
	for _, instruction := range instructions {
		fmt.Println("found ", instruction[0])

		if strings.HasPrefix(instruction[0], "mul") {
			n1, _ := strconv.Atoi(instruction[1])
			n2, _ := strconv.Atoi(instruction[2])
			part1 += n1 * n2
			if barrier {
				part2 += n1 * n2
			}
		} else if instruction[0] == "do()" {
			barrier = true
		} else if instruction[0] == "don't()" {
			barrier = false
		}
	}

	return part1, part2
}
