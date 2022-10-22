package main

import "testing"

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Repeating Character Replacement.
// Memory Usage: 2.5 MB, less than 79.23% of Go online submissions for Longest Repeating Character Replacement.

func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "424-1",
			args: args{
				s: "ABAB",
				k: 2,
			},
			want: 4,
		},
		{
			name: "424-2",
			args: args{
				s: "AABABBA",
				k: 1,
			},
			want: 4,
		},
		{
			name: "424-3",
			args: args{
				s: "ABC",
				k: 0,
			},
			want: 1,
		},
		{
			name: "424-4",
			args: args{
				s: "ABCDE",
				k: 3,
			},
			want: 4,
		},
		{
			name: "424-5",
			args: args{
				s: "AAAA",
				k: 2,
			},
			want: 4,
		},
		{
			name: "424-6",
			args: args{
				s: "AAAA",
				k: 0,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
