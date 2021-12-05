package day4

import (
	"reflect"
	"testing"
)

var mockDoneBoard = board{
	nums: [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	},
	colSum: []int{-2, -1, -1, -1, -1},
	rowSum: []int{-5, -1, 0, 0},
	hitSum: 21,
}

func Test_board_result(t *testing.T) {
	type args struct {
		numCalled int
	}
	tests := []struct {
		name  string
		board *board
		args  args
		want  int
	}{
		{"ok", &mockDoneBoard, args{5}, 945},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.board
			if got := b.result(tt.args.numCalled); got != tt.want {
				t.Errorf("board.result() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseBoards(t *testing.T) {
	mockLines := []string{
		" 1  2  3  4  5",
		" 6  7  8  9 10",
		"11 12 13 14 15",
		"16 17 18 19 20",
		"21 22 23 24 25",
		"",
		" 6  7  8  9 10",
		"21 22 23 24 25",
		"26 27 28 29 30",
		"31 32 33 34 35",
		"36 37 38 39 40",
	}

	parseBoards(mockLines)

	expectedBoards := []board{
		{
			nums: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			marks: [][]bool{
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, false, false, false},
			},
			colSum: []int{0, 0, 0, 0, 0},
			rowSum: []int{0, 0, 0, 0, 0},
			hitSum: 0,
			done:   false,
		},
		{
			nums: [][]int{
				{6, 7, 8, 9, 10},
				{21, 22, 23, 24, 25},
				{26, 27, 28, 29, 30},
				{31, 32, 33, 34, 35},
				{36, 37, 38, 39, 40},
			},
			marks: [][]bool{
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, false, false, false},
			},
			colSum: []int{0, 0, 0, 0, 0},
			rowSum: []int{0, 0, 0, 0, 0},
			hitSum: 0,
			done:   false,
		},
	}

	for i, b := range boards {
		if !reflect.DeepEqual(*b, expectedBoards[i]) {
			t.Errorf("want: %v; got: %v", expectedBoards[i], *b)
		}
	}
}

func Test_start(t *testing.T) {
	mockLines := []string{
		"6,1,2,3,4,5,7,8,9,10",
		"",
		" 1  2  3  4  5",
		" 6  7  8  9 10",
		"11 12 13 14 15",
		"16 17 18 19 20",
		"21 22 23 24 25",
		"",
		" 6  7  8  9 10",
		"21 22 23 24 25",
		"26 27 28 29 30",
		"31 32 33 34 35",
		"36 37 38 39 40",
	}

	expectedResult := [2]int{1520, 6100}

	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{"ok", args{mockLines}, expectedResult},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := start(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("start() = %v, want %v", got, tt.want)
			}
		})
	}
}
