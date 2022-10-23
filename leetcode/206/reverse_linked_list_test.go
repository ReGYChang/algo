package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
//Memory Usage: 2.7 MB, less than 29.35% of Go online submissions for Reverse Linked List.

func Test_reverseList(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "206-1",
			args: args{list: []int{1, 2, 3, 4, 5}},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "206-2",
			args: args{list: []int{1, 2}},
			want: []int{2, 1},
		},
		{
			name: "206-3",
			args: args{list: []int{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := structure.Int2List(tt.args.list)
			want := structure.Int2List(tt.want)
			if got := reverseList(head); !reflect.DeepEqual(got, want) {
				t.Errorf("reverseList() = %v, want %v", got, want)
			}
		})
	}
}
