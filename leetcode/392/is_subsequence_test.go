package main

import "testing"

func Test_isSubsequence(t *testing.T) {
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
			name: "Is subsequence1",
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		}, {
			name: "Is subsequence2",
			args: args{
				s: "axc",
				t: "ahbgdc",
			},
			want: false,
		}, {
			name: "Is subsequence3",
			args: args{
				s: "b",
				t: "c",
			},
			want: false,
		}, {
			name: "Is subsequence4",
			args: args{
				s: "",
				t: "ahbgdc",
			},
			want: true,
		}, {
			name: "Is subsequence5",
			args: args{
				s: "",
				t: "",
			},
			want: true,
		}, {
			name: "Is subsequence6",
			args: args{
				s: "abc",
				t: "",
			},
			want: false,
		}, {
			name: "Is subsequence7",
			args: args{
				s: "acb",
				t: "ahbgdc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Runtime: 3 ms, faster than 24.32% of Go online submissions for Is Subsequence.
//Memory Usage: 1.9 MB, less than 81.38% of Go online submissions for Is Subsequence.

func Test_isSubsequence2(t *testing.T) {
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
			name: "Is subsequence1",
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		}, {
			name: "Is subsequence2",
			args: args{
				s: "axc",
				t: "ahbgdc",
			},
			want: false,
		}, {
			name: "Is subsequence3",
			args: args{
				s: "b",
				t: "c",
			},
			want: false,
		}, {
			name: "Is subsequence4",
			args: args{
				s: "",
				t: "ahbgdc",
			},
			want: true,
		}, {
			name: "Is subsequence5",
			args: args{
				s: "",
				t: "",
			},
			want: true,
		}, {
			name: "Is subsequence6",
			args: args{
				s: "abc",
				t: "",
			},
			want: false,
		}, {
			name: "Is subsequence7",
			args: args{
				s: "acb",
				t: "ahbgdc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence2(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Runtime: 2 ms, faster than 47.70% of Go online submissions for Is Subsequence.
//Memory Usage: 3 MB, less than 5.10% of Go online submissions for Is Subsequence.
