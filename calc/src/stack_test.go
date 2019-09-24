package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	var stack Stack

	stack.Push(1)
	stack.Push("+")

	elem, _ := stack.Pop()
	elem, _ = elem.(string)
	assert.Equal(t, "+", elem)

	elem, _ = stack.Pop()
	elem, _ = elem.(int)
	assert.Equal(t, 1, elem)

	assert.True(t, stack.IsEmpty())

	_, err := stack.Pop()
	assert.Equal(t, ErrorPopEmpty, err)
}
