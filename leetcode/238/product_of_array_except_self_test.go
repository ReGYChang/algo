package main

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"productExceptSelf", args{[]int{1, 2, 3, 4}}, []int{24, 12, 8, 6}},
		{"productExceptSelf", args{[]int{-1, 1, 0, -3, 3}}, []int{0, 0, 9, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Runtime: 59 ms, faster than 19.27% of Go online submissions for Product of Array Except Self.
//Memory Usage: 8.2 MB, less than 30.78% of Go online submissions for Product of Array Except Self.
