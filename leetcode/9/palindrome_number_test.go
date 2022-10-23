package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"palindrome number", args{121}, true},
		{"palindrome number", args{-121}, false},
		{"palindrome number", args{10}, false},
		{"palindrome number", args{1111}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Runtime: 19 ms, faster than 87.13% of Go online submissions for Palindrome Number.
// Memory Usage: 4.7 MB, less than 48.04% of Go online submissions for Palindrome Number.
