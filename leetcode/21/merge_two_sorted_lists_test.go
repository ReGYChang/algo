package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
//Memory Usage: 2.5 MB, less than 51.14% of Go online submissions for Merge Two Sorted Lists.

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 []int
		list2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "21-1",
			args: args{
				list1: []int{},
				list2: []int{},
			},
			want: []int{},
		},
		{
			name: "21-2",
			args: args{
				list1: []int{1},
				list2: []int{1},
			},
			want: []int{1, 1},
		},
		{
			name: "21-3",
			args: args{
				list1: []int{1, 2, 3, 4},
				list2: []int{1, 2, 3, 4},
			},
			want: []int{1, 1, 2, 2, 3, 3, 4, 4},
		},
		{
			name: "21-4",
			args: args{
				list1: []int{1, 2, 3, 4, 5},
				list2: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := structure.Int2List(tt.args.list1)
			list2 := structure.Int2List(tt.args.list2)
			want := structure.Int2List(tt.want)
			if got := mergeTwoLists(list1, list2); !reflect.DeepEqual(got, want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, want)
			}
		})
	}
}
