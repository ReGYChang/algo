package main

import "testing"

//Runtime: 2 ms, faster than 49.28% of Go online submissions.
//Memory Usage: 2.1 MB, less than 67.78% of Go online submissions.

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "202-1",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "202-2",
			args: args{
				n: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
