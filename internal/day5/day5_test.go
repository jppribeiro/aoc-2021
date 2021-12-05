package day5

import (
	"reflect"
	"testing"
)

var mock = []string{
	"0,9 -> 5,9",
	"8,0 -> 0,8",
	"9,4 -> 3,4",
	"3,3 -> 3,9",
}

var mockLines = []line{
	{0, 5, 9, 9, true},
	{8, 0, 0, 8, false},
	{9, 3, 4, 4, true},
	{3, 3, 3, 9, true},
}

var mockWorld = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func Test_buildLines(t *testing.T) {
	expectedOrtho := []line{
		{0, 5, 9, 9, true},
		{9, 3, 4, 4, true},
		{3, 3, 3, 9, true},
	}

	expectedNonOrtho := []line{
		{0, 5, 9, 9, true},
		{8, 0, 0, 8, false},
		{9, 3, 4, 4, true},
		{3, 3, 3, 9, true},
	}

	type args struct {
		lines []string
		ortho bool
	}
	tests := []struct {
		name string
		args args
		want []line
	}{
		{"ok", args{mock, true}, expectedOrtho},
		{"ok", args{mock, false}, expectedNonOrtho},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildLines(tt.args.lines, tt.args.ortho); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_worldSize(t *testing.T) {
	expected := []int{9, 9}

	type args struct {
		lines []line
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"ok", args{mockLines}, expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := worldSize(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("worldSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trace(t *testing.T) {
	var mockOrthoLines = []line{
		{0, 5, 9, 9, true},
		{9, 3, 4, 4, true},
		{3, 3, 3, 9, true},
	}

	var expectedWorld = [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 2, 1, 1, 0, 0, 0, 0},
	}

	type args struct {
		world [][]int
		lines []line
	}
	tests := []struct {
		name string
		args args
	}{
		{"ok", args{mockWorld, mockOrthoLines}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trace(tt.args.world, tt.args.lines)

			if !reflect.DeepEqual(mockWorld, expectedWorld) {
				t.Errorf("expected: %v; got: %v", expectedWorld, mockWorld)
			}
		})
	}
}

func Test_analyze(t *testing.T) {
	type args struct {
		lines []string
		ortho bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ortho", args{mock, true}, 2},
		{"all", args{mock, false}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := analyze(tt.args.lines, tt.args.ortho); got != tt.want {
				t.Errorf("analyze() = %v, want %v", got, tt.want)
			}
		})
	}
}
