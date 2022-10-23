package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

//Runtime: 7 ms, faster than 90.19% of Go online submissions for Remove Linked List Elements.
// Memory Usage: 5 MB, less than 14.07% of Go online submissions for Remove Linked List Elements.

func Test_removeElements(t *testing.T) {
	type args struct {
		head []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "203-1",
			args: args{
				head: []int{1, 2, 6, 3, 4, 5, 6},
				val:  6,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "203-2",
			args: args{
				head: []int{},
				val:  1,
			},
			want: []int{},
		},
		{
			name: "203-3",
			args: args{
				head: []int{7, 7, 7, 7},
				val:  7,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := structure.Int2List(tt.args.head)
			if got := removeElements(head, tt.args.val); !reflect.DeepEqual(structure.List2Int(got), tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
