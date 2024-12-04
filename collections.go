package main

import ()

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 |
		~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

func MakeCounter[T Ordered](slice []T) map[T]int {
	counter := make(map[T]int)
	for _, n := range slice {
		if _, ok := counter[n]; ok {
			counter[n]++
		} else {
			counter[n] = 1
		}
	}

	return counter
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
