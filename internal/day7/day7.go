package day7

import (
	"aoc-2021/internal/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type world struct {
	min int
	max int
	crabs []int
}

func Solution() {
	lines := utils.ReadInput("internal/day7/input7")

	pos := strings.Split(lines[0], ",")

	w := resolveWorld(pos)

	fmt.Println(part1(w))
	fmt.Println(part2(w))
}

func part2(w world) int {
	fuel := math.MaxInt64

	for i := w.min; i < w.max; i++ {
		tFuel := 0

		for _, crb := range w.crabs {
			n := int(math.Abs(float64(crb - i)))
			tFuel += n * (n + 1) / 2
		}

		if tFuel < fuel {
			fuel = tFuel
		}
	}

	return fuel
}

func part1(w world) int {
	fuel := math.MaxInt64

	for i := w.min; i < w.max; i++ {
		tFuel := 0

		for _, crb := range w.crabs {
			tFuel += int(math.Abs(float64(crb - i)))
		}

		if tFuel < fuel {
			fuel = tFuel
		}
	}

	return fuel
}

func resolveWorld(pos []string) world {
	var conv []int
	min := math.MaxInt32
	max := 0

	for _, pos := range pos {
		n, _ := strconv.ParseInt(pos, 10, 32)
		if int(n) > max {
			max = int(n)
		}

		if int(n) < min {
			min = int(n)
		}

		conv = append(conv, int(n))
	}

	return world{min, max, conv}
}

