package main

import "testing"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths.
//Memory Usage: 1.9 MB, less than 75.66% of Go online submissions for Unique Paths.

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "62 - 1",
			args: args{
				m: 3,
				n: 7,
			},
			want: 28,
		},
		{
			name: "62 - 2",
			args: args{
				m: 3,
				n: 2,
			},
			want: 3,
		},
		{
			name: "62 - 3",
			args: args{
				m: 1,
				n: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
