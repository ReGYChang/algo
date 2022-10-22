package main

import "testing"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
//Memory Usage: 1.9 MB, less than 99.53% of Go online submissions for Valid Parentheses.

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "20-1",
			args: args{s: "()"},
			want: true,
		},
		{
			name: "20-2",
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			name: "20-3",
			args: args{s: "(]"},
			want: false,
		},
		{
			name: "20-4",
			args: args{s: "([)]"},
			want: false,
		},
		{
			name: "20-5",
			args: args{s: "("},
			want: false,
		},
		{
			name: "20-6",
			args: args{s: "]"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
