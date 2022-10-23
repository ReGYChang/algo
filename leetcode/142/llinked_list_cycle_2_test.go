package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

//Runtime: 14 ms, faster than 28.12% of Go online submissions for Linked List Cycle II.
//Memory Usage: 5.2 MB, less than 12.07% of Go online submissions for Linked List Cycle II.

func Test_detectCycle(t *testing.T) {
	type args struct {
		list []int
		pos  int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "142-1",
			args: args{
				list: []int{3, 2, 0, -4},
				pos:  1,
			},
		},
		{
			name: "142-2",
			args: args{
				list: []int{1, 2, 3},
				pos:  -1,
			},
		},
		{
			name: "142-3",
			args: args{
				list: []int{3, 2, 0, -4},
				pos:  1,
			},
		},
		{
			name: "142-4",
			args: args{
				list: []int{1, 2},
				pos:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := structure.Int2ListWithCycle(tt.args.list, tt.args.pos)
			var ans *structure.ListNode
			if tt.args.pos >= 0 {
				ans = head.GetNodeWith(tt.args.list[tt.args.pos])
			}
			if got := detectCycle(head); !reflect.DeepEqual(got, ans) {
				t.Errorf("detectCycle() = %v, want %v", got, ans)
			}
		})
	}
}
