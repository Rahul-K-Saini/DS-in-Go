package main

import (
	linkedlist "dsa-in-go/linked-list"
	"fmt"
)

func main(){
  head := linkedlist.CreateLinkedList(10)

  // curr := head
  // for curr.Value != 8 {
  //   curr = curr.Next
  // }
  // head.DeleteNode(curr)

  fmt.Println("LINKED LIST")
  head = head.AddFirst(0)
  head.AddAtPosition(2,12)
  head.AddLast(11)
  head.PrintLinkedList()
  fmt.Println("\nLength of linkedlist is ",head.GetLength())
}
