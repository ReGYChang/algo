package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement Queue using Stacks.
//Memory Usage: 1.9 MB, less than 85.17% of Go online submissions for Implement Queue using Stacks.

func TestConstructor(t *testing.T) {
	obj := Constructor()
	assert.Equal(t, obj, MyQueue{queue: []int{}})

	obj.Push(2)
	assert.Equal(t, obj, MyQueue{queue: []int{2}})

	obj.Push(10)
	assert.Equal(t, obj, MyQueue{queue: []int{2, 10}})

	param2 := obj.Pop()
	assert.Equal(t, param2, 2)

	param3 := obj.Peek()
	assert.Equal(t, param3, 10)

	param4 := obj.Empty()
	assert.Equal(t, param4, false)
}
