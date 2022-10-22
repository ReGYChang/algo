package main

import "testing"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Last Stone Weight.
//Memory Usage: 1.9 MB, less than 82.70% of Go online submissions for Last Stone Weight.

func Test_lastStoneWeight(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1046-1",
			args: args{
				stones: []int{2, 7, 4, 1, 8, 1},
			},
			want: 1,
		},
		{
			name: "1046-2",
			args: args{
				stones: []int{1},
			},
			want: 1,
		},
		{
			name: "1046-3",
			args: args{
				stones: []int{3, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeight(tt.args.stones); got != tt.want {
				t.Errorf("lastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
