package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v", filename)
		os.Exit(1)
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
		sb.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read file: %v", err)
		os.Exit(1)
	}

	return sb.String()
}

func ReadInputInto2DRunes(filename string) [][]rune {
	input := ReadInput(filename)
	lines := strings.Split(input, "\n")
	result := make([][]rune, len(lines)-1)
	for i, line := range lines[:len(lines)-1] {
		result[i] = []rune(line)
	}

	return result
}

func ReadInputInto2DInts(filename string) [][]int {
	input := ReadInput(filename)
	lines := strings.Split(input, "\n")
	result := make([][]int, len(lines)-1)
	for i, line := range lines[:len(lines)-1] {
		for j, c := range line {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatalf("Rune is not a number! %c", c)
			}

			result[i][j] = n
		}
	}

	return result
}
