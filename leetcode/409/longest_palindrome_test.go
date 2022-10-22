package main

import "testing"

// Runtime: 3 ms, faster than 57.51% of Go online submissions for Longest Palindrome.
// Memory Usage: 2.1 MB, less than 82.31% of Go online submissions for Longest Palindrome.

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "longest palindrome_1",
			args: args{
				s: "abccccdd",
			},
			want: 7,
		},
		{
			name: "longest palindrome_2",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "longest palindrome_3",
			args: args{
				s: "aaa",
			},
			want: 3,
		},
		{
			name: "longest palindrome_3",
			args: args{
				s: "aaaccc",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
