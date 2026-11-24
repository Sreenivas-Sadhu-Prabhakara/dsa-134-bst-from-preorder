package main

import (
	"fmt"
	"math"
)

// TreeNode is a binary search tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// FromPreorder reconstructs the unique BST with the given preorder
// traversal in one pass: each recursive call consumes values while they
// stay under an upper bound inherited from the ancestors.
func FromPreorder(preorder []int) *TreeNode {
	i := 0
	var build func(bound int) *TreeNode
	build = func(bound int) *TreeNode {
		if i == len(preorder) || preorder[i] > bound {
			return nil
		}
		n := &TreeNode{Val: preorder[i]}
		i++
		n.Left = build(n.Val)  // left subtree: everything below this value
		n.Right = build(bound) // right subtree: up to the inherited bound
		return n
	}
	return build(math.MaxInt)
}

func preorderOf(n *TreeNode) []int {
	if n == nil {
		return nil
	}
	out := []int{n.Val}
	out = append(out, preorderOf(n.Left)...)
	return append(out, preorderOf(n.Right)...)
}

func main() {
	root := FromPreorder([]int{8, 5, 1, 7, 10, 12})
	fmt.Println(root.Val, preorderOf(root))
}
