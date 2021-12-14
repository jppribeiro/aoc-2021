package day10

import (
	"aoc-2021/internal/utils"
	"fmt"
	"regexp"
)

var pairs = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var points = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

func Solution() {
	lines := utils.ReadInput("internal/day10/input10")

	fmt.Println(part1(lines))
}

func part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		sum += resolveLine(line)
	}

	return sum
}

func resolveLine(line string) int {
	for len(line) != 0 {
		for i := 0; i < len(line)-1; i++ {
			if isStart(line[i:i+1]) && isEnd(line[i+1:i+2]) {
				if isCorrupt(line[i : i+2]) {
					return points[line[i+1:i+2]]
				}

				line = line[0:i] + line[i+2:]
				break
			}
		}

		if isStart(line[len(line)-1:]) {
			line = line[0 : len(line)-1]
		}
	}

	return 0
}

func isCorrupt(s string) bool {
	return pairs[s[0:1]] != s[1:2]
}

func isStart(s string) bool {
	rgx := regexp.MustCompile("\\(|\\<|\\{|\\[")

	return rgx.MatchString(s)
}

func isEnd(s string) bool {
	rgx := regexp.MustCompile("\\)|\\>|\\}|\\]")

	return rgx.MatchString(s)
}
