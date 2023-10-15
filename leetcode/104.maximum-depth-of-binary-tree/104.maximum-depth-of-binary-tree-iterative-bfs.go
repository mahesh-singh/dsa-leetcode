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

	que := make([]*TreeNode, 0)
	que = append(que, root)
	level := 0
	for len(que) > 0 {
		currLen := len(que)

		for i := 0; i < currLen; i++ {

			currNode := que[0]
			que = que[1:]
			if currNode.Left != nil {
				que = append(que, currNode.Left)
			}

			if currNode.Right != nil {
				que = append(que, currNode.Right)
			}

		}
		level = level + 1
	}
	return level
}
