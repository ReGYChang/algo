package main

import (
	"reflect"
	"testing"
)

//Runtime: 15 ms, faster than 67.29% of Go online submissions for Find All Anagrams in a String.
//Memory Usage: 5.1 MB, less than 56.20% of Go online submissions for Find All Anagrams in a String.

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "438-1",
			args: args{
				s: "cbaebabacd",
				p: "abc",
			},
			want: []int{0, 6},
		},
		{
			name: "438-2",
			args: args{
				s: "abab",
				p: "ab",
			},
			want: []int{0, 1, 2},
		},
		{
			name: "438-3",
			args: args{
				s: "ababababab",
				p: "aab",
			},
			want: []int{0, 2, 4, 6},
		},
		{
			name: "438-4",
			args: args{
				s: "aaaaaaaaaa",
				p: "aaaaaaaaaaaaa",
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
