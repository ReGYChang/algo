package main

import (
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"find pivot index", args{nums: []int{1, 7, 3, 6, 5, 6}}, 3},
		{"find pivot index", args{nums: []int{1, 2, 3}}, -1},
		{"find pivot index", args{nums: []int{2, 1, -1}}, 0},
		{"find pivot index", args{nums: []int{-1, -1, 0, 1, 1, 0}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Runtime: 451 ms, faster than 5.63% of Go online submissions for Find Pivot Index.
// Memory Usage: 6.5 MB, less than 43.38% of Go online submissions for Find Pivot Index.

func Test_pivotIndex2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"find pivot index", args{nums: []int{1, 7, 3, 6, 5, 6}}, 3},
		{"find pivot index", args{nums: []int{1, 2, 3}}, -1},
		{"find pivot index", args{nums: []int{2, 1, -1}}, 0},
		{"find pivot index", args{nums: []int{-1, -1, 0, 1, 1, 0}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex2(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Runtime: 42 ms, faster than 28.15% of Go online submissions for Find Pivot Index.
// Memory Usage: 6.4 MB, less than 59.77% of Go online submissions for Find Pivot Index.

func Test_pivotIndex3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"find pivot index", args{nums: []int{1, 7, 3, 6, 5, 6}}, 3},
		{"find pivot index", args{nums: []int{1, 2, 3}}, -1},
		{"find pivot index", args{nums: []int{2, 1, -1}}, 0},
		{"find pivot index", args{nums: []int{-1, -1, 0, 1, 1, 0}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex3(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex3() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Runtime: 27 ms, faster than 56.29% of Go online submissions for Find Pivot Index.
// Memory Usage: 6.5 MB, less than 28.48% of Go online submissions for Find Pivot Index.
