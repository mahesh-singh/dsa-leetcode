package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if abs(depth(root.Left)-depth(root.Right)) <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	} else {
		return false
	}
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := depth(root.Left)
	rd := depth(root.Right)

	if ld > rd {
		return ld + 1
	} else {
		return rd + 1
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
