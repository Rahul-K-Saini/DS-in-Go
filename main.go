package main

import (
	"dsa-in-go/trees"
	"fmt"
)

func main() {
	fmt.Println("Binary Tree")
	bt := trees.NewBinaryTree[int]()
	var arr []int
	for i := 0; i < 20; i++ {
		arr = append(arr, i)
	}
	bt.MakeBinaryTree(arr)
	bt.PrintTree()
	bt.Delete(1)
	bt.Delete(2)
	bt.Delete(3)
	bt.Delete(4)
	fmt.Printf("The length of Binary Tree is %d\n", bt.GetLength())
}
