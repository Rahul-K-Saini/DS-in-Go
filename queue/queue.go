package queue

import "fmt"

type Queue[T any] struct {
	items []T
	start int
	end   int
	size  int
}

func MakeQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		items: make([]T, size),
		start: -1,
		end:   -1,
		size:  size,
	}
}

func (q *Queue[T]) Add(val T) error {
	if q.end == q.size-1 {
		return fmt.Errorf("queue is full")
	}
	if q.start == -1 {
		q.start = 0
	}
	q.end++
	q.items[q.end] = val
	return nil
}

func (q *Queue[T]) Delete() (T, error) {
	if q.start == -1 || q.start > q.end {
		var zeroValue T
		return zeroValue, fmt.Errorf("queue is empty")
	}
	val := q.items[q.start]
	q.start++
	if q.start > q.end {
		q.start = -1
		q.end = -1
	}
	return val, nil
}

func (q *Queue[T]) Peek() (T, error) {
	var value T
	if q.start == -1 {
		return value, fmt.Errorf("queue is empty")
	}
	value = q.items[q.end]
	return value, nil
}
