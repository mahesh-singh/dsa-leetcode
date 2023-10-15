package leetcode

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := maxDepth(root.Left)
	rh := maxDepth(root.Right)
	if lh > rh {
		return 1 + lh
	} else {
		return 1 + rh
	}
}
