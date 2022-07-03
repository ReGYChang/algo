package main

import "testing"

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"single number", args{[]int{2, 2, 1}}, 1},
		{"single number", args{[]int{4, 1, 2, 1, 2}}, 4},
		{"single number", args{[]int{1}}, 1},
		{"single number", args{[]int{3, 2, 3, 4, 4, 1, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Runtime: 16 ms, faster than 89.11% of Go online submissions for Single Number.
//Memory Usage: 6.1 MB, less than 94.51% of Go online submissions for Single Number.
