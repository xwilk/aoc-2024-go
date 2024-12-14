package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type machine struct {
	A     Position
	B     Position
	Prize Position
}

func calcTimesA(v1, v2, target Position) int64 {
	m1, c1, _ := CalculateLineEquation(Position{0, 0}, v1)
	m2, c2, _ := CalculateLineEquation(target.Sub(v2), target)
	x, y, err := FindLinesIntersection(m1, c1, m2, c2)
	if err != nil {
		return 0
	}

	x = RoundToPrecision(x, 2)
	y = RoundToPrecision(y, 2)
	if math.Mod(x, 1) != 0 || math.Mod(y, 1) != 0 {
		return 0
	}

	xi := int64(x)
	yi := int64(y)
	if xi%int64(v1.X) != 0 || yi%int64(v1.Y) != 0 {
		return 0
	}

	return xi / int64(v1.X)
}

func calculatePrizes(machines []machine) int64 {
	var moves int64 = 0
	for _, m := range machines {
		timesA := calcTimesA(m.A, m.B, m.Prize)
		timesB := calcTimesA(m.B, m.A, m.Prize)

		if timesA != 0 && timesB != 0 {
			moves += timesA*3 + timesB
		}
	}

	return moves
}

func day13Solution() (int, int, []time.Duration) {
	input := ReadInput("day13.input")
	input = strings.TrimSuffix(input, "\n")
	machinesBlock := strings.Split(input, "\n\n")
	machines1 := []machine{}
	machines2 := []machine{}
	for _, block := range machinesBlock {
		re := regexp.MustCompile(`X.(\d+), Y.(\d+)`)
		coords := re.FindAllStringSubmatch(block, -1)
		ax, _ := strconv.Atoi(coords[0][1])
		ay, _ := strconv.Atoi(coords[0][2])
		bx, _ := strconv.Atoi(coords[1][1])
		by, _ := strconv.Atoi(coords[1][2])
		px, _ := strconv.Atoi(coords[2][1])
		py, _ := strconv.Atoi(coords[2][2])

		machines1 = append(machines1, machine{Position{ax, ay}, Position{bx, by}, Position{px, py}})
		machines2 = append(machines2, machine{Position{ax, ay}, Position{bx, by}, Position{px + 10000000000000, py + 10000000000000}})
	}

	part1 := calculatePrizes(machines1)
	part2 := calculatePrizes(machines2)

	fmt.Println(part1, part2)
	return int(part1), int(part2), []time.Duration{}
}
