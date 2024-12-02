package main

import (
	"bufio"
	"fmt"
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

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var functionsMap = map[int]interface{}{
	1: day1Solution,
	2: day2Solution,
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Error - no arguments provided")
		os.Exit(1)
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error - invalid argument provided")
		os.Exit(1)
	}

	function, exist := functionsMap[day]
	if exist {
		function.(func())()
	}
}
