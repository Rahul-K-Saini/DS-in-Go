package linkedlist

import (
    "fmt"
    "strings"
)

type Node struct {
    Value int
    Next  *Node
}

func CreateLinkedList(length int) *Node {
    if length < 1 {
        return nil
    }
    
    head := &Node{Value: 1}
    current := head
    
    for i := 2; i <= length; i++ {
        current.Next = &Node{Value: i}
        current = current.Next
    }
    
    return head
}

func (head *Node) DeleteNode(node *Node) bool {
    if node == nil {
        return false
    }
    
    if node == head {
        if head.Next != nil {
            *head = *head.Next
        } else {
            head = nil
        }
        return true
    }
    
    curr := head
    for curr.Next != nil && curr.Next != node {
        curr = curr.Next
    }
    
    if curr.Next == nil {
        return false
    }
    
    curr.Next = curr.Next.Next
    return true
}

func (head *Node) AddLast(value int) {
    newNode := &Node{Value: value}
    
    if head == nil {
        head = newNode
        return
    }
    
    curr := head 
    for curr.Next != nil {
        curr = curr.Next
    }
    curr.Next = newNode
}

func (head *Node) AddAtPosition(pos int, val int) error {
    if pos < 1 {
        return fmt.Errorf("position should be greater than or equal to 1")
    }
    
    if pos == 1 {
        return fmt.Errorf("to add at the start, use AddFirst method")
    }
    
    newNode := &Node{Value: val}
    curr := head
    
    for i := 1; i < pos-1; i++ {
        if curr == nil {
            return fmt.Errorf("position is greater than the length of the linked list")
        }
        curr = curr.Next
    }
    
    if curr == nil {
        return fmt.Errorf("position is greater than the length of the linked list")
    }
    
    newNode.Next = curr.Next
    curr.Next = newNode
    
    return nil
}

func (head *Node) AddFirst(value int) *Node {
    return &Node{
        Value: value,
        Next: head,
    }
}

func (node *Node) String() string {
    if node == nil {
        return "Empty List"
    }
    
    var sb strings.Builder
    current := node
    
    for current != nil {
        sb.WriteString(fmt.Sprintf("%d", current.Value))
        if current.Next != nil {
            sb.WriteString("->")
        }
        current = current.Next
    }
    
    return sb.String()
}

func (node *Node) PrintLinkedList() {
    fmt.Println(node.String())
}

func (head *Node) GetLength() int {
    length := 0
    curr := head
    
    for curr != nil {
        curr = curr.Next
        length++
    } 
    
    return length
}
