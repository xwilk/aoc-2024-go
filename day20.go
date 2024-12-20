package main

import (
	"container/heap"
	"time"
)

func findCheatyPath(start, end Position, grid [][]rune, distances map[Position]int, cheatDist int) []int {
	candidates := []Position{start}
	visited := make(map[Position]bool)
	cheatyDistances := []int{}

	for len(candidates) > 0 {
		current := candidates[0]
		candidates = candidates[1:]

		if ok := visited[current]; ok {
			continue
		}

		visited[current] = true

		for _, distNeighbor := range PositionsWithinDistance(2, cheatDist, current) {
			if PosInBounds(distNeighbor, grid) && distances[distNeighbor] > distances[current] {
				distanceSaved := distances[distNeighbor] - distances[current] - Distance(current, distNeighbor)
				if distanceSaved > 0 {
					cheatyDistances = append(cheatyDistances, distanceSaved)
				}
			}
		}

		for _, direction := range CardinalDirections {
			neighbor := PositionInDirection(current, direction)
			if neighbor == end {
				return cheatyDistances
			}
			if PosInBounds(neighbor, grid) && grid[neighbor.Y][neighbor.X] != '#' {
				candidates = append(candidates, neighbor)
			}
		}
	}

	return cheatyDistances
}

func cpuRace(start, end Position, grid [][]rune) map[Position]int {
	candidates := &MinHeap{&Node{start, 0, nil}}
	heap.Init(candidates)

	distances := make(map[Position]int)
	distances[start] = 0

	for candidates.Len() > 0 {
		current := heap.Pop(candidates).(*Node)

		for _, direction := range CardinalDirections {
			neighborPos := PositionInDirection(current.pos, direction)
			if !PosInBounds(neighborPos, grid) || grid[neighborPos.Y][neighborPos.X] == '#' {
				continue
			}

			newDist := current.distance + 1
			if _, ok := distances[neighborPos]; !ok || newDist < distances[neighborPos] {
				distances[neighborPos] = newDist
				neighborNode := Node{pos: neighborPos, distance: newDist, prev: current}
				heap.Push(candidates, &neighborNode)
			}
		}
	}

	return distances
}

func FindStartAndEnd(grid [][]rune) (Position, Position) {
	var start, end Position
	for y, row := range grid {
		for x, char := range row {
			if char == 'S' {
				start = Position{x, y}
			} else if char == 'E' {
				end = Position{x, y}
			}
		}
	}
	return start, end
}

func day20Solution() (int, int, []time.Duration) {
	grid := ReadInputInto2DRunes("day20.input")
	PrintMap(grid)
	start, end := FindStartAndEnd(grid)
	distances := cpuRace(start, end, grid)
	part1 := 0
	cheatyPathsDistances := findCheatyPath(start, end, grid, distances, 2)
	for _, dist := range cheatyPathsDistances {
		if dist >= 100 {
			part1++
		}
	}

	part2 := 0
	cheatyPathsDistances2 := findCheatyPath(start, end, grid, distances, 20)
	for _, dist := range cheatyPathsDistances2 {
		if dist >= 100 {
			part2++
		}
	}

	return part1, part2, []time.Duration{}
}
