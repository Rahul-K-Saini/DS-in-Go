package doublylinkedlist

import (
	"bytes"
	"fmt"
)

type node struct {
	value int
	next  *node
	prev  *node
}

type DoublyLinkedList struct {
	head   *node
	tail   *node
  length int
}

func CreateDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{head: nil, tail: nil, length: 0}
}

func (list *DoublyLinkedList) AddLast(value int) {
	newNode := &node{value: value}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
	list.length++
}

func (list *DoublyLinkedList) AddFirst(value int) {
	newNode := &node{value: value}
	// this means the list is empty so new node will be the head and tail of list
	if list.length == 0 { 
		list.head = newNode
		list.tail = newNode
		list.length++
		return
	}
  newNode.next = list.head
  list.head = newNode
  list.length++
}

func (list *DoublyLinkedList) DeleteFirst() error {
  if list.length == 0 {
    return fmt.Errorf("list is empty")
  }
  if list.length == 1 {
    list.head = nil 
    list.tail = nil 
    list.length = 0
    return nil
  } 
  list.head = list.head.next
  list.head.prev = nil
  list.length--
  return nil
}

func (list *DoublyLinkedList) DeleteLast() error {
  if list.length == 0 {
    return fmt.Errorf("list is empty")
  }
  if list.length == 1 {
    list.head = nil 
    list.tail = nil 
    list.length = 0
    return nil
  } 
  list.tail = list.tail.prev
  list.tail.next = nil
  list.length--
  return nil
}

func (list *DoublyLinkedList) DeleteAtPos(pos int) error {
  if list.length == 0 {
    return fmt.Errorf("list is empty")
  }
  if pos < 0 || pos >= list.length {
    return fmt.Errorf("invalid position")
  }
  if pos == 0 {
    return list.DeleteFirst()
  }
  if pos == list.length-1 {
    return list.DeleteLast()
  }
  curr := list.head
  for i := 0; i < pos-1; i++ {
    curr = curr.next
  }
  curr.next = curr.next.next
  curr.next.prev = curr
  list.length--
  return nil
}

func (list *DoublyLinkedList) PrintLinkedList() {
    var buffer bytes.Buffer
    current := list.head
    buffer.WriteString("Head -> ")
    for current != nil {
        buffer.WriteString(fmt.Sprintf("%d <-> ", current.value))
        current = current.next
    }
    buffer.WriteString("Tail")
    fmt.Println(buffer.String())
}

