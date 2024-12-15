package main

import (
	"fmt"
	"strings"
	"time"
)

func moveOne(grid [][]rune, pos Position, direction Direction) Position {
	target := PositionInDirection(pos, direction)
	if grid[target.Y][target.X] == '#' {
		return pos
	}

	if grid[target.Y][target.X] == '.' {
		grid[target.Y][target.X] = grid[pos.Y][pos.X]
		grid[pos.Y][pos.X] = '.'
		return target
	}

	if grid[target.Y][target.X] == 'O' {
		newPos := moveOne(grid, target, direction)
		if newPos != target {
			grid[target.Y][target.X] = grid[pos.Y][pos.X]
			grid[pos.Y][pos.X] = '.'
			return target
		}
	}

	return pos
}

func moveOneL(grid [][]rune, pos Position, direction Direction) Position {
	target := PositionInDirection(pos, direction)
	if grid[target.Y][target.X] == '#' {
		return pos
	}

	if grid[target.Y][target.X] == '.' {
		grid[target.Y][target.X] = grid[pos.Y][pos.X]
		grid[pos.Y][pos.X] = '.'
		return target
	}

	if grid[target.Y][target.X] == '[' || grid[target.Y][target.X] == ']' {
		if direction == East || direction == West {
			newPos := moveOneL(grid, target, direction)
			if newPos != target {
				grid[target.Y][target.X] = grid[pos.Y][pos.X]
				grid[pos.Y][pos.X] = '.'
				return target
			}
		}

		r := pos
		l := pos
		rTarget := target
		lTarget := target
		if direction == North {
			if grid[target.Y][target.X] == '[' {
				r = PositionInDirection(pos, East)
				rTarget = PositionInDirection(target, East)
			} else {
				l = PositionInDirection(pos, West)
				lTarget = PositionInDirection(target, West)
			}

			newPosR := moveOneL(grid, rTarget, direction)
			newPosL := moveOneL(grid, lTarget, direction)

			if newPosR != rTarget && newPosL != lTarget {
				grid[lTarget.Y][lTarget.X] = grid[l.Y][l.X]
				grid[l.Y][l.X] = '.'
				grid[rTarget.Y][rTarget.X] = grid[r.Y][r.X]
				grid[r.Y][r.X] = '.'
				return target
			} else {
				return pos
			}
		}
	}

	return pos
}

func findRobot(grid [][]rune) Position {
	for y, row := range grid {
		for x, cell := range row {
			if cell == '@' {
				return Position{x, y}
			}
		}
	}
	return Position{}
}

func InputToLargerGrid(input string) [][]rune {
	lines := strings.Split(input, "\n")
	result := make([][]rune, len(lines))
	for i, line := range lines {
		result[i] = make([]rune, len(line)*2)
		for j, c := range line {
			if c == '@' {
				result[i][j*2] = c
				result[i][j*2+1] = '.'
			}
			if c == '#' || c == '.' {
				result[i][j*2] = c
				result[i][j*2+1] = c
			}
			if c == 'O' {
				result[i][j*2] = '['
				result[i][j*2+1] = ']'
			}
		}
	}

	return result
}

func day15Solution() (int, int, []time.Duration) {
	input := strings.Split(ReadInput("day15.testinput"), "\n\n")
	// input := strings.Split(ReadInput("day15.input"), "\n\n")
	grid := InputTo2DRunes(input[0])
	instructions := strings.Replace(input[1], "\n", "", -1)
	PrintMap(grid)
	fmt.Println(instructions)
	pos := findRobot(grid)
	for _, c := range instructions {
		switch c {
		case '^':
			pos = moveOne(grid, pos, North)
		case '>':
			pos = moveOne(grid, pos, East)
		case 'v':
			// snapshot := Copy2DSlice(grid)
			pos = moveOne(grid, pos, South)
		case '<':
			pos = moveOne(grid, pos, West)
		}
	}

	part1 := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'O' {
				part1 += y*100 + x
			}
		}
	}

	grid2 := InputToLargerGrid(input[0])
	PrintMap(grid2)
	pos = findRobot(grid2)
	for _, c := range instructions {
		switch c {
		case '^':
			snapshot := Copy2DSlice(grid2)
			newPos := moveOneL(grid2, pos, North)
			if newPos == pos {
				grid2 = snapshot
			} else {
				pos = newPos
			}
		case '>':
			pos = moveOneL(grid2, pos, East)
		case 'v':
			pos = moveOneL(grid2, pos, South)
		case '<':
			pos = moveOneL(grid2, pos, West)
		}

		fmt.Println("Move", c)
		PrintMap(grid2)
		time.Sleep(1 * time.Second)
	}
	PrintMap(grid2)
	part2 := 0

	return part1, part2, []time.Duration{}
}
