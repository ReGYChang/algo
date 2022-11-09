package main

//Runtime: 0 ms, faster than 100% of Go online submissions.
//Memory Usage: 2.5 MB, less than 69.90% of Go online submissions.

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "190-1",
			args: args{
				num: 43261596,
			},
			want: 964176192,
		},
		{
			name: "190-2",
			args: args{
				num: 4294967293,
			},
			want: 3221225471,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
