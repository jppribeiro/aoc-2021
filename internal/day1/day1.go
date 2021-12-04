package day1

import (
	"aoc-2021/internal/utils"
	"fmt"
)

func Solution() {
	lines := utils.ReadIntInput("internal/day1/day1-input")

	increases := 0

	for i := 1; i < len(lines); i++ {
		if lines[i] > lines[i-1] {
			increases++
		}
	}

	fmt.Printf("There are %d unit increases\n", increases)

	blockIncreases := 0

	for i := 4; i <= len(lines); i++ {
		if sumSlice(lines[i-3:i]) > sumSlice(lines[i-4:i-1]) {
			blockIncreases++
		}
	}

	fmt.Printf("There are %d block increases", blockIncreases)
}

func sumSlice(slice []int) int {
	sum := 0

	for _, n := range slice {
		sum += n
	}

	return sum
}