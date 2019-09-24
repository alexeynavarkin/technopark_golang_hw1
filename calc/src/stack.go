package main

import (
	"errors"
)

// Stack is custom type for stack
type Stack []interface{}

// Push is lol
func (s *Stack) Push(elem interface{}) {
	*s = append(*s, elem)
}

// ErrorPopEmpty is lol
var ErrorPopEmpty = errors.New("stack is empty")

// Pop is lol
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrorPopEmpty
	}
	elem := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return elem, nil
}

// IsEmpty is lol
func (s Stack) IsEmpty() bool {
	if len(s) > 0 {
		return false
	}
	return true
}
