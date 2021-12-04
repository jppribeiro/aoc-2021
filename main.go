package main

import (
	"aoc-2021/internal/day1"
	"aoc-2021/internal/day2"
	"aoc-2021/internal/day3"
	"aoc-2021/internal/day4"
	"fmt"
	"os"
)

var solutions = map[string]func(){
	"day1": day1.Solution,
	"day2": day2.Solution,
	"day3": day3.Solution,
	"day4": day4.Solution,
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Pass a day as argument: day<num>")
		os.Exit(1)
	}

	arg := os.Args[1]

	if solutions[arg] == nil {
		fmt.Printf("Invalid argument %s\n", arg)
		os.Exit(1)
	}

	solutions[arg]()
}
