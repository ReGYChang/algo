package main

import "testing"

//Runtime: 25 ms, faster than 75.69% of Go online submissions.
//Memory Usage: 6 MB, less than 99.31% of Go online submissions.

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "169-1",
			args: args{
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
		{
			name: "169-2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
