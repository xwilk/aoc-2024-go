package main

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

	digits := 0
	if n < 0 {
		n = -n
	}
	for n > 0 {
		digits++
		n /= 10
	}

	return digits
}
