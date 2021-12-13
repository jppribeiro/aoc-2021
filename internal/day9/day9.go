package day9

import (
	"aoc-2021/internal/utils"
	"fmt"
	"sort"
	"strconv"
)

var heightMap [][]int

func Solution() {
	lines := utils.ReadInput("internal/day9/input9")

	heightMap = parseHeightMap(lines)

	res := part1()

	fmt.Println(res)

	fmt.Println(part2())
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
	lows := lows()

	var basins []int

	for _, l := range lows {
		n := basin(l)

		basins = append(basins, n)
	}

	sort.Ints(basins)

	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

func lows() [][2]int {
	top, right, bottom, left := 0, 0, 0, 0

	var lows [][2]int

	for i:=1;i<len(heightMap)-1;i++ {
		for j:=1;j<len(heightMap[i])-1;j++ {
			top, right, bottom, left = i-1, j+1, i+1, j-1

			n := heightMap[i][j]
			if n < heightMap[i][left] && n < heightMap[i][right] && n < heightMap[top][j] && n < heightMap[bottom][j] {
				lows = append(lows, [2]int{i, j})
			}
		}
	}

	return lows
}

func basin(low [2]int) int {
	basin := make([][]int, len(heightMap))
	for i := range basin {
		basin[i] = make([]int, len(heightMap[i]))
	}

	basin[low[0]][low[1]] = 1

	safe := 0
	next := 1

	for safe != next {
		safe = next
		next = 0

		for i := 1; i < len(basin) - 1; i++ {
			for j := 1; j < len(basin[i]) - 1; j++ {
				if basin[i][j] == 1 {
					// left
					if heightMap[i][j-1] < 9 && basin[i][j-1] < 1 {
						basin[i][j-1]++
					}

					// right
					if heightMap[i][j+1] < 9 && basin[i][j+1] < 1 {
						basin[i][j+1]++
					}

					// top
					if heightMap[i-1][j] < 9 && basin[i-1][j] < 1 {
						basin[i-1][j]++
					}

					// bottom
					if heightMap[i+1][j] < 9 && basin[i+1][j] < 1 {
						basin[i+1][j]++
					}
				}
			}
		}

		for _, l := range basin {
			for _, n := range l {
				if n > 0 {
					next += n
				}
			}
		}
	}

	return next
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