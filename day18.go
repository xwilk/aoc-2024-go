package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	pos      Position
	distance int
	prev     *Node
}

type MinHeap []*Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].distance < h[j].distance } // Min-heap: smallest value at root
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*Node))
}

// Pop removes and returns the smallest element
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func createGrid(width, height, bytes int, lines []string) [][]rune {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	bytes = min(len(lines), bytes)
	for _, line := range lines[:bytes] {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		grid[y][x] = '#'
	}

	return grid
}

func findShortesPath(start, end Position, grid [][]rune) int {
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

	if _, exists := distances[end]; exists {
		return distances[end]
	}

	return 0
}

func day18Solution() (int, int, []time.Duration) {
	input := ReadInput("day18.input")
	const width, height, bytes = 71, 71, 1024
	start, end := Position{0, 0}, Position{70, 70}

	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	grid := createGrid(width, height, bytes, lines)
	PrintMap(grid)

	part1 := findShortesPath(start, end, grid)
	part2 := 0
	x := 0
	y := 0
	for i := 1; i < len(lines); i++ {
		coords := strings.Split(lines[bytes+i], ",")
		x, _ = strconv.Atoi(coords[0])
		y, _ = strconv.Atoi(coords[1])
		grid[y][x] = '#'

		part2 = findShortesPath(start, end, grid)
		if part2 == 0 {
			break
		}
	}

	fmt.Printf("part2: %v,%v", x, y)
	return part1, part2, nil
}
