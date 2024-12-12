package main

import (
	"fmt"
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

	fmt.Println("p", perimeter)
	return perimeter
}

func findTopLeft(area map[Position]bool) Position {
	topLeft := Position{1000000, 1000000}
	for pos := range area {
		if pos.Y < topLeft.Y {
			topLeft = pos
		} else if pos.Y == topLeft.Y && pos.X < topLeft.X {
			topLeft = pos
		}
	}

	return topLeft
}

type posAndDir struct {
	p Position
	d Direction
}

func countSides(area map[Position]bool) int {
	if len(area) == 1 {
		fmt.Println(4)
		return 4
	}

	topLeft := findTopLeft(area)
	direction := East
	sides := 1
	pos := topLeft
	first := true

	for {
		// time.Sleep(1 * time.Second)

		leftTurnDirection := PreviousCardinalDirection(direction)
		if !first {
			if pos.X == topLeft.X && pos.Y == topLeft.Y && leftTurnDirection == North {
				break
			}
		}
		first = false
		leftPos := PositionInDirection(pos, leftTurnDirection)
		_, exists := area[leftPos]
		if exists {
			sides += 1
			fmt.Println("Turning left:", direction, "at:", pos, "sides:", sides)
			direction = leftTurnDirection
			pos = leftPos
			continue
		}

		if pos.X == topLeft.X && pos.Y == topLeft.Y && direction == North {
			break
		}
		forwardPos := PositionInDirection(pos, direction)
		_, exists = area[forwardPos]
		if exists {
			fmt.Println("Moving forward:", direction, "at:", pos, "sides:", sides)
			pos = forwardPos
			if pos.X == topLeft.X && pos.Y == topLeft.Y && direction == North {
				break
			}
			continue
		}

		rightTurnDirection := NextCardinalDirection(direction)
		if pos.X == topLeft.X && pos.Y == topLeft.Y && rightTurnDirection == North {
			sides += 1
			break
		}
		rightPos := PositionInDirection(pos, rightTurnDirection)
		_, exists = area[rightPos]
		if exists {
			sides += 1
			fmt.Println("Turning right:", direction, "at:", pos, "sides:", sides)
			direction = rightTurnDirection
			pos = rightPos
			if pos.X == topLeft.X && pos.Y == topLeft.Y && direction == North {
				break
			}
			continue
		}

		uTurnDirection := NextCardinalDirection(rightTurnDirection)
		uTurnPos := PositionInDirection(pos, uTurnDirection)
		_, exists = area[uTurnPos]
		if exists {
			sides += 2
			fmt.Println("Turning around:", direction, "at:", pos, "sides:", sides)
			direction = uTurnDirection
			pos = uTurnPos
			if pos.X == topLeft.X && pos.Y == topLeft.Y && direction == North {
				break
			}
		}
	}

	fmt.Println(sides)
	return sides
}

func day12Solution() (int, int, []time.Duration) {
	// input := ReadInputInto2DRunes("day12.input")
	input := ReadInputInto2DRunes("day12.testinput")

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

	// 850842 too low
	return part1, part2, []time.Duration{}
}
