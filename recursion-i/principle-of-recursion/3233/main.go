package main

import "fmt"

func main() {
	tree := &TreeNode{
		Val:   4,
		Right: &TreeNode{Val: 7},
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
			Left:  &TreeNode{Val: 1},
		},
	}
	fmt.Println(searchBST(tree, 2))
}

// TreeNode is binary search tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else if root.Val > val {
		return searchBST(root.Left, val)
	}
	return root
}
