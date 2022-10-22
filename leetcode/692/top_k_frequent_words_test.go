package main

import (
	"reflect"
	"testing"
)

//Runtime: 3 ms, faster than 98.02% of Go online submissions for Top K Frequent Words.
//Memory Usage: 4.3 MB, less than 76.98% of Go online submissions for Top K Frequent Words.

func Test_topKFrequent(t *testing.T) {
	type args struct {
		words []string
		k     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "692-1",
			args: args{
				words: []string{"i", "love", "leetcode", "i", "love", "coding"},
				k:     2,
			},
			want: []string{"i", "love"},
		},
		{
			name: "692-2",
			args: args{
				words: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
				k:     4,
			},
			want: []string{"the", "is", "sunny", "day"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.words, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
