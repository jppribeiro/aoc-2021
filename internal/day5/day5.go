package day5

import (
	"aoc-2021/internal/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type line struct {
	x1    int
	x2    int
	y1    int
	y2    int
	ortho bool
}

func Solution() {
	lines := utils.ReadInput("internal/day5/input5")

	mapOrtho(lines)

	mapAll(lines)
}

func mapOrtho(lines []string) {
	ortho := buildLines(lines, true)

	worldSize := worldSize(ortho)

	world := make([][]int, worldSize[1]+1)

	for i := range world {
		world[i] = make([]int, worldSize[0]+1)
	}

	trace(world, ortho)

	dangers := 0

	for _, row := range world {
		for _, pos := range row {
			if pos >= 2 {
				dangers++
			}
		}
	}

	fmt.Printf("Dangers: %d\n", dangers)
}

func mapAll(lines []string) {
	intLines := buildLines(lines, false)

	worldSize := worldSize(intLines)

	world := make([][]int, worldSize[1]+1)

	for i := range world {
		world[i] = make([]int, worldSize[0]+1)
	}

	trace(world, intLines)

	dangers := 0

	for _, row := range world {
		for _, pos := range row {
			if pos >= 2 {
				dangers++
			}
		}
	}

	fmt.Printf("Dangers: %d\n", dangers)
}

func trace(world [][]int, lines []line) {
	// For each line:
	// 1. Determine min x and min y
	// 2. sub-set the world with x1 x2 y1 y2

	// For each position of the sub set add 1

	for _, l := range lines {
		var minX int
		var maxX int
		var minY int
		var maxY int

		if l.x1 > l.x2 {
			minX = l.x2
			maxX = l.x1
		} else {
			minX = l.x1
			maxX = l.x2
		}

		if l.y1 > l.y2 {
			minY = l.y2
			maxY = l.y1
		} else {
			minY = l.y1
			maxY = l.y2
		}

		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				if l.ortho {
					world[y][x] += 1
					continue
				}

				if (l.x2 - l.x1) == 0 {
					spew.Dump(l)
				}

				a := (l.y2 - l.y1) / (l.x2 - l.x1)
				b := l.y2 - a*l.x2

				y0 := a*x + b

				if y0 == y {
					world[y][x] += 1
				}
			}
		}
	}
}

func buildLines(lines []string, ortho bool) []line {
	filteredLines := []line{}
	for _, strline := range lines {
		pts := strings.Split(strline, " -> ")

		x1, _ := strconv.ParseInt(strings.Split(pts[0], ",")[0], 10, 32)
		x2, _ := strconv.ParseInt(strings.Split(pts[1], ",")[0], 10, 32)
		y1, _ := strconv.ParseInt(strings.Split(pts[0], ",")[1], 10, 32)
		y2, _ := strconv.ParseInt(strings.Split(pts[1], ",")[1], 10, 32)

		if ortho {
			if x1 == x2 || y1 == y2 {
				filteredLines = append(filteredLines, line{x1: int(x1), x2: int(x2), y1: int(y1), y2: int(y2), ortho: true})
			}
		} else {
			if x1 == x2 || y1 == y2 {
				filteredLines = append(filteredLines, line{x1: int(x1), x2: int(x2), y1: int(y1), y2: int(y2), ortho: true})
				continue
			}

			filteredLines = append(filteredLines, line{x1: int(x1), x2: int(x2), y1: int(y1), y2: int(y2), ortho: false})
		}
	}

	return filteredLines
}

func worldSize(lines []line) []int {
	maxX := 0
	maxY := 0

	for _, l := range lines {
		if l.x1 > l.x2 {
			if l.x1 > maxX {
				maxX = l.x1
			}
		} else {
			if l.x2 > maxX {
				maxX = l.x2
			}
		}

		if l.y1 > l.y2 {
			if l.y1 > maxY {
				maxY = l.y1
			}
		} else {
			if l.y2 > maxY {
				maxY = l.y2
			}
		}
	}

	return []int{maxX, maxY}
}
