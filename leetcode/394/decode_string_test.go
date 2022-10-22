package main

import "testing"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Decode String.
//Memory Usage: 2 MB, less than 93.38% of Go online submissions for Decode String.

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "394-1",
			args: args{
				s: "3[a]2[bc]",
			},
			want: "aaabcbc",
		},
		{
			name: "394-2",
			args: args{
				s: "3[a2[c]]",
			},
			want: "accaccacc",
		},
		{
			name: "394-3",
			args: args{
				s: "2[abc]3[cd]ef",
			},
			want: "abcabccdcdcdef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
