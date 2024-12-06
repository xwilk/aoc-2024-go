package main

import (
// "fmt"
)

type Position struct {
	X int
	Y int
}

type PathPoint struct {
	Position
	Direction
}

func findStart(input [][]rune) Position {
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input); col++ {
			if input[row][col] == '^' {
				return Position{
					X: col,
					Y: row,
				}
			}
		}
	}

	return Position{0, 0}
}

func Copy2DSlice(input [][]rune) [][]rune {
	candidate := make([][]rune, len(input))
	for i := range input {
		candidate[i] = make([]rune, len(input[i]))
		copy(candidate[i], input[i])
	}
	return candidate
}
func day6Solution() (int, int) {
	input := ReadInputInto2DRunes("day6.input")
	part1 := 0
	part2 := 0

	startPos := findStart(input)
	direction := North
	pos := startPos
	visited := make(map[Position]bool, 0)
	for IsPosInBounds(pos, input) {
		visited[pos] = true

		posD := PositionInDirection(pos, direction)
		for IsPosInBounds(posD, input) && input[posD.Y][posD.X] == '#' {
			direction = NextCardinalDirection(direction)
			posD = PositionInDirection(pos, direction)
		}

		pos = posD
	}

	part1 = len(visited)
	delete(visited, startPos)

	for obstaclePos := range visited {
		candidate := Copy2DSlice(input)
		candidate[obstaclePos.Y][obstaclePos.X] = '#'
		pos = startPos
		direction = North
		path := make(map[PathPoint]bool, 0)
		for IsPosInBounds(pos, candidate) {
			path[PathPoint{pos, direction}] = true

			posD := PositionInDirection(pos, direction)
			for IsPosInBounds(posD, candidate) && candidate[posD.Y][posD.X] == '#' {
				direction = NextCardinalDirection(direction)
				posD = PositionInDirection(pos, direction)
			}

			if _, ok := path[PathPoint{posD, direction}]; ok {
				part2 += 1
				break
			}

			pos = posD
		}
	}

	return part1, part2
}
