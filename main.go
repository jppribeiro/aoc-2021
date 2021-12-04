package main

import (
	"aoc-2021/internal/day1"
	"aoc-2021/internal/day2"
	"fmt"
	"os"
)

var solutions = map[string]func(){
	"day1": day1.Solution,
	"day2": day2.Solution,
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Pass a day as argument: day<num>")
		os.Exit(1)
	}

	arg := os.Args[1]

	solutions[arg]()
}
