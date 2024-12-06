package main

import (
	"fmt"
)

type Direction int

const (
	North Direction = iota
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
)

func DirectionTo2DVector(direction Direction) (int, int) {
	switch direction {
	case North:
		return 0, -1
	case NorthEast:
		return 1, -1
	case East:
		return 1, 0
	case SouthEast:
		return 1, 1
	case South:
		return 0, 1
	case SouthWest:
		return -1, 1
	case West:
		return -1, 0
	case NorthWest:
		return -1, -1
	default:
		fmt.Println("Invalid direction")
		return 0, 0
	}
}

func NextCardinalDirection(direction Direction) Direction {
	switch direction {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	default:
		fmt.Println("Invalid direction")
		return 0
	}
}

type Position struct {
	X int
	Y int
}

type PathPoint struct {
	Position
	Direction
}

func CoordinatesInDirection(x, y int, direction Direction) (int, int) {
	xD, yD := DirectionTo2DVector(direction)
	return x + xD, y + yD
}

func PositionInDirection(pos Position, direction Direction) Position {
	xD, yD := DirectionTo2DVector(direction)
	return Position{
		X: pos.X + xD,
		Y: pos.Y + yD,
	}
}

func CoordsInBounds(x, y int, input [][]rune) bool {
	if y > len(input)-1 || y < 0 || x < 0 || x > len(input[0])-1 {
		return false
	}

	return true
}

func PosInBounds(pos Position, input [][]rune) bool {
	return CoordsInBounds(pos.X, pos.Y, input)
}
