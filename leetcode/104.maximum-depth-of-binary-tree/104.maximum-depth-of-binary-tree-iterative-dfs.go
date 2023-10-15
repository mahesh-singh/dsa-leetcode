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

	stack := make([]*TreeNode, 0)

	mapLevel := make(map[*TreeNode]int)
	stack = append(stack, root)
	mapLevel[root] = 1
	maxLevel := 1
	for len(stack) > 0 {

		popnode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		level := mapLevel[popnode]

		if popnode.Left != nil {
			stack = append(stack, popnode.Left)
			mapLevel[popnode.Left] = level + 1
		}
		if popnode.Right != nil {
			stack = append(stack, popnode.Right)
			mapLevel[popnode.Right] = level + 1
		}
		if level > maxLevel {
			maxLevel = level
		}
	}

	return maxLevel
}
