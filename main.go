package main

import (
	doublylinkedlist "dsa-in-go/doubly-linkedlist"
	"fmt"
	// linkedlist "dsa-in-go/linked-list"
)

func main(){
  fmt.Println("Doubly Linked List")
  dll := doublylinkedlist.CreateDoublyLinkedList()
  // dll.AddFirst(1)
  // dll.AddFirst(2)
  // dll.AddFirst(3)
  // dll.AddLast(4)
  err := dll.DeleteAtPos(0)
  if err != nil {
    fmt.Print("Error :",err)
  }
  dll.PrintLinkedList()

  fmt.Println("LINKED LIST")
  // head = head.AddFirst(0)
  // head.AddAtPosition(2,12)
  // head.AddLast(11)
  // head.PrintLinkedList()
  // fmt.Println("\nLength of linkedlist is ",head.GetLength())
}
