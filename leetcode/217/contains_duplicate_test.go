package main

import "testing"

//Runtime: 92 ms, faster than 80.51% of Go online submissions for Contains Duplicate.
//Memory Usage: 8.9 MB, less than 66.21% of Go online submissions for Contains Duplicate.

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "217-1",
			args: args{[]int{1, 2, 3, 1}},
			want: true,
		},
		{
			name: "217-2",
			args: args{[]int{1, 2, 3, 4}},
			want: false,
		},
		{
			name: "217-3",
			args: args{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			want: true,
		},
		{
			name: "217-4",
			args: args{[]int{1}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
