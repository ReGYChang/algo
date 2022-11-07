package main

import "testing"

//Runtime: 0 ms, faster than 100.00% of Go online submissions.
//Memory Usage: 2 MB, less than 85.65% of Go online submissions.

func Test_titleToNumber(t *testing.T) {
	type args struct {
		columnTitle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "171-1",
			args: args{
				columnTitle: "A",
			},
			want: 1,
		},
		{
			name: "171-2",
			args: args{
				columnTitle: "AB",
			},
			want: 28,
		},
		{
			name: "171-3",
			args: args{
				columnTitle: "ZY",
			},
			want: 701,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.columnTitle); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
