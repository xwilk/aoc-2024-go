package main

import (
	"time"
)

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

func day6Solution() (int, int, []time.Duration) {
	input := ReadInputInto2DRunes("day6.input")
	part1 := 0
	part2 := 0

	startPos := findStart(input)
	direction := North
	pos := startPos
	visited := make(map[Position]bool, 0)
	for PosInBounds(pos, input) {
		visited[pos] = true

		posD := PositionInDirection(pos, direction)
		for PosInBounds(posD, input) && input[posD.Y][posD.X] == '#' {
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
		for PosInBounds(pos, candidate) {
			path[PathPoint{pos, direction}] = true

			posD := PositionInDirection(pos, direction)
			for PosInBounds(posD, candidate) && candidate[posD.Y][posD.X] == '#' {
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

	return part1, part2, []time.Duration{}
}
