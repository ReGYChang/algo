package main

import "testing"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Backspace String Compare.
//Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Backspace String Compare.

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "844-1",
			args: args{
				s: "ab#c",
				t: "ad#c",
			},
			want: true,
		},
		{
			name: "844-2",
			args: args{
				s: "ab##",
				t: "c#d#",
			},
			want: true,
		},
		{
			name: "844-3",
			args: args{
				s: "a#c",
				t: "b",
			},
			want: false,
		},
		{
			name: "844-4",
			args: args{
				s: "a#cb",
				t: "#b",
			},
			want: false,
		},
		{
			name: "844-5",
			args: args{
				s: "##cb",
				t: "#b",
			},
			want: false,
		},
		{
			name: "844-6",
			args: args{
				s: "xywrrmp",
				t: "xywrrmu#p",
			},
			want: true,
		},
		{
			name: "844-7",
			args: args{
				s: "y#fo##f",
				t: "y#f#o##f",
			},
			want: true,
		},
		{
			name: "844-8",
			args: args{
				s: "abc#",
				t: "bac#",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}
