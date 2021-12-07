package day6

import (
	"reflect"
	"testing"
)

func Test_reproduce(t *testing.T) {
	population = []int{3,4,3,1,2}
	type args struct {
		days uint
	}
	tests := []struct {
		name string
		args args
	}{
		{"ok", args{8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reproduce(tt.args.days)

			expected := []int{2,3,2,0,1,2,3,4,4,5}

			if !reflect.DeepEqual(population, expected) {
				t.Errorf("want: %v: got: %v", expected, population)
			}
		})
	}
}
