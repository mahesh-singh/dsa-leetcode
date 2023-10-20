package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	if isSame(root, subRoot) {
		return true
	} else {
		return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
	}

}

func isSame(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val == q.Val {
		return isSame(p.Left, q.Left) && isSame(p.Right, q.Right)
	} else {
		return false
	}
}
