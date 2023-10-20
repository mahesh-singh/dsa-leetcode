package main

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxD := 0
	depth(root, &maxD)
	return maxD
}

func depth(root *TreeNode, maxD *int) int {
	if root == nil {
		return 0
	}

	ld := depth(root.Left, maxD)
	rd := depth(root.Right, maxD)

	if ld+rd > *maxD {
		*maxD = ld + rd
	}

	if ld > rd {
		return ld + 1
	} else {
		return rd + 1
	}
}
