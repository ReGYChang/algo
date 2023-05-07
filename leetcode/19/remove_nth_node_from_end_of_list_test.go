package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		list []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "19-1",
			args: args{
				list: []int{1, 2, 3, 4, 5},
				n:    2,
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "19-2",
			args: args{
				list: []int{1},
				n:    1,
			},
			want: []int{},
		},
		{
			name: "19-3",
			args: args{
				list: []int{1, 2},
				n:    1,
			},
			want: []int{1},
		},
		{
			name: "19-4",
			args: args{
				list: []int{1, 2, 3, 4, 5},
				n:    1,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "19-5",
			args: args{
				list: []int{1, 2, 3},
				n:    3,
			},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := structure.Int2List(tt.args.list)
			want := structure.Int2List(tt.want)
			if got := removeNthFromEnd(head, tt.args.n); !reflect.DeepEqual(got, want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", structure.List2Int(got), structure.List2Int(want))
			}
		})
	}
}
