package main

import "testing"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Sqrt(x).
//Memory Usage: 2 MB, less than 82.21% of Go online submissions for Sqrt(x).

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "69-1",
			args: args{
				x: 4,
			},
			want: 2,
		},
		{
			name: "69-2",
			args: args{
				x: 8,
			},
			want: 2,
		},
		{
			name: "69-3",
			args: args{
				x: 784,
			},
			want: 28,
		},
		{
			name: "69-4",
			args: args{
				x: 441,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
