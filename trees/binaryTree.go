package trees

import (
	"fmt"
	"math"
)

type TreeNode[T comparable] struct {
	data  T
	left  *TreeNode[T]
	right *TreeNode[T]
}

type BinaryTree[T comparable] struct {
	root *TreeNode[T]
}

func NewBinaryTree[T comparable]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

func (bt *BinaryTree[T]) MakeBinaryTree(arr []T) error {
	if len(arr) == 0 || arr == nil {
		return fmt.Errorf("input array is empty")
	}
	bt.root = &TreeNode[T]{
		data: arr[0],
	}
	helper(bt.root, arr, 0)
	return nil
}

func helper[T comparable](node *TreeNode[T], arr []T, index int) {
	if node == nil || index >= len(arr) {
		return
	}

	leftIndex := 2*index + 1
	rightIndex := 2*index + 2

	if leftIndex < len(arr) {
		node.left = &TreeNode[T]{data: arr[leftIndex]}
		helper(node.left, arr, leftIndex)
	}

	if rightIndex < len(arr) {
		node.right = &TreeNode[T]{data: arr[rightIndex]}
		helper(node.right, arr, rightIndex)
	}
}

func (bt *BinaryTree[T]) PrintTree() {
	if bt.root == nil {
		panic("tree is empty cant print empty tree")
	}
	fmt.Println("Enter:")
	fmt.Println("1 for pre-order traversal")
	fmt.Println("2 for post-order traversal")
	fmt.Println("3 for in-order traversal")
	fmt.Println("4 for level-order traversal")
	var choice int
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		preorder(bt.root)
		return
	case 2:
		postOrder(bt.root)
		return
	case 3:
		inOrder(bt.root)
		return
	case 4:
		bfs(bt.root)
		return
	default:
		fmt.Println("Invalid choice please enter 1,2,3 or 4")
	}

}

func preorder[T comparable](node *TreeNode[T]) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	preorder(node.left)
	preorder(node.right)
}

func postOrder[T comparable](node *TreeNode[T]) {
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Println(node.data)
}

func inOrder[T comparable](node *TreeNode[T]) {
	if node == nil {
		return
	}
	inOrder(node.left)
	fmt.Println(node.data)
	inOrder(node.right)
}

func bfs[T comparable](node *TreeNode[T]) {
	if node == nil {
		return
	}
	queue := make([]*TreeNode[T], 0)
	queue = append(queue, node)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		fmt.Println(curr.data)

		if curr.left != nil {
			queue = append(queue, curr.left)
		}

		if curr.right != nil {
			queue = append(queue, curr.right)
		}
	}
}

func (bt *BinaryTree[T]) GetLength() int {
	if bt.root == nil {
		return 0
	}
	return bt.root.GetLength()
}

func (node *TreeNode[T]) GetLength() int {
	if node == nil {
		return 0
	}
	leftH := 0
	if node.left != nil {
		leftH = node.left.GetLength()
	}
	rightH := 0
	if node.right != nil {
		rightH = node.right.GetLength()
	}
	return 1 + int(math.Max(float64(leftH), float64(rightH)))
}

func (bt *BinaryTree[T]) Delete(value T) {
	bt.root = deleteNode(bt.root, value)
}

func deleteNode[T comparable](root *TreeNode[T], value T) *TreeNode[T] {
	if root == nil {
		return nil
	}

	if root.data == value {
		if root.left == nil && root.right == nil {
			return nil
		}
		if root.left == nil {
			return root.right
		}
		if root.right == nil {
			return root.left
		}

		return root.left
	}

	root.left = deleteNode(root.left, value)
	root.right = deleteNode(root.right, value)

	return root
}
