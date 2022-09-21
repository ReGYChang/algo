package main

import "testing"

//Runtime: 33 ms, faster than 91.93% of Go online submissions for Binary Search.
//Memory Usage: 6.6 MB, less than 94.44% of Go online submissions for Binary Search.

func Test_search_704(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "binary search 1",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			name: "binary search 2",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			want: -1,
		},
		{
			name: "binary search 3",
			args: args{
				nums:   []int{-1},
				target: 0,
			},
			want: -1,
		},
		{
			name: "binary search 4",
			args: args{
				nums:   []int{-1},
				target: -1,
			},
			want: 0,
		},
		{
			name: "binary search 5",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 12,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
