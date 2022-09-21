package main

import "testing"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for First Bad Version.
//Memory Usage: 1.8 MB, less than 87.98% of Go online submissions for First Bad Version.

func Test_firstBadVersion_278(t *testing.T) {
	type args struct {
		n   int
		bad int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first bad version",
			args: args{n: 5},
			want: 4,
		},
		{
			name: "first bad version 2",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "first bad version 3",
			args: args{n: 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		Bad = tt.want
		t.Run(tt.name, func(t *testing.T) {
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
