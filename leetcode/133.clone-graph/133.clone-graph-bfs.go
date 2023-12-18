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

	copies := make([]*Node, 101)

	copy := &Node{Val: node.Val}
	copies[node.Val] = copy
	queue := []*Node{}
	queue = append(queue, node)

	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]
		for _, neighbor := range currNode.Neighbors {
			if copies[neighbor.Val] == nil {
				copies[neighbor.Val] = &Node{Val: neighbor.Val}
				queue = append(queue, neighbor)
			}
			copies[currNode.Val].Neighbors = append(copies[currNode.Val].Neighbors, copies[neighbor.Val])
		}

	}

	return copies[node.Val]

}