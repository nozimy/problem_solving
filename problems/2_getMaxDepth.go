package problems

import (
	"fmt"
)

// There is a binary search tree. It is required to develop a method getMaxDepth(BinaryNode root)
// which returns the maximum depth of the tree.

type BinaryNode struct {
	data  int
	Left  *BinaryNode
	Right *BinaryNode
}

func RunGetMaxDepthProblem() {
	root := newNode(8)

	root.Left = newNode(3)
	root.Right = newNode(9)
	root.Left.Left = newNode(1)
	root.Left.Right = newNode(5)
	root.Left.Right.Left = newNode(4)

	result := getMaxDepth(root)

	fmt.Printf("Height of binary search tree is: %d\n", result)
}

func getMaxDepth(root *BinaryNode) int {
	if root == nil {
		return 0
	}

	leftDepth := getMaxDepth(root.Left)
	rightDepth := getMaxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}

func newNode(data int) *BinaryNode {
	return &BinaryNode{
		data:  data,
		Left:  nil,
		Right: nil,
	}
}
