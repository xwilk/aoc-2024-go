package main

import (
	"time"
)

func bfdSearch(start Position, grid [][]rune, assigned map[Position]bool) map[Position]bool {
	candidates := []Position{start}
	visited := make(map[Position]bool)
	area := map[Position]bool{}

	areaType := grid[start.Y][start.X]
	for len(candidates) > 0 {
		current := candidates[0]
		candidates = candidates[1:]

		if ok := visited[current]; ok {
			continue
		}

		visited[current] = true

		if grid[current.Y][current.X] != areaType {
			continue
		}

		area[current] = true
		assigned[current] = true
		for _, direction := range CardinalDirections {
			neighbor := PositionInDirection(current, direction)
			if PosInBounds(neighbor, grid) {
				candidates = append(candidates, neighbor)
			}
		}
	}

	return area
}

func countPerimeter(area map[Position]bool) int {
	perimeter := 0
	for pos := range area {
		p := 4
		for _, direction := range CardinalDirections {
			neighbor := PositionInDirection(pos, direction)
			if _, exists := area[neighbor]; exists {
				p--
			}
		}
		perimeter += p
	}

	return perimeter
}

func findMinY(area map[Position]bool) int {
	minY := 1000000
	for pos := range area {
		if pos.Y < minY {
			minY = pos.Y
		}
	}

	return minY
}

func findMaxY(area map[Position]bool) int {
	maxY := 0
	for pos := range area {
		if pos.Y > maxY {
			maxY = pos.Y
		}
	}

	return maxY
}

func findMinX(area map[Position]bool) int {
	minX := 1000000
	for pos := range area {
		if pos.X < minX {
			minX = pos.X
		}
	}

	return minX
}

func findMaxX(area map[Position]bool) int {
	maxX := 0
	for pos := range area {
		if pos.X > maxX {
			maxX = pos.X
		}
	}

	return maxX
}

func countSides(area map[Position]bool) int {
	minY := findMinY(area)
	maxY := findMaxY(area)
	minX := findMinX(area)
	maxX := findMaxX(area)

	sides := 0
	for row := minY; row <= maxY; row++ {
		for col := minX; col <= maxX; col++ {
			if _, exists := area[Position{X: col, Y: row}]; exists {
				_, aboveExists := area[Position{X: col, Y: row - 1}]
				_, leftExists := area[Position{X: col - 1, Y: row}]
				_, rightExists := area[Position{X: col + 1, Y: row}]
				_, northWestExists := area[Position{X: col - 1, Y: row - 1}]
				_, northEastExists := area[Position{X: col + 1, Y: row - 1}]
				if !leftExists && !aboveExists {
					sides += 1
				} else if !leftExists && northWestExists {
					sides += 1
				}
				if !rightExists && !aboveExists {
					sides += 1
				} else if !rightExists && northEastExists {
					sides += 1
				}
			}
		}
	}

	for col := minX; col <= maxX; col++ {
		for row := minY; row <= maxY; row++ {
			if _, exists := area[Position{X: col, Y: row}]; exists {
				_, aboveExists := area[Position{X: col, Y: row - 1}]
				_, belowExists := area[Position{X: col, Y: row + 1}]
				_, leftExists := area[Position{X: col - 1, Y: row}]
				_, northWestExists := area[Position{X: col - 1, Y: row - 1}]
				_, southWestExists := area[Position{X: col - 1, Y: row + 1}]
				if !aboveExists && !leftExists {
					sides += 1
				} else if !aboveExists && northWestExists {
					sides += 1
				}
				if !belowExists && !leftExists {
					sides += 1
				} else if !belowExists && southWestExists {
					sides += 1
				}
			}
		}
	}

	return sides
}

func day12Solution() (int, int, []time.Duration) {
	input := ReadInputInto2DRunes("day12.input")

	areas := []map[Position]bool{}
	assigned := map[Position]bool{}
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[row]); col++ {
			if _, exists := assigned[Position{X: col, Y: row}]; !exists {
				areas = append(areas, bfdSearch(Position{X: col, Y: row}, input, assigned))
			}
		}
	}

	part1 := 0
	part2 := 0

	for _, area := range areas {
		part1 += len(area) * countPerimeter(area)
		part2 += len(area) * countSides(area)
	}

	return part1, part2, []time.Duration{}
}
