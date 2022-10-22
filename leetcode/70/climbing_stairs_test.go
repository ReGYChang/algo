package main

import "testing"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
//Memory Usage: 1.9 MB, less than 91.21% of Go online submissions for Climbing Stairs.

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "70 - 1",
			args: args{2},
			want: 2,
		},
		{
			name: "70 - 2",
			args: args{3},
			want: 3,
		},
		{
			name: "70 - 3",
			args: args{4},
			want: 5,
		},
		{
			name: "70 - 4",
			args: args{5},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
