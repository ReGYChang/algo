package main

import (
	"reflect"
	"testing"

	"algo/structure"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "876-1",
			args: args{list: []int{1, 2, 3, 4, 5}},
			want: []int{3, 4, 5},
		},
		{
			name: "876-2",
			args: args{list: []int{1, 2, 3, 4, 5, 6}},
			want: []int{4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := structure.Int2List(tt.args.list)
			want := structure.Int2List(tt.want)
			if got := middleNode(head); !reflect.DeepEqual(got, want) {
				t.Errorf("middleNode() = %v, want %v", got, want)
			}
		})
	}
}
