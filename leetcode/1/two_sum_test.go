package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1-1",
			args: args{
				nums:   []int{2, 7, 11, 5},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "1-2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "1-3",
			args: args{
				nums:   []int{-2, 7, -11, 5},
				target: -4,
			},
			want: []int{1, 2},
		},
		{
			name: "1-4",
			args: args{
				nums:   []int{3, 5, 16, -9},
				target: 21,
			},
			want: []int{1, 2},
		},
		{
			name: "1-5",
			args: args{
				nums:   []int{3, 2, 4},
				target: 8,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
