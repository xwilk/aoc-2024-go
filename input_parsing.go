package main

import (
	"bufio"
	"log"
	"os"
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
