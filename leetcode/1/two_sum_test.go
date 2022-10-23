package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"two sum", args{[]int{2, 7, 11, 5}, 9}, []int{0, 1}},
		{"two sum", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{"two sum", args{[]int{3, 3}, 6}, []int{0, 1}},
		{"two sum", args{[]int{-2, 7, -11, 5}, -4}, []int{1, 2}},
		{"two sum", args{[]int{3, 5, 16, -9}, 7}, []int{2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Runtime: 9 ms, faster than 65.01% of Go online submissions for Two Sum.
// Memory Usage: 4.4 MB, less than 29.00% of Go online submissions for Two Sum.
