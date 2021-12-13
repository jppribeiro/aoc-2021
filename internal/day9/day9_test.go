package day9

import (
	"reflect"
	"testing"
)

func Test_parseHeightMap(t *testing.T) {
	lines := []string{
		"11111",
		"12111",
		"11311",
	}

	expect := [][]int{
		{9,9,9,9,9,9,9},
		{9,1,1,1,1,1,9},
		{9,1,2,1,1,1,9},
		{9,1,1,3,1,1,9},
		{9,9,9,9,9,9,9},
	}

	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{ "ok", args{lines: lines}, expect},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := parseHeightMap(tt.args.lines); !reflect.DeepEqual(got, tt.want) {

				t.Errorf("parseHeightMap() = %v, want %v", got, tt.want)

			}

		})
	}
}

func Test_part1(t *testing.T) {
	heightMap = [][]int{
		{2,1,9,9,9,4,3,2,1,0},
		{3,9,8,7,8,9,4,9,2,1},
		{9,8,5,6,7,8,9,8,9,2},
		{8,7,6,7,8,9,6,7,8,9},
		{9,8,9,9,9,6,5,6,7,8},
	}
	tests := []struct {
		name string
		want int
	}{
		{"ok", 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
/*

		{0,0,0,0,0,0,0,0,0,0,0,0},
		{0,0,0,1,0,0,0,0,0,0,0,0},
		{0,0,1,1,1,0,0,0,0,0,0,0},
		{0,0,0,1,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0,0,0},

 */
func Test_basin(t *testing.T) {
	heightMap = [][]int{
		{9,9,9,9,9,9,9,9,9,9,9,9},
		{9,2,1,9,9,9,4,3,2,1,0,9},
		{9,3,9,8,7,8,9,4,9,2,1,9},
		{9,9,8,5,6,7,8,9,8,9,2,9},
		{9,8,7,6,7,8,9,6,7,8,9,9},
		{9,9,8,9,9,9,6,5,6,7,8,9},
		{9,9,9,9,9,9,9,9,9,9,9,9},
	}

	type args struct {
		low [2]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ok", args{[2]int{3,3}}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basin(tt.args.low); got != tt.want {
				t.Errorf("basin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lows(t *testing.T) {
	heightMap = [][]int{
		{9,9,9,9,9,9,9,9,9,9,9,9},
		{9,2,1,9,9,9,4,3,2,1,0,9},
		{9,3,9,8,7,8,9,4,9,2,1,9},
		{9,9,8,5,6,7,8,9,8,9,2,9},
		{9,8,7,6,7,8,9,6,7,8,9,9},
		{9,9,8,9,9,9,6,5,6,7,8,9},
		{9,9,9,9,9,9,9,9,9,9,9,9},
	}

	expect := [][2]int{
		{1,2},
		{1, 10},
		{3,3},
		{5,7},
	}

	tests := []struct {
		name string
		want [][2]int
	}{
		{"ok", expect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lows(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lows() = %v, want %v", got, tt.want)
			}
		})
	}
}