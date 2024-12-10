package stack

import "fmt"

type Stack[T any] struct {
	items []T
}

func MakeStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: []T{},
	}
}

func (s *Stack[T]) Push(val T) {
	s.items = append(s.items, val)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}
