package main

import "testing"

//Runtime: 5 ms, faster than 94.03% of Go online submissions for First Unique Character in a String.
//Memory Usage: 5.3 MB, less than 94.36% of Go online submissions for First Unique Character in a String.

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "387-1",
			args: args{"leetcode"},
			want: 0,
		},
		{
			name: "387-2",
			args: args{"loveleetcode"},
			want: 2,
		},
		{
			name: "387-3",
			args: args{"aabb"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
