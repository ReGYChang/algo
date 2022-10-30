package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Runtime: 8 ms, faster than 92.44% of Go online submissions for Remove Duplicates from Sorted Array.
//Memory Usage: 4.3 MB, less than 82.73% of Go online submissions for Remove Duplicates from Sorted Array.

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "26-1",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: []int{1, 2},
		},
		{
			name: "26-2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "26-3",
			args: args{
				nums: []int{-100, -100, 0, 1, 2, 3, 3, 4, 4, 4, 5, 5, 100},
			},
			want: []int{-100, 0, 1, 2, 3, 4, 5, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != len(tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i], tt.args.nums[i])
			}
		})
	}
}

//Runtime: 1198 ms, faster than 5.04% of Go online submissions for Remove Duplicates from Sorted Array.
//Memory Usage: 4.3 MB, less than 82.73% of Go online submissions for Remove Duplicates from Sorted Array.

func Test_removeDuplicatesBubble(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "26-1",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: []int{1, 2},
		},
		{
			name: "26-2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "26-3",
			args: args{
				nums: []int{-100, -100, 0, 1, 2, 3, 3, 4, 4, 4, 5, 5, 100},
			},
			want: []int{-100, 0, 1, 2, 3, 4, 5, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicatesBubble(tt.args.nums); got != len(tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i], tt.args.nums[i])
			}
		})
	}
}
