package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

//Runtime: 76 ms, faster than 34.35% of Go online submissions.
//Memory Usage: 7.5 MB, less than 44.19% of Go online submissions.

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		listA []int
		listB []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "160-1",
			args: args{
				listA: []int{4, 1},
				listB: []int{5, 6, 1},
			},
			want: []int{8, 4, 5},
		},
		{
			name: "160-2",
			args: args{
				listA: []int{1, 9, 1},
				listB: []int{3},
			},
			want: []int{2, 4},
		},
		{
			name: "160-3",
			args: args{
				listA: []int{2, 6, 4},
				listB: []int{1, 5},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tail := structure.Int2List(tt.want)

			headA := structure.Int2List(tt.args.listA)
			tailA := structure.TailOfList(headA)
			tailA.Next = tail

			headB := structure.Int2List(tt.args.listB)
			tailB := structure.TailOfList(headB)
			tailB.Next = tail

			want := structure.Int2List(tt.want)

			if got := getIntersectionNode(headA, headB); !reflect.DeepEqual(got, want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, want)
			}
		})
	}
}
