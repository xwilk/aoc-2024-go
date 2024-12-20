package main

import (
	// "fmt"
	"reflect"
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

func All[T any](s []T, f func(elem T) bool) bool {
	for _, elem := range s {
		if !f(elem) {
			return false
		}
	}

	return true
}

func Any[T any](s []T, f func(elem T) bool) bool {
	for _, elem := range s {
		if f(elem) {
			return true
		}
	}

	return false
}

func Filter[T any](s []T, f func(T) bool) []T {
	vsf := make([]T, 0)
	for _, v := range s {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func moveOneL(grid [][]rune, positions []Position, direction Direction) []Position {
	// fmt.Println("moveOneL", positions, direction)

	targets := []Position{}
	for _, p := range positions {
		targets = append(targets, PositionInDirection(p, direction))
	}

	if Any(targets, func(elem Position) bool { return grid[elem.Y][elem.X] == '#' }) {
		return positions
	}

	if All(targets, func(elem Position) bool { return grid[elem.Y][elem.X] == '.' }) {
		for i, t := range targets {
			grid[t.Y][t.X] = grid[positions[i].Y][positions[i].X]
			grid[positions[i].Y][positions[i].X] = '.'
		}
		return targets
	}

	newTargets := Filter(targets, func(elem Position) bool { return grid[elem.Y][elem.X] != '.' })
	if direction == North || direction == South {
		for i, t := range targets {
			if grid[t.Y][t.X] == '[' && grid[positions[i].Y][positions[i].X] != '[' {
				newTargets = append(newTargets, PositionInDirection(t, East))
			} else if grid[t.Y][t.X] == ']' && grid[positions[i].Y][positions[i].X] != ']' {
				newTargets = append(newTargets, PositionInDirection(t, West))
			}
		}
	}

	newPositions := moveOneL(grid, newTargets, direction)
	if !reflect.DeepEqual(newPositions, newTargets) {
		for i, p := range positions {
			grid[targets[i].Y][targets[i].X] = grid[p.Y][p.X]
			grid[p.Y][p.X] = '.'
		}
		return targets
	}

	return positions
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
	// PrintMap(grid)
	// fmt.Println(instructions)
	pos := findRobot(grid)
	for _, c := range instructions {
		switch c {
		case '^':
			pos = moveOne(grid, pos, North)
		case '>':
			pos = moveOne(grid, pos, East)
		case 'v':
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
		// fmt.Printf("Move %c\n", c)
		switch c {
		case '^':
			pos = moveOneL(grid2, []Position{pos}, North)[0]
		case '>':
			pos = moveOneL(grid2, []Position{pos}, East)[0]
		case 'v':
			pos = moveOneL(grid2, []Position{pos}, South)[0]
		case '<':
			pos = moveOneL(grid2, []Position{pos}, West)[0]
		}

		// PrintMap(grid2)
		// time.Sleep(1 * time.Second)
	}
	PrintMap(grid2)
	part2 := 0

	for y := 0; y < len(grid2); y++ {
		for x := 0; x < len(grid2[0]); x++ {
			if grid2[y][x] == '[' {
				part2 += 100*y + x
			}
		}
	}

	// 1534133 too low
	return part1, part2, []time.Duration{}
}
