package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

//Runtime: 3 ms, faster than 89.65% of Go online submissions for Remove Duplicates from Sorted List.
//Memory Usage: 3.1 MB, less than 85.61% of Go online submissions for Remove Duplicates from Sorted List.

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name         string
		args         args
		expectedList []int
	}{
		{
			name: "83-1",
			args: args{
				list: []int{1, 1, 2},
			},
			expectedList: []int{1, 2},
		},
		{
			name: "83-2",
			args: args{
				list: []int{1, 1, 2, 3, 3},
			},
			expectedList: []int{1, 2, 3},
		},
		{
			name: "83-3",
			args: args{
				list: []int{1, 1, 1},
			},
			expectedList: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := structure.Int2List(tt.args.list)
			expectedHead := structure.Int2List(tt.expectedList)
			if got := deleteDuplicates(head); !reflect.DeepEqual(got, expectedHead) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, expectedHead)
			}
		})
	}
}
