package main

import (
	"errors"
	"math"
)

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func NumOfDigits(n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		n = -n
	}

	return int(math.Log10(float64(n))) + 1
}

func FindLinesIntersection(m1, c1, m2, c2 float64) (float64, float64, error) {
	if m1 == m2 {
		return 0, 0, errors.New("No intersection")
	}

	x := (c2 - c1) / (m1 - m2)
	y := m1*x + c1
	return math.Round(x), y, nil
}

func CalculateLineEquation(a, b Position) (float64, float64, error) {
	if a.X == b.X {
		return 0, 0, errors.New("No slope")
	}

	m := (float64(a.Y) - float64(b.Y)) / (float64(a.X) - float64(b.X))
	c := float64(a.Y) - m*float64(a.X)
	return m, c, nil
}

func RoundToPrecision(value float64, precision int) float64 {
	multiplier := math.Pow10(precision)
	return math.Round(value*multiplier) / multiplier
}
