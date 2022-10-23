package main

import (
	"testing"

	"algo/structure"
)

//Runtime: 18 ms, faster than 31.96% of Go online submissions for Linked List Cycle.
//Memory Usage: 6.4 MB, less than 16.93% of Go online submissions for Linked List Cycle.

func Test_hasCycleHashMap(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "141-1",
			args: args{
				head: []int{3, 2, 0, -4},
			},
			want: false,
		},
		{
			name: "141-2",
			args: args{
				head: []int{1, 2},
			},
			want: false,
		},
		{
			name: "141-3",
			args: args{
				head: []int{1},
			},
			want: false,
		},
		{
			name: "141-4",
			args: args{
				head: []int{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := structure.Int2List(tt.args.head)
			if got := hasCycleHashMap(head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Linked List Cycle.
//Memory Usage: 4.3 MB, less than 100.00% of Go online submissions for Linked List Cycle.

func Test_hasCycleTwoPointer(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "141-1",
			args: args{
				head: []int{3, 2, 0, -4},
			},
			want: false,
		},
		{
			name: "141-2",
			args: args{
				head: []int{1, 2},
			},
			want: false,
		},
		{
			name: "141-3",
			args: args{
				head: []int{1},
			},
			want: false,
		},
		{
			name: "141-4",
			args: args{
				head: []int{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := structure.Int2List(tt.args.head)
			if got := hasCycleTwoPointer(head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
