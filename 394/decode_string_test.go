package main

import "testing"

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
			want: "aaabcbc",
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
