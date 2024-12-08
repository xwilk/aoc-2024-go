package main

import (
	"time"
)

func day8Solution() (int, int, []time.Duration) {
	input := ReadInputInto2DRunes("day8.input")
	antennas := make(map[rune][]Position)
	for row := range len(input) {
		for col := range len(input[0]) {
			c := input[row][col]
			if c != '.' {
				antennas[c] = append(antennas[c], Position{X: col, Y: row})
			}
		}
	}

	part1Positions := make(map[Position]bool)
	part2Positions := make(map[Position]bool)
	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				v := MakeVector(positions[i], positions[j])
				a1 := positions[i].Sub(v)
				a2 := positions[j].Add(v)
				if PosInBounds(a1, input) {
					part1Positions[a1] = true
				}
				if PosInBounds(a2, input) {
					part1Positions[a2] = true
				}

				a1 = positions[i].Add(v)
				a2 = positions[i].Sub(v)
				for PosInBounds(a1, input) {
					part2Positions[a1] = true
					a1 = a1.Add(v)
				}
				for PosInBounds(a2, input) {
					part2Positions[a2] = true
					a2 = a2.Sub(v)
				}

				b1 := positions[j].Add(v)
				b2 := positions[j].Sub(v)
				for PosInBounds(b1, input) {
					part2Positions[b1] = true
					b1 = b1.Add(v)
				}
				for PosInBounds(b2, input) {
					part2Positions[b2] = true
					b2 = b2.Sub(v)
				}
			}
		}
	}

	part1 := len(part1Positions)
	part2 := len(part2Positions)
	return part1, part2, []time.Duration{}
}
