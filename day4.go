package main

import "time"

var runeArray = [...]rune{'M', 'A', 'S'}

func searchInDirection(input [][]rune, x, y int, direction Direction) bool {
	for _, r := range runeArray {
		x, y = CoordinatesInDirection(x, y, direction)
		if !CoordsInBounds(x, y, input) {
			return false
		}
		if input[y][x] != r {
			return false
		}
	}

	return true
}

func directionalSearch(input [][]rune, x, y int) int {
	words := 0
	for _, direction := range []Direction{North, NorthEast, East, SouthEast, South, SouthWest, West, NorthWest} {
		if searchInDirection(input, x, y, direction) {
			words++
		}
	}

	return words
}

func xSearch(input [][]rune, x, y int) int {
	xnw, ynw := CoordinatesInDirection(x, y, NorthWest)
	xse, yse := CoordinatesInDirection(x, y, SouthEast)
	xne, yne := CoordinatesInDirection(x, y, NorthEast)
	xsw, ysw := CoordinatesInDirection(x, y, SouthWest)

	if !CoordsInBounds(xnw, ynw, input) ||
		!CoordsInBounds(xse, yse, input) ||
		!CoordsInBounds(xne, yne, input) ||
		!CoordsInBounds(xsw, ysw, input) {
		return 0
	}

	arm1 := (input[ynw][xnw] == 'M' && input[yse][xse] == 'S') || (input[ynw][xnw] == 'S' && input[yse][xse] == 'M')
	arm2 := (input[yne][xne] == 'M' && input[ysw][xsw] == 'S') || (input[yne][xne] == 'S' && input[ysw][xsw] == 'M')
	if arm1 && arm2 {
		return 1
	}

	return 0
}

func day4Solution() (int, int, []time.Duration) {
	input := ReadInputInto2DRunes("day4.input")
	part1 := 0
	part2 := 0
	for y, line := range input {
		for x, c := range line {
			if c == 'X' {
				part1 += directionalSearch(input, x, y)
			}
			if c == 'A' {
				part2 += xSearch(input, x, y)
			}
		}
	}

	return part1, part2, []time.Duration{}
}
