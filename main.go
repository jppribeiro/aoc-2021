package main

import (
	"aoc-2021/internal/day1"
	"aoc-2021/internal/day2"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Pass a day as argument: day<num>")
		os.Exit(1)
	}

	arg := os.Args[1]

	switch arg {
	case "day1":
		day1.Solution()
	case "day2":
		day2.Solution()
	default:
		fmt.Println("Pass a day as argument: day<num>")
	}
}
