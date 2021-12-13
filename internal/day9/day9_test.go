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