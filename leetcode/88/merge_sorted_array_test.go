package main

import (
	"reflect"
	"testing"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Sorted Array.
//Memory Usage: 2.2 MB, less than 85.10% of Go online submissions for Merge Sorted Array.

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "88-1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "88-2",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			want: []int{1},
		},
		{
			name: "88-3",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1},
		},
		{
			name: "88-4",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{0},
				n:     0,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("merge() = %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}
