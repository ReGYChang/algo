package main

import "testing"

//Runtime: 0 ms, faster than 100% of Go online submissions.
//Memory Usage: 1.9 MB, less than 92.10% of Go online submissions.

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "191-1",
			args: args{
				num: 11,
			},
			want: 3,
		},
		{
			name: "191-2",
			args: args{
				num: 128,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
