package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadIntInput(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %s -> %s\n", path, err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, num)
	}
	return lines
}

func ReadInput(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %s -> %s\n", path, err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
