
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	clonedNodes := make([]*Node, 101)
	cloneNode := &Node{Val: node.Val}
	clonedNodes[node.Val] = cloneNode
	dfs(node, clonedNodes)
	return clonedNodes[node.Val]
}

func dfs(currNode *Node, clonedNodes []*Node) {
	for _, neighbor := range currNode.Neighbors {
		if clonedNodes[neighbor.Val] == nil {
			cloneNode := &Node{Val: neighbor.Val}
			clonedNodes[neighbor.Val] = cloneNode
			dfs(neighbor, clonedNodes)
		}
		clonedNodes[currNode.Val].Neighbors = append(clonedNodes[currNode.Val].Neighbors, clonedNodes[neighbor.Val])
	}
}
