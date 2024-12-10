package doublylinkedlist

import (
	"bytes"
	"fmt"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type DoublyLinkedList[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

func MakeDoublyLL[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{Head: nil, Tail: nil, Length: 0}
}

func (list *DoublyLinkedList[T]) AddLast(value T) {
	newNode := &Node[T]{Value: value}

	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		list.Tail.Next = newNode
		newNode.Prev = list.Tail
		list.Tail = newNode
	}
	list.Length++
}

func (list *DoublyLinkedList[T]) AddFirst(value T) {
	newNode := &Node[T]{Value: value}

	if list.Length == 0 {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.Next = list.Head
		list.Head.Prev = newNode
		list.Head = newNode
	}
	list.Length++
}

func (list *DoublyLinkedList[T]) DeleteFirst() error {
	if list.Length == 0 {
		return fmt.Errorf("list is empty")
	}
	if list.Length == 1 {
		list.Head = nil
		list.Tail = nil
		list.Length = 0
		return nil
	}
	list.Head = list.Head.Next
	list.Head.Prev = nil
	list.Length--
	return nil
}

func (list *DoublyLinkedList[T]) DeleteLast() error {
	if list.Length == 0 {
		return fmt.Errorf("list is empty")
	}
	if list.Length == 1 {
		list.Head = nil
		list.Tail = nil
		list.Length = 0
		return nil
	}
	list.Tail = list.Tail.Prev
	list.Tail.Next = nil
	list.Length--
	return nil
}

func (list *DoublyLinkedList[T]) DeleteAtPos(pos int) error {
	if list.Length == 0 {
		return fmt.Errorf("list is empty")
	}
	if pos < 0 || pos >= list.Length {
		return fmt.Errorf("invalid position")
	}
	if pos == 0 {
		return list.DeleteFirst()
	}
	if pos == list.Length-1 {
		return list.DeleteLast()
	}

	curr := list.Head
	for i := 0; i < pos; i++ {
		curr = curr.Next
	}
	curr.Prev.Next = curr.Next
	curr.Next.Prev = curr.Prev
	list.Length--
	return nil
}

func (list *DoublyLinkedList[T]) PrintLinkedList() {
	var buffer bytes.Buffer
	current := list.Head
	buffer.WriteString("Head -> ")
	for current != nil {
		buffer.WriteString(fmt.Sprintf("%v <-> ", current.Value))
		current = current.Next
	}
	buffer.WriteString("Tail")
	fmt.Println(buffer.String())
}
