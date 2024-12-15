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
	fmt.Printf("The length of Binary Tree is %d\n", bt.GetLength())
}
