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

// GCD using the Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM for two numbers
func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0 // LCM is undefined if either number is 0
	}
	return AbsInt(a*b) / gcd(a, b)
}

// LCM for multiple numbers
func lcmMultiple(nums []int) int {
	if len(nums) == 0 {
		return 0 // No numbers, no LCM
	}

	result := nums[0]
	for _, num := range nums[1:] {
		result = lcm(result, num)
	}
	return result
}
