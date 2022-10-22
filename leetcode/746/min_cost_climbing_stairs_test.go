package main

import "testing"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Min Cost Climbing Stairs.
//Memory Usage: 3.1 MB, less than 43.33% of Go online submissions for Min Cost Climbing Stairs.

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "746 - 1",
			args: args{
				cost: []int{10, 15, 20},
			},
			want: 15,
		},
		{
			name: "746 - 2",
			args: args{
				cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			},
			want: 6,
		},
		{
			name: "746 - 3",
			args: args{
				cost: []int{0, 0, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
