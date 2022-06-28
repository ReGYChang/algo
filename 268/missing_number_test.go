package main

import (
	"testing"
)

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"missing number", args{nums: []int{3, 0, 1}}, 2},
		{"missing number", args{nums: []int{0, 1}}, 2},
		{"missing number", args{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Runtime: 40 ms, faster than 8.86% of Go online submissions for Missing Number.
// Memory Usage: 6.8 MB, less than 10.80% of Go online submissions for Missing Number.

func Test_missingNumber2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"missing number", args{nums: []int{3, 0, 1}}, 2},
		{"missing number", args{nums: []int{0, 1}}, 2},
		{"missing number", args{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber2(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Runtime: 29 ms, faster than 28.98% of Go online submissions for Missing Number.
// Memory Usage: 6.2 MB, less than 97.54% of Go online submissions for Missing Number.
