package main

import (
	"reflect"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"find all numbers disappeared numbers in an array", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{5, 6}},
		{"find all numbers disappeared numbers in an array", args{[]int{1, 1}}, []int{2}},
		{"find all numbers disappeared numbers in an array", args{[]int{1, 2, 2, 2, 2}}, []int{3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDisappearedNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Runtime: 118 ms, faster than 15.90% of Go online submissions for Find All Numbers Disappeared in an Array.
// Memory Usage: 9.5 MB, less than 7.23% of Go online submissions for Find All Numbers Disappeared in an Array.

func Test_findDisappearedNumbers2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"find all numbers disappeared numbers in an array", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{5, 6}},
		{"find all numbers disappeared numbers in an array", args{[]int{1, 1}}, []int{2}},
		{"find all numbers disappeared numbers in an array", args{[]int{1, 2, 2, 2, 2}}, []int{3, 4, 5}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDisappearedNumbers2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Runtime: 51 ms, faster than 88.73% of Go online submissions for Find All Numbers Disappeared in an Array.
// Memory Usage: 7.5 MB, less than 88.44% of Go online submissions for Find All Numbers Disappeared in an Array.
