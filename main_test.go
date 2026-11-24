package main

import (
	"reflect"
	"sort"
	"testing"
)

func inorderOf(n *TreeNode) []int {
	if n == nil {
		return nil
	}
	return append(append(inorderOf(n.Left), n.Val), inorderOf(n.Right)...)
}

func TestFromPreorder(t *testing.T) {
	pre := []int{8, 5, 1, 7, 10, 12}
	root := FromPreorder(pre)
	// Round trip: rebuilt tree reproduces the preorder...
	if got := preorderOf(root); !reflect.DeepEqual(got, pre) {
		t.Errorf("preorder = %v, want %v", got, pre)
	}
	// ...and its inorder is sorted, which proves BST-ness.
	in := inorderOf(root)
	want := append([]int{}, pre...)
	sort.Ints(want)
	if !reflect.DeepEqual(in, want) {
		t.Errorf("inorder = %v, want %v", in, want)
	}

	if FromPreorder(nil) != nil {
		t.Error("empty preorder should build nil")
	}
	// Descending preorder: a pure left chain.
	chain := FromPreorder([]int{5, 4, 3})
	if chain.Val != 5 || chain.Left.Val != 4 || chain.Left.Left.Val != 3 || chain.Right != nil {
		t.Error("descending preorder should build a left chain")
	}
	// Ascending preorder: a pure right chain.
	rchain := FromPreorder([]int{3, 4, 5})
	if rchain.Val != 3 || rchain.Right.Val != 4 || rchain.Right.Right.Val != 5 || rchain.Left != nil {
		t.Error("ascending preorder should build a right chain")
	}
}
