package queue

import "fmt"

type CircularQueue[T any] struct {
	items []T
	start int
	end   int
	size  int
	count int
}

func MakeCircularQueue[T any](size int) *CircularQueue[T] {
	return &CircularQueue[T]{
		size:  size,
		items: make([]T, size),
		start: 0,
		end:   -1,
		count: 0,
	}
}

func (cq *CircularQueue[T]) Enqueue(val T) error {
	if cq.IsFull() {
		return fmt.Errorf("queue is full")
	}
	cq.end = (cq.end + 1) % cq.size
	cq.items[cq.end] = val
	cq.count++
	return nil
}

func (cq *CircularQueue[T]) Dequeue() (T, error) {
	var zeroValue T
	if cq.IsEmpty() {
		return zeroValue, fmt.Errorf("queue is empty")
	}
	value := cq.items[cq.start]
	cq.start = (cq.start + 1) % cq.size
	cq.count--
	return value, nil
}

func (cq *CircularQueue[T]) Peek() (T, error) {
	var zeroValue T
	if cq.IsEmpty() {
		return zeroValue, fmt.Errorf("queue is empty")
	}
	return cq.items[cq.start], nil
}

func (cq *CircularQueue[T]) IsFull() bool {
	return cq.count == cq.size
}

func (cq *CircularQueue[T]) IsEmpty() bool {
	return cq.count == 0
}
