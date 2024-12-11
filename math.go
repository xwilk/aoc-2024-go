package main

import "math"

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
