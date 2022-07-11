package main

import "testing"

func Test_numberOfSteps(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test number of steps to reduce a number to zero", args{num: 14}, 6},
		{"test number of steps to reduce a number to zero", args{num: 8}, 4},
		{"test number of steps to reduce a number to zero", args{num: 123}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfSteps(tt.args.num); got != tt.want {
				t.Errorf("numberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of Steps to Reduce a Number to Zero.
// Memory Usage: 1.9 MB, less than 90% of Go online submissions for Number of Steps to Reduce a Number to Zero.
