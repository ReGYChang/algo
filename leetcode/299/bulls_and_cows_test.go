package main

import "testing"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Bulls and Cows.
//Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Bulls and Cows.

func Test_getHint(t *testing.T) {
	type args struct {
		secret string
		guess  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "299-1",
			args: args{
				secret: "1807",
				guess:  "7810",
			},
			want: "1A3B",
		},
		{
			name: "299-2",
			args: args{
				secret: "1123",
				guess:  "0111",
			},
			want: "1A1B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHint(tt.args.secret, tt.args.guess); got != tt.want {
				t.Errorf("getHint() = %v, want %v", got, tt.want)
			}
		})
	}
}
