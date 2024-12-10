package main

import (
	"time"
)

func findTrailheads(input [][]int) []Position {
	var trailheads []Position
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input); col++ {
			if input[row][col] == 0 {
				trailheads = append(trailheads, Position{X: col, Y: row})
			}
		}
	}

	return trailheads
}

func countReachablePaths(trailhead Position, grid [][]int) int {
	candidates := []Position{trailhead}
	visited := make(map[Position]bool)
	stops := make(map[Position]bool)
	for len(candidates) > 0 {
		current := candidates[0]
		candidates = candidates[1:]

		if ok := visited[current]; ok {
			continue
		}

		visited[current] = true

		if grid[current.Y][current.X] == 9 {
			stops[current] = true
			continue
		}

		for _, direction := range CardinalDirections {
			neighbor := PositionInDirection(current, direction)
			if PosInBounds(neighbor, grid) && grid[neighbor.Y][neighbor.X]-grid[current.Y][current.X] == 1 {
				candidates = append(candidates, neighbor)
			}
		}
	}

	return len(stops)
}

func countTrails(trailhead Position, grid [][]int) int {
	candidates := []Position{trailhead}
	visited := make(map[Position]bool)
	trails := 0
	for len(candidates) > 0 {
		current := candidates[0]
		candidates = candidates[1:]

		visited[current] = true

		if grid[current.Y][current.X] == 9 {
			trails += 1
			continue
		}

		for _, direction := range CardinalDirections {
			neighbor := PositionInDirection(current, direction)
			if PosInBounds(neighbor, grid) && grid[neighbor.Y][neighbor.X]-grid[current.Y][current.X] == 1 {
				candidates = append(candidates, neighbor)
			}
		}
	}

	return trails
}

func day10Solution() (int, int, []time.Duration) {
	input := ReadInputInto2DInts("day10.input")

	part1 := 0
	part2 := 0
	traiheads := findTrailheads(input)
	for _, t := range traiheads {
		part1 += countReachablePaths(t, input)
		part2 += countTrails(t, input)
	}

	return part1, part2, []time.Duration{}
}
