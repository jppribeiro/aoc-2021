package day10

import (
	"testing"
)

func Test_isCorrupt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"ok", args{"()"}, false},
		{"corrupt", args{"(>"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCorrupt(tt.args.s); got != tt.want {
				t.Errorf("isCorrupt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isStart(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"yes", args{"("}, true},
		{"no", args{">"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStart(tt.args.s); got != tt.want {
				t.Errorf("isStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEnd(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"yes", args{")"}, true},
		{"no", args{"{"}, false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEnd(tt.args.s); got != tt.want {
				t.Errorf("isEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resolveLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1197 pts", args{"{([(<{}[<>[]}>{[]{[(<()>"}, 1197},
		{"25137 pts", args{"<{([([[(<>()){}]>(<<{{"}, 25137},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolveLine(tt.args.line); got != tt.want {
				t.Errorf("resolveLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	lines := []string{
		"{([(<{}[<>[]}>{[]{[(<()>",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
	}
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ok", args{lines}, 26397},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.lines); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
