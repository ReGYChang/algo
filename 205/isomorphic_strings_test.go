package main

import "testing"

func Test_isIsomorphic(t *testing.T) {
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
			name: "isomorphic strings",
			args: args{
				s: "egg",
				t: "add",
			},
			want: true,
		}, {
			name: "isomorphic strings",
			args: args{
				s: "foo",
				t: "bar",
			},
			want: false,
		}, {
			name: "isomorphic strings",
			args: args{
				s: "paper",
				t: "title",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Runtime: 13 ms, faster than 19.41% of Go online submissions for Isomorphic Strings.
//Memory Usage: 2.6 MB, less than 85.04% of Go online submissions for Isomorphic Strings.
