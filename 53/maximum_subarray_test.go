package main

import "testing"

//Runtime: 169 ms, faster than 64.52% of Go online submissions for Maximum Subarray.
//Memory Usage: 9.4 MB, less than 58.13% of Go online submissions for Maximum Subarray.

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "53-1",
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			want: 6,
		},
		{
			name: "53-2",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "53-3",
			args: args{nums: []int{5, 4, -1, 7, 8}},
			want: 23,
		},
		{
			name: "53-4",
			args: args{nums: []int{-1, -2}},
			want: -1,
		},
		{
			name: "53-5",
			args: args{nums: []int{1, 2}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
