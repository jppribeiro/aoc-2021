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

	part1(lines)

	part2(lines)

}

func part1(lines []string) {
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

		calc[i] = c / float32(len(lines))

		if calc[i] >= 0.500000000 {
			gamma += int(math.Exp2(float64(e)))
		} else {
			epsilon += int(math.Exp2(float64(e)))
		}
	}

	fmt.Printf("Part 1: gamma * epsilon = %d\n", gamma*epsilon)
}

func part2(lines []string) {

	oxy := strings.Split(filter(lines, 0, "oxy"), "")
	co2 := strings.Split(filter(lines, 0, "co2"), "")

	o := 0
	c := 0

	for i := range oxy {
		n, _ := strconv.ParseInt(oxy[i], 10, 8)
		m, _ := strconv.ParseInt(co2[i], 10, 8)

		e := len(oxy) - 1 - i
		o += int(n) * int(math.Exp2(float64(e)))
		c += int(m) * int(math.Exp2(float64(e)))
	}

	fmt.Printf("Part 2: oxygen * co2 = %d\n", o*c)
}

func filter(lines []string, i int, criteria string) string {
	valid := []string{}

	iavg := calcAvg(lines)

	if len(lines) == 1 || i > 11 {
		return lines[0]
	}

	for _, l := range lines {
		switch criteria {
		case "oxy":
			if iavg[i] == 0.5 && l[i:i+1] == "1" {
				valid = append(valid, l)
			}

			if (iavg[i] > 0.5 && l[i:i+1] == "1") || (iavg[i] < 0.5 && l[i:i+1] == "0") {
				valid = append(valid, l)
			}
		case "co2":
			if iavg[i] == 0.5 && l[i:i+1] == "0" {
				valid = append(valid, l)
			}

			if (iavg[i] > 0.5 && l[i:i+1] == "0") || (iavg[i] < 0.5 && l[i:i+1] == "1") {
				valid = append(valid, l)
			}
		}
	}

	return filter(valid, i+1, criteria)
}

func calcAvg(lines []string) []float32 {
	calc := make([]float32, len(strings.Split(lines[0], "")))

	for _, l := range lines {
		t := strings.Split(l, "")

		for i, s := range t {
			n, err := strconv.ParseFloat(s, 32)

			if err != nil {
				panic("error")
			}
			calc[i] = calc[i] + float32(n)
		}
	}

	for i, c := range calc {
		calc[i] = c / float32(len(lines))
	}

	return calc
}
