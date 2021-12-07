package day6

import (
	"aoc-2021/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

var population []int

var popByAge [9]int

type ageGroup struct {
	age int
	num int
}

func Solution() {
	input := utils.ReadInput("internal/day6/input6")

	inputs := strings.Split(input[0], ",")

	part1(inputs)
	part2(inputs, 256)
}

func part2(inputs []string, days int) {
	for _, el := range inputs {
		n, _ := strconv.ParseInt(el, 10, 32)
		popByAge[n]++
	}


	var newPop = [9]int{}

	for i := 0; i < days; i++ {
		if i < 3 {
			fmt.Println(popByAge)
		}
		for j := 0; j < 9; j++ {
			if j == 8 {
				newPop[j] = popByAge[0]
				continue
			}

			newPop[j] = popByAge[j+1]

			if j == 6 {
				newPop[j] += popByAge[0]
			}
		}

		popByAge = newPop
	}

	total := 0

	for _, el := range popByAge {
		total += el
	}

	fmt.Println(total)
}

func part1(inputs []string) {
	population = make([]int, len(inputs))

	for i, el := range inputs {
		n, _ := strconv.ParseInt(el, 10, 32)

		population[i] = int(n)
	}

	reproduce(80)

	fmt.Println(len(population))
}

func reproduce(days uint) {
	// 1.a Cycle population, subtract 1 OR reset to 6 if zero
	// 1.b If zero -> increase append to pop with 8
	for i := uint(0); i < days; i++ {
		for j, el := range population {
			if el == 0 {
				population[j] = 6
				population = append(population, 8)
				continue
			}

			population[j] -= 1
		}
	}
}