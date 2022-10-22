package main

import "testing"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Anagram.
//Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Valid Anagram.

func Test_isAnagram(t *testing.T) {
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
			name: "242-1",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "242-2",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
		{
			name: "242-3",
			args: args{
				s: "ab",
				t: "a",
			},
			want: false,
		},
		{
			name: "242-4",
			args: args{
				s: "a",
				t: "ab",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
