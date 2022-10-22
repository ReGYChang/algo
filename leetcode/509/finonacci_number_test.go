package main

import "testing"

//Runtime: 13 ms, faster than 21.35% of Go online submissions for Fibonacci Number.
//Memory Usage: 2 MB, less than 19.04% of Go online submissions for Fibonacci Number.

func Test_fibRecursive_509(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "finonacci number",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "finonacci number",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "finonacci number",
			args: args{n: 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibRecursive(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
