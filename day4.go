package main

import ()

var runeArray = [...]rune{'M', 'A', 'S'}

func searchInDirection(input [][]rune, x, y int, direction Direction) bool {
	for _, r := range runeArray {
		x, y = PositionInDirection(x, y, direction)
		if !IsInBounds(x, y, input) {
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
	xnw, ynw := PositionInDirection(x, y, NorthWest)
	xse, yse := PositionInDirection(x, y, SouthEast)
	xne, yne := PositionInDirection(x, y, NorthEast)
	xsw, ysw := PositionInDirection(x, y, SouthWest)

	if !IsInBounds(xnw, ynw, input) || !IsInBounds(xse, yse, input) || !IsInBounds(xne, yne, input) || !IsInBounds(xsw, ysw, input) {
		return 0
	}

	arm1 := (input[ynw][xnw] == 'M' && input[yse][xse] == 'S') || (input[ynw][xnw] == 'S' && input[yse][xse] == 'M')
	arm2 := (input[yne][xne] == 'M' && input[ysw][xsw] == 'S') || (input[yne][xne] == 'S' && input[ysw][xsw] == 'M')
	if arm1 && arm2 {
		return 1
	}

	return 0
}

func day4Solution() (int, int) {
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

	return part1, part2
}
