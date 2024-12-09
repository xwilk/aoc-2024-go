package main

import (
	"fmt"
	"strings"
	"time"
)

func expandInput(input string) []rune {
	var res []rune
	for col, c := range input {
		n := int(c) - 48
		for i := 0; i < n; i++ {
			if col%2 == 0 {
				res = append(res, rune('0'+col/2))
			} else {
				res = append(res, '.')
			}
		}
	}
	return res
}

func shiftWithFragmentation(disk []rune) []rune {
	i := 0
	j := len(disk) - 1
	for j > i {
		if disk[i] != '.' {
			i++
		} else if disk[j] == '.' {
			j--
		} else {
			tmp := disk[j]
			disk[j] = disk[i]
			disk[i] = tmp
		}
	}

	return disk
}

func shiftNoFragmentation(disk []rune) []rune {
	countDots := func(i int) int {
		count := 0
		for i < len(disk) && disk[i] == '.' {
			count++
			i++
		}

		return count
	}

	countFileSize := func(i int, c rune) int {
		count := 0
		for i >= 0 && disk[i] == c {
			count++
			i--
		}

		return count
	}

	i := 0
	j := len(disk) - 1
	for j > 0 {
		if disk[i] != '.' {
			i++
		} else if disk[j] == '.' {
			j--
		} else {
			dotCount := 0
			if i < j {
				dotCount = countDots(i)
			}

			fileSize := countFileSize(j, disk[j])

			if i < j && fileSize <= dotCount {
				for k := i; k < i+fileSize; k++ {
					disk[k] = disk[j]
				}

				for k := j; k > j-fileSize; k-- {
					disk[k] = '.'
				}

				j -= fileSize
				i = 0
			} else {
				i += dotCount
				if i < len(disk) && (disk[i] == disk[j] || i >= j) {
					i = 0
					j -= fileSize
				}
			}
		}
	}

	return disk
}

func countChecksum(disk []rune) int {
	checksum := 0
	for i, c := range disk {
		if c == '.' {
			continue
		}
		n := int(c) - 48
		checksum += i * n
	}

	return checksum
}

func printDisk(disk []rune) {
	for _, c := range disk {
		fmt.Printf("%c", c)
	}
}

func day9Solution() (int, int, []time.Duration) {
	input := ReadInput("day9.input")
	input = strings.TrimSuffix(input, "\n")

	disk1 := expandInput(input)
	disk2 := CopySlice(disk1)

	disk1 = shiftWithFragmentation(disk1)
	part1 := countChecksum(disk1)
	disk2 = shiftNoFragmentation(disk2)
	part2 := countChecksum(disk2)

	// 6448989155953 - part1
	// 6476642796832 - part2
	return part1, part2, []time.Duration{}
}
