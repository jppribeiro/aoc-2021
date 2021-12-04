package day3

import (
	"aoc-2021/internal/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solution() {
	lines := utils.ReadInput("internal/day3/input3")

	gamma := 0
	epsilon := 0

	calc := make([]float32, len(strings.Split(lines[0], "")))

	for _, l := range lines {
		t := strings.Split(l, "")

		for i, s := range t {
			n, _ := strconv.ParseFloat(s, 32)
			calc[i] = calc[i] + float32(n)
		}
	}

	for i, c := range calc {
		e := len(calc) - 1 - i

		fmt.Println(c / float32(len(lines)))

		if c/float32(len(lines)) >= 0.500000000 {
			gamma += int(math.Exp2(float64(e)))
		} else {
			epsilon += int(math.Exp2(float64(e)))
		}
	}

	fmt.Printf("Part 1: gamma * epsilon = %d\n", gamma*epsilon)
}
