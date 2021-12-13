package day9

import (
	"aoc-2021/internal/utils"
	"fmt"
	"strconv"
)

var heightMap [][]int

func Solution() {
	lines := utils.ReadInput("internal/day9/input9")

	heightMap = parseHeightMap(lines)

	res := part1()

	fmt.Println(res)
}

func part1() int {
	top, right, bottom, left := 0, 0, 0, 0

	danger := 0

	for i:=1;i<len(heightMap)-1;i++ {
		for j:=1;j<len(heightMap[i])-1;j++ {
			top, right, bottom, left = i-1, j+1, i+1, j-1

			n := heightMap[i][j]
			if n < heightMap[i][left] && n < heightMap[i][right] && n < heightMap[top][j] && n < heightMap[bottom][j] {
				danger += n + 1
			}
		}
	}

	return danger
}

func part2() int {

}

func basin(low [2]int) int {

}

func parseHeightMap(lines []string) [][]int {
	h := make([][]int,len(lines) + 2, len(lines) + 2)

	for i := 0; i < len(lines) + 2; i++ {
		h[i] = make([]int, len(lines[0])+2, len(lines[0])+2)

		for j := 0; j < len(lines[0]) + 2; j++ {
			if j == 0 ||
				j == len(lines[0]) + 1 ||
				i == 0 ||
				i == len(lines) + 1 {
				h[i][j] = 9
				continue
			}

			h[i][j], _ = strconv.
				Atoi(lines[i-1][j-1:j])
		}
	}

	return h
}