package main

import (
	"fmt"
)

type Position struct {
	X int
	Y int
}

func MakeVector(pos, dest Position) Position {
	return Position{X: dest.X - pos.X, Y: dest.Y - pos.Y}
}

func (p *Position) Add(other Position) Position {
	return Position{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p *Position) Sub(other Position) Position {
	return Position{X: p.X - other.X, Y: p.Y - other.Y}
}

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

func DirectionToVector2D(direction Direction) Position {
	switch direction {
	case North:
		return Position{0, -1}
	case NorthEast:
		return Position{1, -1}
	case East:
		return Position{1, 0}
	case SouthEast:
		return Position{1, 1}
	case South:
		return Position{0, 1}
	case SouthWest:
		return Position{-1, 1}
	case West:
		return Position{-1, 0}
	case NorthWest:
		return Position{-1, -1}
	default:
		fmt.Println("Invalid direction")
		return Position{0, 0}
	}
}

var CardinalDirections = []Direction{North, East, South, West}

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

type PathPoint struct {
	Position
	Direction
}

func CoordinatesInDirection(x, y int, direction Direction) (int, int) {
	v := DirectionToVector2D(direction)
	return x + v.X, y + v.Y
}

func PositionInDirection(pos Position, direction Direction) Position {
	v := DirectionToVector2D(direction)
	return pos.Add(v)
}

func CoordsInBounds[T Ordered](x, y int, input [][]T) bool {
	if y > len(input)-1 || y < 0 || x < 0 || x > len(input[0])-1 {
		return false
	}

	return true
}

func PosInBounds[T Ordered](pos Position, input [][]T) bool {
	return CoordsInBounds(pos.X, pos.Y, input)
}

func PrintMap(input [][]rune) {
	for _, row := range input {
		fmt.Println(string(row))
	}
}
