package main

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"contains duplicate", args{[]int{1, 2, 3, 1}}, true},
		{"contains duplicate", args{[]int{1, 2, 3, 4}}, false},
		{"contains duplicate", args{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, true},
		{"contains duplicate", args{[]int{1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
