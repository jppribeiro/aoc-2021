package day4

import (
	"aoc-2021/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

var boards []*board

var isWin bool = false

var lastWinner *board
var lastWinnerN int

type board struct {
	nums   [][]int
	marks  [][]bool
	colSum []int
	rowSum []int
	hitSum int
	done   bool
}

func Solution() {
	lines := utils.ReadInput("internal/day4/input4")

	res := start(lines)

	fmt.Printf("First winner: %d\n", res[0])
	fmt.Printf("Last winner: %d\n", res[1])
}

func start(lines []string) [2]int {
	rawResult := strings.Split(lines[0], ",")

	parseBoards(lines[2:])

	results := [2]int{}

	for _, r := range rawResult {
		n, _ := strconv.ParseInt(r, 10, 32)

		for _, b := range boards {
			if b.done {
				continue
			}

			if b.isDone(int(n)) {
				if !isWin {
					isWin = true
					results[0] = b.result(int(n))
				}

				b.done = true
				lastWinner = b
				lastWinnerN = int(n)
			}
		}
	}

	results[1] = lastWinner.result(lastWinnerN)
	return results
}

func parseBoards(lines []string) {
	newBoard := false
	b := &board{[][]int{}, [][]bool{}, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 0, false}

	for i, l := range lines {
		row := []int{}
		mark := []bool{}

		l = strings.TrimSpace(l)
		l = strings.ReplaceAll(l, "  ", " 0")
		if l == "" {
			boards = append(boards, b)
			newBoard = true
			continue
		}

		if newBoard {
			b = &board{[][]int{}, [][]bool{}, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 0, false}
		}

		for _, n := range strings.Split(l, " ") {
			d, _ := strconv.ParseInt(n, 10, 32)

			row = append(row, int(d))
			mark = append(mark, false)
		}

		b.nums = append(b.nums, row)
		b.marks = append(b.marks, mark)
		newBoard = false

		if i == len(lines)-1 {
			boards = append(boards, b)
		}
	}
}

func (bo *board) isDone(numCalled int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if (*bo).nums[i][j] == numCalled {
				(*bo).marks[i][j] = true
				bo.colSum[j] = bo.colSum[j] - 1
				bo.rowSum[i] = bo.rowSum[i] - 1
				bo.hitSum = bo.hitSum + numCalled
			}
		}
	}

	for i := 0; i < 5; i++ {
		if bo.colSum[i] == -5 || bo.rowSum[i] == -5 {
			return true
		}
	}

	return false
}

func (b *board) result(numCalled int) int {
	res := 0

	for _, row := range b.nums {
		for _, c := range row {
			res += c
		}
	}

	return (res - b.hitSum) * numCalled
}
